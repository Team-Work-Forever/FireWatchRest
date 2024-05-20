package repositories

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/adapters"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"gorm.io/gorm"
)

type TokenRepostory struct {
	dbContext *gorm.DB
}

func NewTokenRepository(database adapters.Database) *TokenRepostory {
	return &TokenRepostory{
		dbContext: database.GetDatabase(),
	}
}

func (repo *TokenRepostory) Create(token *entities.Token) error {
	return repo.dbContext.Create(token).Error
}
