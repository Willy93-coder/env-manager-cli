// Package models contains al database models
package models

import "time"

type Environment struct {
	ID        int64
	ProjectID int64
	Key       string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
