package usecases

import (
	repo "github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type UpdateProfileUseCase struct {
	authRepo    *repo.AuthRepository
	profileRepo *repo.ProfileRepository
}

func NewUpdateProfileUIseCase(
	authRepo *repo.AuthRepository,
	profileRepo *repo.ProfileRepository,
) *UpdateProfileUseCase {
	return &UpdateProfileUseCase{
		authRepo:    authRepo,
		profileRepo: profileRepo,
	}
}

func (uc *UpdateProfileUseCase) Handle(request contracts.UpdateProfileResponse) (*contracts.ProfileResponse, error) {
	foundProfile, err := uc.profileRepo.GetUserByAuthId(request.UserId)

	if err != nil {
		return nil, exec.USER_NOT_FOUND
	}

	foundAuth, err := uc.authRepo.GetAuthById(request.UserId)

	if err != nil {
		return nil, exec.USER_NOT_FOUND
	}

	if request.Email != "" {
		email, err := vo.NewEmail(request.Email)

		if err != nil {
			return nil, err
		}

		foundAuth.Email = *email

		if err := uc.authRepo.Update(foundAuth); err != nil {
			return nil, err
		}
	}

	if request.UserName != "" {
		foundProfile.UserName = request.UserName
	}

	if request.PhoneCode != "" || request.PhoneNumber != "" {
		if request.PhoneCode == "" || request.PhoneNumber == "" {
			return nil, exec.PHONE_PROVIDE
		}

		phone, err := vo.NewPhone(request.PhoneCode, request.PhoneNumber)
		if err != nil {
			return nil, err
		}

		foundProfile.PhoneNumber = *phone
	}

	if request.ZipCode != "" {
		zipCode, err := vo.NewZipCode(request.ZipCode)

		if err != nil {
			return nil, err
		}

		foundProfile.Address.ZipCode = zipCode.GetValue()
	}

	if request.Street != "" {
		foundProfile.Address.Street = request.Street
	}

	if request.StreetPort != nil {
		foundProfile.Address.SetStreetNumber(*request.StreetPort)
	}

	if request.City != "" {
		foundProfile.Address.City = request.City
	}

	if err := uc.profileRepo.Update(foundProfile); err != nil {
		return nil, err
	}

	return &contracts.ProfileResponse{
		Email:     foundAuth.Email.GetValue(),
		UserName:  foundProfile.UserName,
		FirstName: foundProfile.FirstName,
		LastName:  foundProfile.LastName,
		Phone: contracts.PhoneResponse{
			CountryCode: foundProfile.PhoneNumber.CountryCode,
			Number:      foundProfile.PhoneNumber.Number,
		},
		Address: contracts.AddressResponse{
			Street: foundProfile.Address.Street,
			Number: foundProfile.Address.Number,
			ZipCode: contracts.ZipCodeResponse{
				Value: foundProfile.Address.ZipCode,
			},
			City: foundProfile.Address.City,
		},
		UserType: foundAuth.GetRole(),
	}, nil
}
