package usescases

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type DeleteAutarchyUseCase struct {
	autarchyRepo *repositories.AutarchyRepository
}

func NewDeleteAutarchyUseCase(autarchyRepo *repositories.AutarchyRepository) *DeleteAutarchyUseCase {
	return &DeleteAutarchyUseCase{
		autarchyRepo: autarchyRepo,
	}
}

func (uc *DeleteAutarchyUseCase) Handle(request contracts.DeleteAutarchyRequest) (*contracts.AutarchyActionResponse, error) {
	foundAutarchy, err := uc.autarchyRepo.GetAutarchyById(request.AutarchyId)

	if err != nil {
		return nil, exec.AUTARCHY_NOT_FOUND
	}

	// remove it
	if err := uc.autarchyRepo.Delete(foundAutarchy); err != nil {
		return nil, exec.AUTARCHY_FAILED_REMOVAL
	}

	return &contracts.AutarchyActionResponse{
		AutarchyId: foundAutarchy.ID,
	}, nil
}
