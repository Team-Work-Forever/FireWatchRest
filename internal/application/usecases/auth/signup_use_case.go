package usecases

import (
	"errors"
	"fmt"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
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
	email, err := vo.NewEmail(request.Email)

	if err != nil {
		return nil, err
	}

	password, err := vo.NewPassword(request.Password)

	if err != nil {
		return nil, err
	}

	nif, err := vo.NewNIF(request.NIF)

	if err != nil {
		return nil, err
	}

	if request.FirstName == "" {
		return nil, errors.New("provide an first name")
	}

	if request.LastName == "" {
		return nil, errors.New("provide an last name")
	}

	phone, err := vo.NewPhone(request.PhoneCode, request.PhoneNumber)

	if err != nil {
		return nil, err
	}

	zipCode, err := vo.NewZipCode(request.ZipCode)

	if err != nil {
		return nil, err
	}

	address, err := vo.NewAddress(request.Street, request.StreetPort, *zipCode, request.City)

	if err != nil {
		return nil, err
	}

	if ok := uc.authRepository.ExistsUserWithEmail(email); ok {
		return nil, fmt.Errorf("o utilizador já existe")
	}

	if ok := uc.authRepository.ExistsUserWithNif(nif); ok {
		return nil, fmt.Errorf("o utilizador já existe")
	}

	auth := entities.NewAuth(
		*email,
		*password,
		*nif,
	)

	user := entities.NewUser(
		"adashdahhhad",
		request.FirstName,
		request.LastName,
		*phone,
		*address,
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
