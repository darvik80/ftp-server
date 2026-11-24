package engine

type AuthState int

const (
	AuthNone AuthState = iota
	AuthNeedPass
	AuthOK
)

type DataMode int

const (
	DataNone DataMode = iota
	DataActive
	DataPassive
)

type Representation struct {
	Type byte // 'A' или 'I'; старт — 'A'
	Mode byte // 'S'
	Stru byte // 'F'
}

type DataEndpoint struct {
	Mode DataMode
	Host string
	Port int
}

type Session struct {
	ID     string
	Remote string

	Auth     AuthState
	User     string
	Home     string // абсолютный корень пользователя, уже проверенный
	CanWrite bool

	Dir  string // относительный каталог внутри Home
	Repr Representation

	Data       *DataEndpoint // последний PORT или PASV
	RenameFrom string        // путь после RNFR, пока не пришёл RNTO

	LastCommand string
}
