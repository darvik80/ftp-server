package engine

import (
	"context"
	"strconv"
	"strings"
)

type Engine struct {
	cmd map[string]FtpHandler
}

func (*Engine) Name() string {
	return "engine"
}

func NewEngine() *Engine {
	return &Engine{
		cmd: make(map[string]FtpHandler),
	}
}

func (e *Engine) RegisterCommand(name string, handler FtpHandler) {
	e.cmd[name] = handler
}

func (e *Engine) Exec(ctx context.Context, cmd string) (string, error) {
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

func (e Engine) reply(ctx context.Context, reply Reply) string {
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
