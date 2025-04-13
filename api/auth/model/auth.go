package model

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Username string    `gorm:"column:title" json:"title"`
	Password string    `gorm:"column:password" json:"password"`
	FullName string    `gorm:"column:full_name" json:"fullName"`
}

func (t *User) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	log.Info("Generated UUIT before create: ", t.ID)
	return
}
