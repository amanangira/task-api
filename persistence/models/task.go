package models

import (
	"time"
)

// TODO - omit ID from being read from request input, introduce DTOs
// TODO - update Priority type to enum in DB and deal with it's own go type

// Task - represents a user task in DB
type Task struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Priority    TaskPriority `json:"priority"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DueAt       time.Time    `json:"due_at"`
}
