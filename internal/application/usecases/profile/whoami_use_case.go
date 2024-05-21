package usecases

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
)

type WhoamiUseCase struct {
	authRepository    *repositories.AuthRepository
	profileRepository *repositories.ProfileRepository
}

func NewWhoamiUseCase(
	authRepository *repositories.AuthRepository,
	profileRepository *repositories.ProfileRepository,
) *WhoamiUseCase {
	return &WhoamiUseCase{
		authRepository:    authRepository,
		profileRepository: profileRepository,
	}
}

func (w *WhoamiUseCase) Handle(request contracts.WhoamiRequest) (*contracts.ProfileResponse, error) {
	// fetch auth
	foundAuth, err := w.authRepository.GetAuthById(request.UserId)

	if err != nil {
		return nil, err
	}

	// fetch user
	profileFound, err := w.profileRepository.GetUserByAuthId(foundAuth.ID)

	if err != nil {
		return nil, err
	}

	return &contracts.ProfileResponse{
		Email:     foundAuth.Email.GetValue(),
		FirstName: profileFound.FirstName,
		LastName:  profileFound.LastName,
		Phone: contracts.PhoneResponse{
			CountryCode: profileFound.PhoneNumber.CountryCode,
			Number:      profileFound.PhoneNumber.Number,
		},
		Address: contracts.AddressResponse{
			Street: profileFound.Address.Street,
			Number: profileFound.Address.Number,
			ZipCode: contracts.ZipCodeResponse{
				Value: profileFound.Address.ZipCode,
			},
			City: profileFound.Address.City,
		},
		UserType: "admin",
	}, nil
}
