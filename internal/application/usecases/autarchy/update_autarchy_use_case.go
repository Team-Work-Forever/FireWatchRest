package usescases

import (
	"strconv"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/geojson"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/services"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/upload"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type UpdateAutarchyUseCase struct {
	autarchyRepo *repositories.AutarchyRepository
	authRepo     *repositories.AuthRepository
	fileService  *upload.BlobService
}

func NewUpdateAutarchyUseCase(
	autarchyRepo *repositories.AutarchyRepository,
	authRepo *repositories.AuthRepository,
	fileService *upload.BlobService,
) *UpdateAutarchyUseCase {
	return &UpdateAutarchyUseCase{
		autarchyRepo: autarchyRepo,
		authRepo:     authRepo,
		fileService:  fileService,
	}
}

func (uc *UpdateAutarchyUseCase) Handle(request contracts.UpdateAutarchyRequest) (*geojson.GeoJsonFeature, error) {
	foundAutarchy, err := uc.autarchyRepo.GetAutarchyById(request.AutarchyId)

	if err != nil {
		return nil, exec.AUTARCHY_NOT_FOUND
	}

	foundAuth, err := uc.authRepo.GetAuthById(foundAutarchy.AuthKeyId)

	if err != nil {
		return nil, exec.USER_NOT_FOUND
	}

	if request.Title != "" {
		foundAutarchy.Title = request.Title
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

	if request.PhoneCode != "" || request.PhoneNumber != "" {
		if request.PhoneCode == "" || request.PhoneNumber == "" {
			return nil, exec.PHONE_PROVIDE
		}

		phone, err := vo.NewPhone(request.PhoneCode, request.PhoneNumber)
		if err != nil {
			return nil, err
		}

		foundAutarchy.PhoneNumber = *phone
	}

	if request.ZipCode != "" {
		zipCode, err := vo.NewZipCode(request.ZipCode)

		if err != nil {
			return nil, err
		}

		foundAutarchy.Address.ZipCode = zipCode.GetValue()
	}

	if request.Street != "" {
		foundAutarchy.Address.Street = request.Street
	}

	if request.StreetPort != nil {
		foundAutarchy.Address.SetStreetNumber(*request.StreetPort)
	}

	if request.City != "" {
		foundAutarchy.Address.City = request.City
	}

	if request.Lat != "" && request.Lon != "" {
		lat, err := strconv.ParseFloat(request.Lat, 64)

		if err != nil {
			return nil, exec.AUTARCHY_PROVIDE_LAT
		}

		lon, err := strconv.ParseFloat(request.Lon, 64)

		if err != nil {
			return nil, exec.AUTARCHY_PROVIDE_LON
		}

		foundAutarchy.Coordinates = *vo.NewCoordinate(lat, lon)
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

		foundAutarchy.Picture = url
	}

	_, err = services.GetAutarchy(foundAutarchy.Address)

	if err != nil {
		return nil, err
	}

	if err := uc.autarchyRepo.Update(foundAutarchy); err != nil {
		return nil, err
	}

	result, err := uc.autarchyRepo.GetAutarchtDetailById(request.AutarchyId)

	if err != nil {
		return nil, exec.AUTARCHY_NOT_ABLE_UPDATE
	}

	totalofBurns, err := uc.autarchyRepo.GetAutarchyBurnCount(result.Id)

	if err != nil {
		return nil, exec.AUTARCHY_FAILED_DETAILS_FETCH
	}

	return geojson.NewFeature(
		result.Lat,
		result.Lon,
		contracts.AutarchyResponse{
			Id:    request.AutarchyId,
			Title: result.Title,
			Email: result.Email,
			NIF:   result.NIF,
			Phone: contracts.PhoneResponse{
				CountryCode: result.PhoneNumber.CountryCode,
				Number:      result.PhoneNumber.Number,
			},
			Address: contracts.AddressResponse{
				Street: result.Address.Street,
				Number: result.Address.Number,
				ZipCode: contracts.ZipCodeResponse{
					Value: result.Address.ZipCode,
				},
				City: result.Address.City,
			},
			Avatar:       result.AutarchyAvatar,
			TotalOfBurns: totalofBurns,
		},
	), nil
}
