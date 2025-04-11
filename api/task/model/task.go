package model

import (
	"time"
)

type TaskStatus string

const (
	Todo       TaskStatus = "TODO"
	InProgress TaskStatus = "IN_PROGRESS"
	Rejected   TaskStatus = "REJECTED"
	Resolved   TaskStatus = "RESOLVED"
)

type Task struct {
	Title       string     `validate:"required"`
	Description string     `validate:"required"`
	Tags        []string   `validate:"required"`
	Slug        string     `validate:"required"`
	Status      TaskStatus `validate:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
