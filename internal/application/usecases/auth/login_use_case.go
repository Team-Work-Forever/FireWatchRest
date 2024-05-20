package usecases

import (
	"errors"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/jwt"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pwd"
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
	// validate email

	foundAuth, err := uc.authRepository.GetAuthByEmail(request.Email)

	if err != nil {
		return nil, err
	}

	// verify if the password is correct
	if ok := pwd.CheckPasswordHash(request.Password, foundAuth.Salt, foundAuth.Password); !ok {
		return nil, errors.New("the email or password is wrong")
	}

	// generate jwt tokens
	accessToken, refreshToken, err := jwt.CreateAuthTokens(foundAuth)

	if err != nil {
		return nil, err
	}

	return &contracts.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
