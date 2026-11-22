package engine

type SessionStatus int

const (
	SessionInactive = iota
	SessionActive
	SessionTerminated
	SessionClosed
)

type Session struct {
	User   string
	Status SessionStatus
}
