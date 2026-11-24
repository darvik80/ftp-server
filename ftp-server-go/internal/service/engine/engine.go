package engine

import (
	"context"
)

type Engine interface {
	RegisterCommands()
	OpenSession(remote string) *Session
	Greet(s *Session) Reply
	Exec(ctx context.Context, s *Session, line string) Decision
	OpenData(ctx context.Context, s *Session, t Transfer) (DataStream, error)
	CompleteTransfer(ctx context.Context, s *Session, result TransferResult) Reply
	CloseSession(s *Session, reason CloseReason) Reply
}

type PassivePorts interface {
	Listen(ctx context.Context, sessionID string) ( /*host*/ string /*port */, int, error)
	Close(sessionID string)
}
