package usecases

import (
	"errors"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
)

type TerminateBurnUseCase struct {
	burnRepository *repositories.BurnRepository
}

func NewTerminateBurnUseCase(
	burnRepository *repositories.BurnRepository,
) *TerminateBurnUseCase {
	return &TerminateBurnUseCase{
		burnRepository: burnRepository,
	}
}

func (uc *TerminateBurnUseCase) Handle(request contracts.TerminateBurnRequest) (*contracts.BurnActionResponse, error) {
	foundBurn, err := uc.burnRepository.GetBurnDetailById(request.UserId, request.BurnId)

	if err != nil {
		return nil, errors.New("burn not found")
	}

	if foundBurn.State != uint16(vo.Active) {
		return nil, errors.New("burn is not active, it cannot perform this action")
	}

	state, err := uc.burnRepository.SetBurnStatus(
		request.UserId,
		foundBurn.Autarchy,
		foundBurn.Id,
		vo.Completed,
	)

	if err != nil {
		return nil, errors.New("cannot complete burn")
	}

	return &contracts.BurnActionResponse{
		BurnId: request.BurnId,
		State:  state.GetState(),
	}, nil
}
