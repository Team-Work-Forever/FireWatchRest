package usescases

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/geojson"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pagination"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type GetAllAutarchies struct {
	autarchyRepo *repositories.AutarchyRepository
}

func NewGetAllAutarchies(autarchyRepo *repositories.AutarchyRepository) *GetAllAutarchies {
	return &GetAllAutarchies{
		autarchyRepo: autarchyRepo,
	}
}

func (uc *GetAllAutarchies) Handle(request contracts.GetAllAutarchiesRequest) (*geojson.GeoJsonCollection, error) {
	var features = make([]geojson.GeoJsonFeature, 0)

	params := map[string]interface{}{
		"search": request.Search,
	}

	pag := pagination.Pagination{
		PageSize: request.PageSize,
		Page:     request.Page,
	}

	result, err := uc.autarchyRepo.GetAll(params, &pag)

	for _, v := range result {
		totalofBurns, err := uc.autarchyRepo.GetAutarchyBurnCount(v.Id)

		if err != nil {
			return nil, exec.AUTARCHY_FAILED_DETAILS_FETCH
		}

		features = append(features, *geojson.NewFeature(
			v.Lat,
			v.Lon,
			contracts.AutarchyResponse{
				Id:    v.Id,
				Title: v.Title,
				Email: v.Email,
				NIF:   v.NIF,
				Phone: contracts.PhoneResponse{
					CountryCode: v.PhoneNumber.CountryCode,
					Number:      v.PhoneNumber.Number,
				},
				Address: contracts.AddressResponse{
					Street: v.Address.Street,
					Number: v.Address.Number,
					ZipCode: contracts.ZipCodeResponse{
						Value: v.Address.ZipCode,
					},
					City: v.Address.City,
				},
				Avatar:       v.AutarchyAvatar,
				TotalOfBurns: totalofBurns,
			},
		))
	}

	if err != nil {
		return nil, err
	}

	return geojson.NewCollection(
		features,
		pag,
	), nil
}
