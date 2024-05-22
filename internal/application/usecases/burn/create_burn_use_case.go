package burn

import (
	"errors"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/daos"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/date"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
)

type CreateBurnUseCase struct {
	burnRepo *repositories.BurnRepository
}

func NewCreateBurnUseCase(burnRepo *repositories.BurnRepository) *CreateBurnUseCase {
	return &CreateBurnUseCase{
		burnRepo: burnRepo,
	}
}

func (uc *CreateBurnUseCase) Handler(request contracts.CreateBurnRequest) (*contracts.BurnActionResponse, error) {
	// validate
	initDate, err := date.ParseString(request.InitDate)

	if err != nil {
		return nil, err
	}

	reason, ok := vo.GetBurnReasonKey(request.Reason)

	if !ok {
		return nil, errors.New("reason type does not exists")
	}

	burnType, ok := vo.GetBurnTypeKey(request.Type)

	if !ok {
		return nil, errors.New("burn type does not exists")
	}

	// call api to check if possible to create an burn

	// create burn
	burn, err := entities.NewBurn(
		request.Title,
		reason,
		burnType,
		*vo.NewCoordinate(request.Lat, request.Lon),
		*initDate,
	)

	if err != nil {
		return nil, err
	}

	burnRequest, err := uc.burnRepo.CreateBurn(daos.CreateBurnDao{
		AuthId:         request.UserId,
		Burn:           burn,
		InitialPropose: request.InitialProprose,
	})

	if err != nil {
		return nil, err
	}

	return &contracts.BurnActionResponse{
		BurnId: burnRequest.BurnId,
		State:  burnRequest.State.GetState(),
	}, nil
}
