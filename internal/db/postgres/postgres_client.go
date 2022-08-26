package postgres

import (
	"context"
	"database/sql"
)

// Client for the PostgreSQL.
type Client struct {
	*sql.DB
}

// NewPostgresClient returns a new client postgres.
func NewPostgresClient(source string) (*Client, error) {
	db, err := sql.Open("postgres", source)
	if err != nil {
		return nil, err
	}

	// confirm connection to ddbb
	err = db.PingContext(context.Background())
	if err != nil {
		return nil, err
	}

	return &Client{db}, nil
}
