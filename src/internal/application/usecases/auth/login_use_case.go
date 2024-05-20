package usecases

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	jwtService "github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
)

type LoginUseCase struct {
	authRepository *repositories.AuthRepository
}

func NewLoginUseCase(authRepository *repositories.AuthRepository) *LoginUseCase {
	return &LoginUseCase{
		authRepository: authRepository,
	}
}

func (uc *LoginUseCase) Handle(request *contracts.LoginRequest) (*contracts.AuthResponse, error) {
	// verify if there is an user called like that by email
	// verify if the password is correct

	// generate jwt tokens
	accessToken, refreshToken, err := jwtService.CreateAuthTokens()

	if err != nil {
		return nil, err
	}

	return &contracts.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
