package database

import (
	"time"

	"github.com/KhizarShabir1/foodji-tinder/foodji"
	"github.com/google/uuid"
)

var _ foodji.SessionProvider = (*Provider)(nil)

func (p *Provider) CreateSession() (*foodji.Session, error) {

	// Generate a unique session ID using Google's uuid package
	sessionID := uuid.New().String()

	// Create a new Session instance with the generated session ID and current time
	session := &foodji.Session{
		ID:        sessionID,
		CreatedAt: time.Now(),
	}

	// Insert the new session into the database
	_, err := p.db.Exec("INSERT INTO session (id, created_at) VALUES ($1, $2)", session.ID, session.CreatedAt)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (p *Provider) GetSessions() ([]foodji.Session, error) {
	rows, err := p.db.Query("SELECT id, created_at FROM session")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sessions := []foodji.Session{}
	for rows.Next() {
		var session foodji.Session
		err := rows.Scan(&session.ID, &session.CreatedAt)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}

	return sessions, nil
}
