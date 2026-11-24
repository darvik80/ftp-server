package engine

import (
	"context"
	"errors"
	"ftp-server-go/internal/service/engine/account"
)

func (e *defEngine) handleUSER(ctx context.Context, s *Session, args []string) Reply {
	if len(args) != 1 {
		return ReplyByCode(FTPCodeSyntaxErrorCommandNotRecognized)
	}

	s.User = args[0]
	s.Auth = AuthNeedPass

	return ReplyByCode(FTPCodeNameUserAccepted)
}

func (e *defEngine) handlePASS(ctx context.Context, s *Session, args []string) Reply {
	if s.LastCommand != "USER" {
		return ReplyByCode(FTPCodeIncorrectSequenceOfCommands)
	}
	if len(args) != 1 {
		return ReplyByCode(FTPCodeSyntaxErrorCommandNotRecognized)
	}

	_, err := e.acc.Authenticate(ctx, s.User, args[0])
	switch {
	case errors.Is(err, account.ErrAuthFailed) || errors.Is(err, account.ErrUserNotFound):
		return ReplyByCode(FTPCodeLoginIncorrect)
	case err != nil:
		return ReplyByCode(FTPCodeServiceUnavailable)
	default:
	}

	s.Auth = AuthOK

	return ReplyByCode(FTPCodeCommandCompleted)
}

func (e *defEngine) handleHELP(ctx context.Context, s *Session, args []string) Reply {
	return Reply{Code: FTPCodeCommandCompleted, Messages: []string{
		"HELP command",
	}}
}
