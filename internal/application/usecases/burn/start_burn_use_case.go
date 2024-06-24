package usecases

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
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
		return nil, exec.BURN_NOT_FOUND
	}

	if foundBurn.State != uint16(vo.Scheduled) {
		return nil, exec.BURN_FAILED_SCHEDULED_ACTION
	}

	state, err := uc.burnRepository.SetBurnStatus(
		request.UserId,
		foundBurn.Autarchy,
		foundBurn.Id,
		vo.Active,
	)

	if err != nil {
		return nil, exec.BURN_NOT_ABLE_START
	}

	return &contracts.BurnActionResponse{
		BurnId: request.BurnId,
		State:  state.GetState(),
	}, nil
}
