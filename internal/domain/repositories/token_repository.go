package repositories

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/key"
)

type TokenRepostory struct {
	kvService *key.KeyValueService
}

func NewTokenRepository(
	kvService *key.KeyValueService,
) *TokenRepostory {
	return &TokenRepostory{
		kvService: kvService,
	}
}

func (repo *TokenRepostory) GetByToken(tokenString string, tt entities.TokenType) (*entities.Token, error) {
	var token entities.Token

	if err := repo.kvService.Get(key.Key{
		Tag:   tt.GetType(),
		Value: tokenString,
	}, &token); err != nil {
		return nil, err
	}

	return &token, nil
}

func (repo *TokenRepostory) Create(token *entities.Token) error {
	return repo.kvService.Store(token)
}

func (repo *TokenRepostory) Delete(token *entities.Token) error {
	return repo.kvService.Delete(token)
}
