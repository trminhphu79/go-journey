package model

import (
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TaskStatus string

const (
	Todo       TaskStatus = "TODO"
	InProgress TaskStatus = "IN_PROGRESS"
	Rejected   TaskStatus = "REJECTED"
	Resolved   TaskStatus = "RESOLVED"
)

type Task struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;" json:"id"`
	Title       string     `gorm:"column:title" json:"title"`
	Description string     `gorm:"column:description" json:"description"`
	Tags        []string   `gorm:"type:text[];column:tags" json:"tags,omitempty"`
	Slug        string     `gorm:"column:slug" json:"slug"`
	Status      TaskStatus `gorm:"column:status" json:"status"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updatedAt"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	log.Info("Generated UUIT before create: ", t.ID)
	return
}
