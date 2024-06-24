package usecases

import (
	"strconv"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/date"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/geojson"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
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
		return nil, exec.BURN_NOT_FOUND
	}

	if ok := uc.burnRepository.UserOwnsBurn(request.UserId, request.BurnId); !ok {
		return nil, exec.BURN_DENIAL_OF_ACCESS
	}

	if request.Title != "" {
		foundBurn.Title = request.Title
	}

	if request.Type != "" {
		burnType, ok := vo.GetBurnTypeKey(request.Type)

		if !ok {
			return nil, exec.BURN_PROVIDE_NOT_EXISTING_TYPE
		}

		foundBurn.Type = burnType
	}

	if request.HasBackUpTeam != "" {
		hasBackUpTeam, err := strconv.ParseBool(request.HasBackUpTeam)

		if err != nil {
			return nil, exec.BOOLEAN_PROVIDE_AN_VALID
		}

		foundBurn.HasAidTeam = hasBackUpTeam
	}

	if request.Reason != "" {
		reason, ok := vo.GetBurnReasonKey(request.Reason)

		if !ok {
			return nil, exec.BURN_PROVIDE_NOT_EXISTING_REASON
		}

		foundBurn.Reason = reason
	}

	if request.InitDate != "" {
		initDate, err := date.ParseString(request.InitDate)

		if err != nil {
			return nil, exec.DATE_PROVIDE_AN_VALID
		}

		foundBurn.BeginAt = *initDate
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

		foundBurn.Coordinates = *vo.NewCoordinate(lat, lon)
	}

	if err := uc.burnRepository.Update(foundBurn); err != nil {
		return nil, err
	}

	result, err := uc.burnRepository.GetBurnDetailById(request.UserId, request.BurnId)

	if err != nil {
		return nil, exec.BURN_NOT_ABLE_UPDATE
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
