package usecases

import (
	"errors"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
)

type StartBurnUseCase struct {
	burnRepository *repositories.BurnRepository
}

func NewStartBurnUserCase(
	burnRepository *repositories.BurnRepository,
) *StartBurnUseCase {
	return &StartBurnUseCase{
		burnRepository: burnRepository,
	}
}

func (uc *StartBurnUseCase) Handle(request contracts.StartBurnRequest) (*contracts.BurnActionResponse, error) {
	foundBurn, err := uc.burnRepository.GetBurnDetailById(request.UserId, request.BurnId)

	if err != nil {
		return nil, errors.New("burn not found")
	}

	if foundBurn.State != uint16(vo.Scheduled) {
		return nil, errors.New("burn is not schedualed, it cannot perform this action")
	}

	state, err := uc.burnRepository.SetBurnStatus(
		request.UserId,
		foundBurn.Autarchy,
		foundBurn.Id,
		vo.Active,
	)

	if err != nil {
		return nil, errors.New("cannot start burn")
	}

	return &contracts.BurnActionResponse{
		BurnId: request.BurnId,
		State:  state.GetState(),
	}, nil
}
