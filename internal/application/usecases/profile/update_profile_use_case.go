package usecases

import (
	"strconv"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
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

func (uc *UpdateProfileUseCase) updateIdentityUser(identityUser entities.UpdateIdentity, foundAuth *entities.Auth, request contracts.UpdateProfileResponse) (interface{}, error) {
	if request.PhoneCode != "" || request.PhoneNumber != "" {
		if request.PhoneCode == "" || request.PhoneNumber == "" {
			return nil, exec.PHONE_PROVIDE
		}

		phone, err := vo.NewPhone(request.PhoneCode, request.PhoneNumber)

		if err != nil {
			return nil, err
		}

		identityUser.SetPhone(phone)
	}

	if request.ZipCode != "" {
		zipCode, err := vo.NewZipCode(request.ZipCode)

		if err != nil {
			return nil, err
		}

		identityUser.SetZipCode(zipCode)
	}

	if request.Street != "" {
		identityUser.SetStreet(request.Street)
	}

	if request.StreetPort != nil {
		identityUser.SetStreetNumber(*request.StreetPort)
	}

	if request.City != "" {
		identityUser.SetCity(request.City)
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

		identityUser.SetPicture(url)
	}

	_, err := services.GetAutarchy(identityUser.GetAddress())

	if err != nil {
		return nil, err
	}

	return identityUser, nil
}

func (uc *UpdateProfileUseCase) updateUser(foundAuth *entities.Auth, request contracts.UpdateProfileResponse) (interface{}, error) {
	foundProfile, err := uc.profileRepo.GetUserByAuthId(request.UserId)

	if err != nil {
		return nil, err
	}

	if request.UserName != "" {
		foundProfile.UserName = request.UserName
	}

	updatedProfile, err := uc.updateIdentityUser(foundProfile, foundAuth, request)

	if err != nil {
		return nil, err
	}

	if err := uc.profileRepo.Update(updatedProfile); err != nil {
		return nil, err
	}

	return contracts.GetProfileResponse(foundAuth, updatedProfile, uc.autarchyRepository)
}

func (uc *UpdateProfileUseCase) updateAutarchy(foundAuth *entities.Auth, request contracts.UpdateProfileResponse) (interface{}, error) {
	foundProfile, err := uc.profileRepo.GetAutarchyByAuthId(request.UserId)

	if err != nil {
		return nil, err
	}

	if request.Title != "" {
		foundProfile.Title = request.Title
	}

	if request.Lat != "" && request.Lon != "" {
		lat, err := strconv.ParseFloat(request.Lat, 64)

		if err != nil {
			return nil, err
		}

		lon, err := strconv.ParseFloat(request.Lon, 64)

		if err != nil {
			return nil, err
		}

		foundProfile.Coordinates = *vo.NewCoordinate(lat, lon)
	}

	updatedProfile, err := uc.updateIdentityUser(foundProfile, foundAuth, request)

	if err != nil {
		return nil, err
	}

	if err := uc.autarchyRepository.Update(updatedProfile); err != nil {
		return nil, err
	}

	return contracts.GetProfileResponse(foundAuth, updatedProfile, uc.autarchyRepository)
}

func (uc *UpdateProfileUseCase) Handle(request contracts.UpdateProfileResponse) (interface{}, error) {
	foundAuth, err := uc.authRepo.GetAuthById(request.UserId)

	if err != nil {
		return nil, exec.USER_NOT_FOUND
	}

	if request.Email != "" {
		email, err := vo.NewEmail(request.Email)

		if err != nil {
			return nil, err
		}

		if uc.authRepo.ExistsUserWithEmail(email) && foundAuth.Email != *email {
			return nil, exec.USER_ALREADY_EXISTS
		}

		foundAuth.Email = *email

		if err := uc.authRepo.Update(foundAuth); err != nil {
			return nil, err
		}
	}

	switch foundAuth.UserType {
	case int(vo.User), int(vo.Admin):
		return uc.updateUser(foundAuth, request)
	case int(vo.Autarchy):
		return uc.updateAutarchy(foundAuth, request)
	default:
		return uc.updateUser(foundAuth, request)
	}
}
