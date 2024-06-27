package usecases

import (
	"errors"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/geojson"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type GetBurnByIdUseCase struct {
	burnRepository *repositories.BurnRepository
	authRepository *repositories.AuthRepository
}

func NewGetBurnByIdUseCase(
	burnRepository *repositories.BurnRepository,
	authRepository *repositories.AuthRepository,
) *GetBurnByIdUseCase {
	return &GetBurnByIdUseCase{
		burnRepository: burnRepository,
		authRepository: authRepository,
	}
}

func (uc *GetBurnByIdUseCase) Handle(request contracts.GetBurnRequest) (*geojson.GeoJsonFeature, error) {
	var isAdmin bool = false
	foundAuth, err := uc.authRepository.GetAuthById(request.AuthId)

	if err != nil {
		return nil, exec.USER_ALREADY_EXISTS
	}

	if foundAuth.UserType == int(vo.Admin) || foundAuth.UserType == int(vo.Autarchy) {
		isAdmin = true
	}

	result, err := uc.burnRepository.GetBurnDetailById(request.AuthId, request.BurnId, isAdmin)

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
			Author: contracts.PublicProfileResponse{
				Email:    result.Profile.Email,
				UserName: result.Profile.UserName,
				Avatar:   result.Profile.ProfileAvatar,
				NIF:      result.Profile.NIF,
				Phone: contracts.PhoneResponse{
					CountryCode: result.Profile.Phone.CountryCode,
					Number:      result.Profile.Phone.Number,
				},
			},
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
