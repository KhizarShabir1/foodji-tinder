package database

import (
	"database/sql"
)

type Provider struct {
	db *sql.DB
}

// NewProvider ...
func NewProvider(dbp *sql.DB) *Provider {
	p := &Provider{
		db: dbp,
	}

	return p
}
