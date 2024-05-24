package repositories

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"gorm.io/gorm"
)

type TokenRepostory struct {
	dbContext *gorm.DB
}

func NewTokenRepository(database *gorm.DB) *TokenRepostory {
	return &TokenRepostory{
		dbContext: database,
	}
}

func (repo *TokenRepostory) GetByToken(tokenString string) (*entities.Token, error) {
	var token *entities.Token

	if err := repo.dbContext.Where("token = ?", tokenString).First(&token).Error; err != nil {
		return nil, err
	}

	return token, nil
}

func (repo *TokenRepostory) Create(token *entities.Token) error {
	return repo.dbContext.Create(token).Error
}

func (repo *TokenRepostory) Delete(token *entities.Token) error {
	return repo.dbContext.Delete(&token).Error
}
