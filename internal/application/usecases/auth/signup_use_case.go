package usecases

import (
	"fmt"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	valueobjects "github.com/Team-Work-Forever/FireWatchRest/internal/domain/value_objects"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/jwt"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
)

type SignUpUseCase struct {
	authRepository *repositories.AuthRepository
}

func NewSignUpUseCase(authRepository *repositories.AuthRepository) *SignUpUseCase {
	return &SignUpUseCase{
		authRepository: authRepository,
	}
}

func (uc *SignUpUseCase) Handle(request *contracts.SignUpRequest) (*contracts.AuthResponse, error) {
	// validate data

	if ok := uc.authRepository.ExistsUserWithEmail(request.Email); ok {
		return nil, fmt.Errorf("o utilizador já existe")
	}

	if ok := uc.authRepository.ExistsUserWithNif(request.NIF); ok {
		return nil, fmt.Errorf("o utilizador já existe")
	}

	auth := entities.NewAuth(
		request.Email,
		request.Password,
		request.NIF,
	)

	user := entities.NewUser(
		"adashdahhhad",
		request.FirstName,
		request.LastName,
		valueobjects.Phone{
			PhoneCode:   request.PhoneCode,
			PhoneNumber: request.PhoneNumber,
		},
		valueobjects.Address{
			Street:  request.Street,
			Number:  request.StreetPort,
			ZipCode: request.ZipCode,
			City:    request.City,
		},
	)

	if err := uc.authRepository.CreateUser(auth, user); err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := jwt.CreateAuthTokens(auth)

	if err != nil {
		return nil, err
	}

	return &contracts.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
