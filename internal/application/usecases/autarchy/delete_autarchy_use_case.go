package usescases

import (
	"errors"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
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
		return nil, errors.New("autarchy not found")
	}

	// remove it
	if err := uc.autarchyRepo.Delete(foundAutarchy); err != nil {
		return nil, errors.New("the autarchy failed to be removed")
	}

	return &contracts.AutarchyActionResponse{
		AutarchyId: foundAutarchy.ID,
	}, nil
}
