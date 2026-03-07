// Package models contains al database models
package models

import "time"

type Project struct {
	ID          int64
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
