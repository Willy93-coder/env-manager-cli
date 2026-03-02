// Package db handles the PostgreSQL database connection
package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DB wraps sql.DB to provide a custom type
type DB struct {
	*sql.DB
}

// New opens a connection to PostgreSQL and verifies it with a ping
func New(dsn string) (*DB, error) {
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not open database connection: %w", err)
	}

	// Ping actually tries to connect
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping database: %w", err)
	}

	return &DB{sqlDB}, nil
}
