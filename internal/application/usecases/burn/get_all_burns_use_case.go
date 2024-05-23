package usecases

import (
	"errors"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/date"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/geojson"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pagination"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
)

type GetAllBurnsUseCase struct {
	burnRepository *repositories.BurnRepository
}

func NewGetAllBurnsUseCase(burnRepository *repositories.BurnRepository) *GetAllBurnsUseCase {
	return &GetAllBurnsUseCase{
		burnRepository: burnRepository,
	}
}

func (uc *GetAllBurnsUseCase) Handle(request contracts.GetAllBurnsRequest) (*geojson.GeoJsonCollection, error) {
	var features = make([]geojson.GeoJsonFeature, 0)

	params := map[string]interface{}{
		"search": request.Search,
	}

	if request.State != "" {
		state, ok := vo.GetBurnRequestStateKey(request.State)

		if !ok {
			return nil, errors.New("state is not valid")
		}

		params["state"] = state
	}

	if request.StartDate != "" {
		startDate, err := date.ParseString(request.StartDate)

		if err != nil {
			return nil, err
		}

		params["start_date"] = startDate
	}

	if request.EndDate != "" {
		endDate, err := date.ParseString(request.EndDate)

		if err != nil {
			return nil, err
		}

		params["end_date"] = endDate
	}

	pag := pagination.Pagination{
		PageSize: request.PageSize,
		Page:     request.Page,
	}

	result, err := uc.burnRepository.GetAllBurns(request.AuthId, params, &pag)

	for _, v := range result {
		features = append(features, *geojson.NewFeature(
			v.Lat,
			v.Lon,
			contracts.BurnResponse{
				Id:          v.Id,
				Title:       v.Title,
				HasAidTeam:  v.HasAidTeam,
				Reason:      vo.MustGetBurnReason(v.Reason),
				Type:        vo.MustGetBurnType(v.Type),
				BeginAt:     v.BeginAt,
				CompletedAt: v.CompletedAt,
				Picture:     v.MapPicture,
				State:       vo.MustGetBurnRequestState(v.State),
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
