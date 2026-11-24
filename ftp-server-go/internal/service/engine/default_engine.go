package engine

import (
	"context"
	"ftp-server-go/internal/service/engine/account"
	"strings"
)

type defEngine struct {
	Engine

	cmd     map[string]FtpHandler
	authCmd map[string]FtpHandler

	acc account.Accounts
}

func NewDefaultEngine(acc account.Accounts) Engine {
	return &defEngine{
		cmd:     map[string]FtpHandler{},
		authCmd: make(map[string]FtpHandler),
		acc:     acc,
	}
}

func (e *defEngine) RegisterCommands() {
	// rfc 959
	e.cmd["USER"] = e.handleUSER
	e.cmd["PASS"] = e.handlePASS
	e.cmd["HELP"] = e.handleHELP

	// команды доступные после авторизации
}

func (e *defEngine) OpenSession(remote string) *Session {
	return &Session{
		ID:     "123",
		Remote: remote,
		Auth:   AuthNone,
	}
}

func (e *defEngine) Greet(s *Session) Reply {
	return Reply{
		Code:     FTPCodeServiceReady,
		Messages: []string{"Service ready"},
	}
}

func (e *defEngine) parseCommand(line string) (string, []string) {
	args := strings.Split(line, " ")
	return args[0], args[1:]
}

func (e *defEngine) Exec(ctx context.Context, s *Session, line string) Decision {
	cmd, args := e.parseCommand(line)
	defer func() {
		s.LastCommand = cmd
	}()

	if handler, ok := e.cmd[cmd]; ok {
		return Decision{
			Reply: handler(ctx, s, args),
		}
	}
	return Decision{
		Reply: Reply{
			Code:     FTPCodeCommandNotImplemented,
			Messages: []string{"Command not implemented"},
		},
	}
}

func (e *defEngine) OpenData(ctx context.Context, s *Session, t Transfer) (DataStream, error) {
	return DataStream{}, nil
}

func (e *defEngine) CompleteTransfer(ctx context.Context, s *Session, result TransferResult) Reply {
	return Reply{}
}

func (e *defEngine) CloseSession(s *Session, reason CloseReason) Reply {
	return Reply{
		Code:     FTPCodeServiceClosing,
		Messages: []string{"Service closing"},
	}
}
