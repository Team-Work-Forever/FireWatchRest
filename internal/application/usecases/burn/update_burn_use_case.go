package usecases

import (
	"errors"
	"strconv"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/date"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/geojson"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
)

type UpdateBurnUseCase struct {
	burnRepository *repositories.BurnRepository
}

func NewUpdateBurnUseCase(burnRepository *repositories.BurnRepository) *UpdateBurnUseCase {
	return &UpdateBurnUseCase{
		burnRepository: burnRepository,
	}
}

func (uc *UpdateBurnUseCase) Handle(request contracts.UpdateBurnRequest) (*geojson.GeoJsonFeature, error) {
	foundBurn, err := uc.burnRepository.GetBurnById(request.BurnId)

	if err != nil {
		return nil, errors.New("burn not found")
	}

	if ok := uc.burnRepository.UserOwnsBurn(request.UserId, request.BurnId); !ok {
		return nil, errors.New("you don't have access to this burn")
	}

	if request.Title != "" {
		foundBurn.Title = request.Title
	}

	if request.Type != "" {
		burnType, ok := vo.GetBurnTypeKey(request.Type)

		if !ok {
			return nil, errors.New("burn type does not exists")
		}

		foundBurn.Type = burnType
	}

	if request.HasBackUpTeam != "" {
		hasBackUpTeam, err := strconv.ParseBool(request.HasBackUpTeam)

		if err != nil {
			return nil, errors.New("provide an valid boolean")
		}

		foundBurn.HasAidTeam = hasBackUpTeam
	}

	if request.Reason != "" {
		reason, ok := vo.GetBurnReasonKey(request.Reason)

		if !ok {
			return nil, errors.New("provide an valid reason")
		}

		foundBurn.Reason = reason
	}

	if request.InitDate != "" {
		initDate, err := date.ParseString(request.InitDate)

		if err != nil {
			return nil, errors.New("provide an valid date yyyy-mm-dd")
		}

		foundBurn.BeginAt = *initDate
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

		foundBurn.Coordinates = *vo.NewCoordinate(float32(lat), float32(lon))
	}

	if err := uc.burnRepository.Update(foundBurn); err != nil {
		return nil, err
	}

	result, err := uc.burnRepository.GetBurnDetailById(request.UserId, request.BurnId)

	if err != nil {
		return nil, errors.New("burn could not be updated")
	}

	return geojson.NewFeature(
		result.Lat,
		result.Lon,
		contracts.BurnResponse{
			Id:          result.Id,
			Title:       result.Title,
			HasAidTeam:  result.HasAidTeam,
			Reason:      vo.MustGetBurnReason(result.Reason),
			Type:        vo.MustGetBurnType(result.Type),
			BeginAt:     result.BeginAt,
			CompletedAt: result.CompletedAt,
			Picture:     result.MapPicture,
			State:       vo.MustGetBurnRequestState(result.State),
		},
	), nil
}
