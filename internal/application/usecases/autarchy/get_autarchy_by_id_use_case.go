package usescases

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/geojson"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
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
		return nil, exec.AUTARCHY_NOT_FOUND
	}

	totalofBurns, err := uc.autarchyRepo.GetAutarchyBurnCount(result.Id)

	if err != nil {
		return nil, exec.AUTARCHY_FAILED_DETAILS_FETCH
	}

	return geojson.NewFeature(
		result.Lat,
		result.Lon,
		contracts.AutarchyResponse{
			Id:    result.Id,
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
