package server

import (
	"bufio"
	"context"
	"errors"
	"ftp-server-go/internal/config"
	logger "ftp-server-go/internal/log"
	"ftp-server-go/internal/service"
	"ftp-server-go/internal/service/engine"
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
	cfg    *config.ServerConfig
	engine engine.Engine
}

func NewServer(cfg *config.Config, e engine.Engine) service.Service {
	return &server{
		cfg:    &cfg.Server,
		engine: e,
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

	sess := s.engine.OpenSession(c.RemoteAddr().String())
	s.engine.Greet(sess)

	br := bufio.NewReaderSize(c, 2048)
	if err := s.reply(c, s.engine.Greet(sess)); err != nil {
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
				_ = s.reply(c, engine.Reply{
					Code:     engine.FTPCodeSyntaxErrorCommandNotRecognized,
					Messages: []string{"command line too long"},
				})
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
			if err := s.reply(c, engine.Reply{
				Code:     engine.FTPCodeSyntaxErrorInParameters,
				Messages: []string{"empty command"},
			}); err != nil {
				return
			}
			continue
		}

		log.Debug("ftp: command", slog.String("cmd", cmd))
		d := s.engine.Exec(ctx, sess, cmd)
		if err := s.reply(c, d.Reply); err != nil {
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

func (s *server) reply(w io.Writer, reply engine.Reply) error {
	var builder strings.Builder
	code := strconv.Itoa(int(reply.Code))
	idx := 0

	if len(reply.Messages) > 1 || len(reply.Body) > 0 {
		for ; idx < len(reply.Messages)-1; idx++ {
			builder.WriteString(code)
			builder.WriteString("-")
			builder.WriteString(reply.Messages[idx])
			builder.WriteString("\r\n")
		}

		if len(reply.Body) > 0 {
			for _, b := range reply.Body {
				builder.WriteString(" ")
				builder.WriteString(b)
				builder.WriteString("\r\n")
			}
			builder.WriteString(code)
			builder.WriteString(" ")
			builder.WriteString("END")
		}
	} else {
		builder.WriteString(code)
		builder.WriteString(" ")
		builder.WriteString(reply.Messages[idx])
	}
	builder.WriteString("\r\n")

	_, err := w.Write([]byte(builder.String()))
	return err
}
