package usecases

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/jwt"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pwd"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type LoginUseCase struct {
	authRepository *repositories.AuthRepository
}

func NewLoginUseCase(authRepository *repositories.AuthRepository) *LoginUseCase {
	return &LoginUseCase{
		authRepository: authRepository,
	}
}

func (uc *LoginUseCase) Handle(request contracts.LoginRequest) (*contracts.AuthResponse, error) {
	email, err := vo.NewEmail(request.Email)

	if err != nil {
		return nil, err
	}

	foundAuth, err := uc.authRepository.GetAuthByEmail(email)

	if err != nil {
		return nil, exec.USER_NOT_FOUND
	}

	password, err := vo.NewPassword(request.Password)

	if err != nil {
		return nil, err
	}

	// verify if the password is correct
	if ok := pwd.CheckPasswordHash(password.GetValue(), foundAuth.Salt, foundAuth.Password.GetValue()); !ok {
		return nil, exec.PASSWORD_WRONG
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
