package engine

import "io"

type DataStream struct {
	Out io.ReadCloser  // RETR, LIST, NLST
	In  io.WriteCloser // STOR, APPE
}

type TransferStatus int

const (
	TransferDone       TransferStatus = iota
	TransferConnFailed                // 425, канал данных не открылся
	TransferAborted                   // 426, обрыв посередине
	TransferLocalError                // 451, файл или каталог
)

type TransferResult struct {
	Kind   TransferKind
	Status TransferStatus
	Err    error // в журнал; клиенту уходит только код из Reply
}

type CloseReason int

const (
	CloseClient   CloseReason = iota // EOF, ответ не слать
	CloseQuit                        // 221 уже ушёл из Exec
	CloseIdle                        // 421
	CloseShutdown                    // 421
)
