package foodji

import (
	"time"
)

// SessionProvider gives access to sessions
type SessionProvider interface {
	CreateSession() (*Session, error)
	GetSessions() ([]Session, error)
}

type Session struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
