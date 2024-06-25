package usecases

import (
	repo "github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/services"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/upload"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type UpdateProfileUseCase struct {
	authRepo           *repo.AuthRepository
	profileRepo        *repo.ProfileRepository
	fileService        *upload.BlobService
	autarchyRepository *repo.AutarchyRepository
}

func NewUpdateProfileUIseCase(
	authRepo *repo.AuthRepository,
	profileRepo *repo.ProfileRepository,
	fileService *upload.BlobService,
	autarchyRepository *repo.AutarchyRepository,
) *UpdateProfileUseCase {
	return &UpdateProfileUseCase{
		authRepo:           authRepo,
		profileRepo:        profileRepo,
		fileService:        fileService,
		autarchyRepository: autarchyRepository,
	}
}

func (uc *UpdateProfileUseCase) Handle(request contracts.UpdateProfileResponse) (interface{}, error) {
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

	if request.Avatar != nil {
		file, err := request.Avatar.Open()

		if err != nil {
			return nil, err
		}

		defer file.Close()

		url, err := uc.fileService.UploadFile(&upload.UploadFile{
			Bucket:   upload.ClientBucket,
			FileName: request.Avatar.Filename,
			FileId:   foundAuth.GetId(),
			FileBody: file,
		})

		if err != nil {
			return nil, err
		}

		foundProfile.Picture = url
	}

	_, err = services.GetAutarchy(foundProfile.Address)

	if err != nil {
		return nil, err
	}

	if err := uc.profileRepo.Update(foundProfile); err != nil {
		return nil, err
	}

	return contracts.GetProfileResponse(foundAuth, foundProfile, uc.autarchyRepository)
}
