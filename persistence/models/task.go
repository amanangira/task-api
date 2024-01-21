package models

import (
	"time"
)

// TODO - omit ID from being read from request input, introduce DTOs
// TODO - update Priority type to enum in DB and deal with it's own go type

// Task - represents a user task in DB
type Task struct {
	ID          string       `json:"id" db:"id"`
	Title       string       `json:"title" db:"title"`
	Description string       `json:"description" db:"description"`
	Priority    TaskPriority `json:"priority" db:"priority"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	DueAt       time.Time    `json:"due_at" db:"due_at"`
}
