package usecases

import (
	"errors"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
)

type DeleteBurnUseCase struct {
	burnRepository *repositories.BurnRepository
}

func NewDeleteBurnUseCase(burnRepository *repositories.BurnRepository) *DeleteBurnUseCase {
	return &DeleteBurnUseCase{
		burnRepository: burnRepository,
	}
}

func (uc *DeleteBurnUseCase) Handle(request contracts.DeleteRequest) (*contracts.BurnActionResponse, error) {
	// get burn
	foundBurn, err := uc.burnRepository.GetBurnById(request.BurnId)

	if err != nil {
		return nil, errors.New("burn not found")
	}

	if ok := uc.burnRepository.UserOwnsBurn(request.UserId, request.BurnId); !ok {
		return nil, errors.New("you don't have access to this burn")
	}

	// verify status
	status, err := uc.burnRepository.GetBurnStatus(request.UserId, request.BurnId)

	if err != nil {
		return nil, errors.New("could not fetch the status")
	}

	if *status != uint16(vo.Scheduled) {
		return nil, errors.New("is not possible to remove an burn in action")
	}

	// remove it
	if err := uc.burnRepository.Delete(foundBurn); err != nil {
		return nil, errors.New("the burn failed to be removed")
	}

	return &contracts.BurnActionResponse{
		BurnId: foundBurn.ID,
		State:  "deleted",
	}, nil
}
