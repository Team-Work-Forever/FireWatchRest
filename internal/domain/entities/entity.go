package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EntityBase struct {
	ID        string    `gorm:"type:uuid;primaryKey;column:id"`
	CreatedAt time.Time `gorm:"column:created_at;<-:update"`
	UpdatedAt time.Time `gorm:"column:updated_at;<-:update"`
	DeletedAt time.Time `gorm:"column:deleted_at;<-:update"`
}

func (u *EntityBase) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	return nil
}
