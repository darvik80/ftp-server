package engine

type Decision struct {
	Reply    Reply
	Close    bool // QUIT и ответы класса 421
	Transfer *Transfer
}

type TransferKind int

const (
	TransferRetr TransferKind = iota
	TransferStor
	TransferAppe
	TransferList
	TransferNlst
)

type Transfer struct {
	Kind TransferKind
	Path string // Абсолютный путь внутри Home, уже проверенный
}
