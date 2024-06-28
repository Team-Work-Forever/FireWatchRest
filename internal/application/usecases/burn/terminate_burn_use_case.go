package usecases

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
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
	foundBurn, err := uc.burnRepository.GetBurnDetailById(request.UserId, request.BurnId, false)

	if err != nil {
		return nil, exec.BURN_NOT_FOUND
	}

	if foundBurn.State != uint16(vo.Active) {
		return nil, exec.BURN_NOT_ACTIVE
	}

	state, err := uc.burnRepository.SetBurnStatus(
		request.UserId,
		foundBurn.Autarchy,
		foundBurn.Id,
		vo.Completed,
	)

	if err != nil {
		return nil, exec.BURN_NOT_ABLE_COMPLETE
	}

	return &contracts.BurnActionResponse{
		BurnId: request.BurnId,
		State:  state.GetState(),
	}, nil
}
