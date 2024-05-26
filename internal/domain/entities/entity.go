package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Entity struct {
	CreatedAt time.Time       `gorm:"column:created_at;<-:update"`
	UpdatedAt time.Time       `gorm:"column:updated_at;<-:update"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

type EntityBase struct {
	Entity
	ID string `gorm:"type:uuid;primaryKey;column:id"`
}

func (e *EntityBase) GetId() string {
	if e.ID == "" {
		e.ID = uuid.New().String()
	}

	return e.ID
}

func (u *EntityBase) BeforeCreate(tx *gorm.DB) error {
	u.GetId()

	u.DeletedAt = nil
	return nil
}
