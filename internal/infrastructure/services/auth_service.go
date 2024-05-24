package services

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/adapters"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
)

func GetAuthUser(userId string) (*entities.Auth, error) {
	db := adapters.GetDatabase()
	authRepository := repositories.NewAuthRepository(db)

	foundAuth, err := authRepository.GetAuthById(userId)

	if err != nil {
		return nil, err
	}

	return foundAuth, err
}
