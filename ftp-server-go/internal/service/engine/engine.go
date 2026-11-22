package engine

import (
	"context"
	"io"
	"strconv"
	"strings"
)

type Engine interface {
	OpenSession(remote string) *Session
	Greet(s *Session) Reply
	Exec(ctx context.Context, s *Session, line string) Decision
	OpenData(ctx context.Context, s *Session, t Transfer) (DataStream, error)
	CompleteTransfer(ctx context.Context, s *Session, result TransferResult) Reply
	CloseSession(s *Session, reason CloseReason) Reply
}

type Account struct {
	User     string
	Home     string
	CanWrite bool
	CanRead  bool
	CanList  bool
}

type Accounts interface {
	Authenticate(ctx context.Context, user, password string) (Account, error)
}

type FileInfo struct {
	Name        string
	Size        int64
	Permissions string
}

type DirEntry struct {
	Name        string
	Permissions string
}

type Files interface {
	OpenRead(ctx context.Context, path string) (io.ReadCloser, error)
	OpenWrite(ctx context.Context, path string, append bool) (io.WriteCloser, error)
	ReadDir(ctx context.Context, path string) ([]DirEntry, error)
	Stat(ctx context.Context, path string) (FileInfo, error)
	Remove(ctx context.Context, path string) error
	Rename(ctx context.Context, from, to string) error
	Mkdir(ctx context.Context, path string) error
	Rmdir(ctx context.Context, path string) error
}

type PassivePorts interface {
	Listen(ctx context.Context, sessionID string) ( /*host*/ string /*port */, int, error)
	Close(sessionID string)
}

//func New(accounts Accounts, files Files, ports PassivePorts) Engine

type engine struct {
	cmd map[string]FtpHandler
}

func (*engine) Name() string {
	return "engine"
}

func NewEngine() *engine {
	return &engine{
		cmd: make(map[string]FtpHandler),
	}
}

func (e *engine) RegisterCommand(name string, handler FtpHandler) {
	e.cmd[name] = handler
}

func (e *engine) Exec(ctx context.Context, cmd string) (string, error) {
	args := strings.Fields(cmd)
	if len(args) == 0 {
		return e.reply(ctx,
			Reply{
				Code: FTPCodeSyntaxErrorCommandNotRecognized, Messages: []string{"Syntax error, command unrecognized"},
			},
		), nil
	}

	handler, ok := e.cmd[args[0]]
	if !ok {
		return e.reply(ctx,
			Reply{
				Code: FTPCodeSyntaxErrorCommandNotRecognized, Messages: []string{"Syntax error, command unrecognized"},
			},
		), nil
	}

	return e.reply(ctx, handler(ctx, args[1:])), nil
}

func (e engine) reply(ctx context.Context, reply Reply) string {
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

	return builder.String()
}
