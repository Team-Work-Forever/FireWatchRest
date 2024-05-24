package usescases

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/geojson"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pagination"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
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
		features = append(features, *geojson.NewFeature(
			v.Lat,
			v.Lon,
			contracts.AutarchyResponse{
				Title:       v.Title,
				Email:       v.Email,
				PhoneCode:   v.PhoneNumber.CountryCode,
				PhoneNumber: v.PhoneNumber.Number,
				Street:      v.Address.Street,
				StreetPort:  v.Address.Number,
				ZipCode:     v.Address.ZipCode,
				City:        v.Address.City,
				Lat:         v.Lat,
				Lon:         v.Lon,
				Avatar:      v.AutarchyAvatar,
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
