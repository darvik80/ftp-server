package server

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"ftp-server-go/internal/config"
	logger "ftp-server-go/internal/log"
	"ftp-server-go/internal/service"
	"io"
	"log/slog"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type replyCode int

const (
	codeServiceReady   replyCode = 220
	codeSyntaxError    replyCode = 500
	codeNotImplemented replyCode = 502
)

type server struct {
	cfg *config.ServerConfig
}

func NewServer(cfg *config.Config) service.Service {
	return &server{
		cfg: &cfg.Server,
	}
}

func (s *server) Name() string {
	return "ftp-server"
}

func (s *server) Run(ctx context.Context) error {
	log := logger.FromContextNamed(ctx, "ftp-server")
	var lc net.ListenConfig

	log.Info("tcp: listening", slog.String("address", s.cfg.Address), slog.Int("port", s.cfg.Port))

	l, err := lc.Listen(ctx, "tcp", s.cfg.Address+":"+strconv.Itoa(s.cfg.Port))
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		log.Info("tcp: closing listener...")
		l.Close()
	}()

	guard := make(chan struct{}, s.cfg.MaxSessions)
	var connWG sync.WaitGroup
	for {
		conn, err := l.Accept()
		if err != nil {
			if ctx.Err() != nil {
				break
			}
			return err
		}

		// wait free sessions
		if !acquire(ctx, guard) {
			_ = conn.Close()
			break
		}

		connWG.Add(1)
		go func(c net.Conn) {
			defer func() {
				connWG.Done()
				<-guard
			}()
			s.handleConnection(ctx, c)
		}(conn)
	}

	log.Info("tcp: waiting for active connections to complete...")
	connWG.Wait()
	log.Info("tcp: all connections closed.")

	return nil
}

func (s *server) Stop(ctx context.Context) error {
	return nil
}

func (s *server) handleConnection(ctx context.Context, c net.Conn) {
	log := logger.FromContextNamed(ctx, "ftp-client")
	defer func() {
		log.Info("tcp: client disconnected", slog.String("address", c.RemoteAddr().String()))
		_ = c.Close()
	}()

	done := make(chan struct{})
	defer close(done)
	go func() {
		select {
		case <-ctx.Done():
			_ = c.Close()
		case <-done:
		}
	}()

	log.Info("tcp: client connected", slog.String("address", c.RemoteAddr().String()))

	br := bufio.NewReaderSize(c, 2048)
	if err := s.reply(c, codeServiceReady, "Service ready"); err != nil {
		return
	}

	for {
		err := c.SetReadDeadline(time.Now().Add(time.Duration(s.cfg.IdleTimeout) * time.Second))
		if err != nil {
			log.Error("ftp: set read deadline", slog.Any("err", err))
			return
		}
		line, err := br.ReadString('\n')
		if err != nil {
			var ne net.Error
			switch {
			case errors.Is(err, bufio.ErrBufferFull):
				_ = s.reply(c, codeSyntaxError, "command line too long")
				log.Warn("ftp: command line too long", slog.Int("bytes", len(line)))
			case errors.Is(err, io.EOF):
				log.Info("ftp: client closed the connection")
			case ctx.Err() != nil:
				log.Info("ftp: session interrupted by shutdown")
			case errors.As(err, &ne) && ne.Timeout():
				log.Info("ftp: idle timeout")
			default:
				log.Error("ftp: read failed", slog.Any("err", err))
			}
			return
		}

		cmd := strings.TrimRight(line, "\r\n")
		if cmd == "" {
			if err := s.reply(c, codeSyntaxError, "empty command"); err != nil {
				return
			}
			continue
		}

		log.Debug("ftp: command", slog.String("cmd", cmd))
		if err := s.reply(c, codeNotImplemented, "command not implemented"); err != nil {
			return
		}
	}
}

func acquire(ctx context.Context, guard chan struct{}) bool {
	select {
	case guard <- struct{}{}:
		return true
	case <-ctx.Done():
		return false
	}
}

func (s *server) reply(w io.Writer, code replyCode, text string) error {
	if _, err := w.Write(fmt.Appendf(nil, "%d %s\r\n", code, text)); err != nil {
		return fmt.Errorf("answer %d: %w", code, err)
	}
	return nil
}
