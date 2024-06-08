package usecases

import (
	"errors"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/geojson"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
)

type GetBurnByIdUseCase struct {
	burnRepository *repositories.BurnRepository
}

func NewGetBurnByIdUseCase(burnRepository *repositories.BurnRepository) *GetBurnByIdUseCase {
	return &GetBurnByIdUseCase{
		burnRepository: burnRepository,
	}
}

func (uc *GetBurnByIdUseCase) Handle(request contracts.GetBurnRequest) (*geojson.GeoJsonFeature, error) {
	result, err := uc.burnRepository.GetBurnDetailById(request.AuthId, request.BurnId)

	if err != nil {
		return nil, errors.New("burn not found")
	}

	return geojson.NewFeature(
		result.Lat,
		result.Lon,
		contracts.BurnResponse{
			Id:         result.Id,
			Title:      result.Title,
			HasAidTeam: result.HasAidTeam,
			Reason:     vo.MustGetBurnReason(result.Reason),
			Type:       vo.MustGetBurnType(result.Type),
			Address: contracts.AddressResponse{
				Street: result.Street,
				Number: result.Number,
				ZipCode: contracts.ZipCodeResponse{
					Value: result.ZipCode,
				},
				City: result.City,
			},
			BeginAt:     result.BeginAt,
			CompletedAt: result.CompletedAt,
			Picture:     result.MapPicture,
			State:       vo.MustGetBurnRequestState(result.State),
		},
	), nil
}
