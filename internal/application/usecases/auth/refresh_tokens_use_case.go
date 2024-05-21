package usecases

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/jwt"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type RefreshTokensUseCase struct {
	authRepository *repositories.AuthRepository
}

func NewRefreshTokesUseCase(authRepository *repositories.AuthRepository) *RefreshTokensUseCase {
	return &RefreshTokensUseCase{
		authRepository: authRepository,
	}
}

func (r *RefreshTokensUseCase) Handle(request contracts.RefreshTokensRequest) (*contracts.AuthResponse, error) {
	// validate token
	claims, err := jwt.GetClaims(request.RefreshToken, &jwt.AuthClaims{})

	if err != nil {
		return nil, exec.TOKEN_ISNT_VALD
	}

	authClaims, ok := claims.(*jwt.AuthClaims)

	if !ok {
		return nil, exec.TOKEN_ISNT_VALD
	}

	// generatte new ones
	auth, err := r.authRepository.GetAuthById(authClaims.Subject)

	if err != nil {
		return nil, exec.USER_NOT_FOUND
	}

	accessToken, refreshToken, err := jwt.CreateAuthTokens(auth)

	if err != nil {
		return nil, err
	}

	// return them
	return &contracts.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
