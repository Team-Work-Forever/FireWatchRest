package usecases

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/date"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/geojson"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type GetAllBurnsUseCase struct {
	burnRepository *repositories.BurnRepository
	autarchyRepo   *repositories.AutarchyRepository
	authRepo       *repositories.AuthRepository
	profileRepo    *repositories.ProfileRepository
}

func NewGetAllBurnsUseCase(
	burnRepository *repositories.BurnRepository,
	autarchyRepo *repositories.AutarchyRepository,
	authRepo *repositories.AuthRepository,
	profileRepo *repositories.ProfileRepository,
) *GetAllBurnsUseCase {
	return &GetAllBurnsUseCase{
		burnRepository: burnRepository,
		autarchyRepo:   autarchyRepo,
		authRepo:       authRepo,
		profileRepo:    profileRepo,
	}
}

func (uc *GetAllBurnsUseCase) Handle(request contracts.GetAllBurnsRequest) (*geojson.GeoJsonCollection, error) {
	var features = make([]geojson.GeoJsonFeature, 0)

	params := map[string]interface{}{
		"search": request.Search,
		"sort":   request.Sort,
	}

	foundProfile, err := uc.authRepo.GetAuthById(request.AuthId)

	if err != nil {
		return nil, exec.USER_NOT_FOUND
	}

	params["userType"] = foundProfile.UserType

	if request.Sort != "" {
		if request.Sort != "asc" && request.Sort != "desc" {
			return nil, exec.QUERY_PARAMETER_SORT_NOT_VALID
		}
	}

	if request.AutarchyId != "" {
		if foundProfile.UserType == int(vo.Autarchy) {
			foundAutarchy, err := uc.profileRepo.GetAutarchyByAuthId(request.AuthId)

			if err != nil {
				return nil, err
			}

			request.AutarchyId = foundAutarchy.ID
		} else {
			if _, err := uc.autarchyRepo.GetAutarchyById(request.AutarchyId); err != nil {
				return nil, exec.AUTARCHY_NOT_FOUND
			}
		}

		params["autarchyId"] = request.AutarchyId
	}

	if request.State != "" {
		state, ok := vo.GetBurnRequestStateKey(request.State)

		if !ok {
			return nil, exec.STATE_NOT_VALID
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

	result, err := uc.burnRepository.GetAllBurns(request.AuthId, params, request.Pagination)

	for _, v := range result {
		features = append(features, *geojson.NewFeature(
			v.Lat,
			v.Lon,
			contracts.BurnResponse{
				Id:    v.Id,
				Title: v.Title,
				Author: contracts.PublicProfileResponse{
					Email:    v.Profile.Email,
					UserName: v.Profile.UserName,
					Avatar:   v.Profile.ProfileAvatar,
					NIF:      v.Profile.NIF,
					Phone: contracts.PhoneResponse{
						CountryCode: v.Profile.Phone.CountryCode,
						Number:      v.Profile.Phone.Number,
					},
				},
				HasAidTeam: v.HasAidTeam,
				Reason:     vo.MustGetBurnReason(v.Reason),
				Type:       vo.MustGetBurnType(v.Type),
				Address: contracts.AddressResponse{
					Street: v.Street,
					Number: v.Number,
					ZipCode: contracts.ZipCodeResponse{
						Value: v.ZipCode,
					},
					City: v.City,
				},
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
		*request.Pagination,
	), nil
}
