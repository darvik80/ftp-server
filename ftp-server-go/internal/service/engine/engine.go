package engine

import (
	"context"
	"ftp-server-go/internal/service"
)

type ftpCode int

const (
	FTPCodeOK ftpCode = iota
)

type Command struct {
	Name    string
	Handler func(ctx context.Context, args []string) error
}

type Reply struct {
	Code    ftpCode
	Message string
}

type Engine struct {
	service.Service
}

func (*Engine) Name() string {
	return "engine"
}

func (*Engine) Run(ctx context.Context) error {
	return nil
}
