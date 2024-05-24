package usescases

import (
	"errors"
	"strconv"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/geojson"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type UpdateAutarchyUseCase struct {
	autarchyRepo *repositories.AutarchyRepository
	authRepo     *repositories.AuthRepository
}

func NewUpdateAutarchyUseCase(
	autarchyRepo *repositories.AutarchyRepository,
	authRepo *repositories.AuthRepository,
) *UpdateAutarchyUseCase {
	return &UpdateAutarchyUseCase{
		autarchyRepo: autarchyRepo,
		authRepo:     authRepo,
	}
}

func (uc *UpdateAutarchyUseCase) Handle(request contracts.UpdateAutarchyRequest) (*geojson.GeoJsonFeature, error) {
	foundAutarchy, err := uc.autarchyRepo.GetAutarchyById(request.AutarchyId)

	if err != nil {
		return nil, errors.New("autarchy not found")
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
			return nil, errors.New("provide an valid lat")
		}

		lon, err := strconv.ParseFloat(request.Lon, 64)

		if err != nil {
			return nil, errors.New("provide an valid lon")
		}

		foundAutarchy.Coordinates = *vo.NewCoordinate(lat, lon)
	}

	if err := uc.autarchyRepo.Update(foundAutarchy); err != nil {
		return nil, err
	}

	result, err := uc.autarchyRepo.GetAutarchtDetailById(request.AutarchyId)

	if err != nil {
		return nil, errors.New("autarchy could not be updated")
	}

	return geojson.NewFeature(
		result.Lat,
		result.Lon,
		contracts.AutarchyResponse{
			Id:          request.AutarchyId,
			Title:       result.Title,
			Email:       result.Email,
			PhoneCode:   result.PhoneNumber.CountryCode,
			PhoneNumber: result.PhoneNumber.Number,
			Street:      result.Address.Street,
			StreetPort:  result.Address.Number,
			ZipCode:     result.Address.ZipCode,
			City:        result.Address.City,
			Avatar:      result.AutarchyAvatar,
		},
	), nil
}
