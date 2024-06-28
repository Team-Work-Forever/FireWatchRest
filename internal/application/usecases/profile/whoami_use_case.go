package usecases

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type WhoamiUseCase struct {
	authRepository     *repositories.AuthRepository
	profileRepository  *repositories.ProfileRepository
	autarchyRepository *repositories.AutarchyRepository
}

func NewWhoamiUseCase(
	authRepository *repositories.AuthRepository,
	profileRepository *repositories.ProfileRepository,
	autarchyRepository *repositories.AutarchyRepository,
) *WhoamiUseCase {
	return &WhoamiUseCase{
		authRepository:     authRepository,
		profileRepository:  profileRepository,
		autarchyRepository: autarchyRepository,
	}
}

func (w *WhoamiUseCase) fetchUser(auth *entities.Auth) (interface{}, error) {
	profileFound, err := w.profileRepository.GetUserByAuthId(auth.ID)

	if err != nil {
		return nil, exec.USER_NOT_FOUND
	}

	return contracts.GetProfileResponse(auth, profileFound, w.autarchyRepository)
}

func (w *WhoamiUseCase) fetchAutarchyResponse(auth *entities.Auth) (interface{}, error) {
	profileFound, err := w.profileRepository.GetAutarchyByAuthId(auth.ID)

	if err != nil {
		return nil, exec.USER_NOT_FOUND
	}

	return contracts.GetProfileResponse(auth, profileFound, w.autarchyRepository)
}

func (w *WhoamiUseCase) Handle(request contracts.WhoamiRequest) (interface{}, error) {
	foundAuth, err := w.authRepository.GetAuthById(request.UserId)

	if err != nil {
		return nil, exec.USER_NOT_FOUND
	}

	switch foundAuth.UserType {
	case int(vo.User), int(vo.Admin):
		return w.fetchUser(foundAuth)
	case int(vo.Autarchy):
		return w.fetchAutarchyResponse(foundAuth)
	}

	return nil, exec.USER_NOT_FOUND
}
