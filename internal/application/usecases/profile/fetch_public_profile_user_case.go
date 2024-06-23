package usecases

import (
	repo "github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type FetchPublicProfileUseCase struct {
	authRepository *repo.AuthRepository
	profileRepo    *repo.ProfileRepository
}

func NewFetchPublicProfileUseCase(
	authRepository *repo.AuthRepository,
	profileRepo *repo.ProfileRepository,
) *FetchPublicProfileUseCase {
	return &FetchPublicProfileUseCase{
		authRepository: authRepository,
		profileRepo:    profileRepo,
	}
}

func (uc *FetchPublicProfileUseCase) Handle(request contracts.PublicProfileRequest) (*contracts.PublicProfileResponse, error) {
	email, err := vo.NewEmail(request.Email)

	if err != nil {
		return nil, err
	}

	if !uc.authRepository.ExistsUserWithEmail(email) {
		return nil, exec.USER_NOT_FOUND
	}

	profileFound, err := uc.profileRepo.GetPublicProfile(email)

	if err != nil {
		return nil, err
	}

	return &contracts.PublicProfileResponse{
		Email:    email.Value,
		UserName: profileFound.UserName,
		Avatar:   profileFound.ProfileAvatar,
		NIF:      profileFound.NIF,
		Phone: contracts.PhoneResponse{
			CountryCode: profileFound.Phone.CountryCode,
			Number:      profileFound.Phone.Number,
		},
	}, nil
}
