package usescases

import (
	"errors"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type CreateAutarchyUseCase struct {
	autarchyRepo *repositories.AutarchyRepository
	authRepo     *repositories.AuthRepository
}

func NewCreateAutarchyUseCase(
	autarchyRepo *repositories.AutarchyRepository,
	authRepo *repositories.AuthRepository,
) *CreateAutarchyUseCase {
	return &CreateAutarchyUseCase{
		autarchyRepo: autarchyRepo,
		authRepo:     authRepo,
	}
}

func (uc *CreateAutarchyUseCase) Handle(request contracts.CreateAutarchyRequest) (*contracts.AutarchyActionResponse, error) {
	email, err := vo.NewEmail(request.Email)

	if err != nil {
		return nil, err
	}

	if request.Title == "" {
		return nil, errors.New("provide an name for the autarchy")
	}

	password, err := vo.NewPassword(request.Password)

	if err != nil {
		return nil, err
	}

	nif, err := vo.NewNIF(request.NIF)

	if err != nil {
		return nil, err
	}

	phone, err := vo.NewPhone(request.PhoneCode, request.PhoneNumber)

	if err != nil {
		return nil, err
	}

	zipCode, err := vo.NewZipCode(request.ZipCode)

	if err != nil {
		return nil, err
	}

	address, err := vo.NewAddress(request.Street, request.StreetPort, *zipCode, request.City)

	if err != nil {
		return nil, err
	}

	// email, title, phone
	if ok := uc.authRepo.ExistsUserWithEmail(email); ok {
		return nil, exec.USER_ALREADY_EXISTS
	}

	if ok := uc.authRepo.ExistsUserWithNif(nif); ok {
		return nil, exec.USER_ALREADY_EXISTS_NIF
	}

	if ok := uc.autarchyRepo.ExistsAutarchyWithTitle(request.Title); ok {
		return nil, errors.New("that autarchy is already registered")
	}

	// create and store
	auth := entities.NewAuth(
		*email,
		*password,
		*nif,
		int(vo.Autarchy),
	)

	autarchy, err := entities.NewAutarchy(
		request.Title,
		*vo.NewCoordinate(request.Lat, request.Lon),
		*phone,
		*address,
	)

	if err != nil {
		return nil, err
	}

	if err := uc.authRepo.CreateAccount(auth, autarchy); err != nil {
		return nil, err
	}

	return &contracts.AutarchyActionResponse{
		AutarchyId: autarchy.ID,
	}, nil
}
