package usescases

import (
	"errors"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/geojson"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
)

type GetAutarchyByIdUseCase struct {
	autarchyRepo *repositories.AutarchyRepository
}

func NewGetAutarchyByIdUseCase(autarchyRepo *repositories.AutarchyRepository) *GetAutarchyByIdUseCase {
	return &GetAutarchyByIdUseCase{
		autarchyRepo: autarchyRepo,
	}
}

func (uc *GetAutarchyByIdUseCase) Handle(request contracts.GetAutarchyRequest) (*geojson.GeoJsonFeature, error) {
	result, err := uc.autarchyRepo.GetAutarchtDetailById(request.AutarchyId)

	if err != nil {
		return nil, errors.New("autarchy not found")
	}

	return geojson.NewFeature(
		result.Lat,
		result.Lon,
		contracts.AutarchyResponse{
			Id:          result.Id,
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
