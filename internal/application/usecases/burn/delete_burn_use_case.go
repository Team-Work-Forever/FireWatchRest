package usecases

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type DeleteBurnUseCase struct {
	burnRepository *repositories.BurnRepository
}

func NewDeleteBurnUseCase(burnRepository *repositories.BurnRepository) *DeleteBurnUseCase {
	return &DeleteBurnUseCase{
		burnRepository: burnRepository,
	}
}

func (uc *DeleteBurnUseCase) Handle(request contracts.DeleteBurnRequest) (*contracts.BurnActionResponse, error) {
	// get burn
	foundBurn, err := uc.burnRepository.GetBurnById(request.BurnId)

	if err != nil {
		return nil, exec.BURN_NOT_FOUND
	}

	if ok := uc.burnRepository.UserOwnsBurn(request.UserId, request.BurnId); !ok {
		return nil, exec.BURN_DENIAL_OF_ACCESS
	}

	// verify status
	status, err := uc.burnRepository.GetBurnStatus(request.UserId, request.BurnId)

	if err != nil {
		return nil, exec.BURN_FAILED_FETCH
	}

	if *status != uint16(vo.Scheduled) {
		return nil, exec.BURN_NOT_ABLE_REMOVAL
	}

	// remove it
	if err := uc.burnRepository.Delete(foundBurn); err != nil {
		return nil, exec.BURN_FAILED_REMOVAL
	}

	return &contracts.BurnActionResponse{
		BurnId: foundBurn.ID,
		State:  "deleted",
	}, nil
}
