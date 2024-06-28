package usescases

import (
	"fmt"
	"log"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pwd"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/services"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/smtp"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/upload"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type CreateAutarchyUseCase struct {
	autarchyRepo *repositories.AutarchyRepository
	authRepo     *repositories.AuthRepository
	fileService  *upload.BlobService
}

func NewCreateAutarchyUseCase(
	autarchyRepo *repositories.AutarchyRepository,
	authRepo *repositories.AuthRepository,
	fileService *upload.BlobService,
) *CreateAutarchyUseCase {
	return &CreateAutarchyUseCase{
		autarchyRepo: autarchyRepo,
		authRepo:     authRepo,
		fileService:  fileService,
	}
}

func (uc *CreateAutarchyUseCase) Handle(request contracts.CreateAutarchyRequest) (*contracts.AutarchyActionResponse, error) {
	email, err := vo.NewEmail(request.Email)

	if err != nil {
		return nil, err
	}

	if request.Title == "" {
		return nil, exec.AUTARCHY_PROVIDE_NAME
	}

	password, err := pwd.GeneratePasswordFixed()

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

	_, err = services.GetAutarchy(*address)

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
		return nil, exec.AUTARCHY_ALREADY_REGISTERED
	}

	auth := entities.NewAuth(
		*email,
		*password,
		*nif,
		int(vo.Autarchy),
	)

	file, err := request.Avatar.Open()

	if err != nil {
		return nil, err
	}

	defer file.Close()

	url, err := uc.fileService.UploadFile(&upload.UploadFile{
		Bucket:   upload.ClientBucket,
		FileName: request.Avatar.Filename,
		FileId:   auth.GetId(),
		FileBody: file,
	})

	if err != nil {
		return nil, err
	}

	autarchy, err := entities.NewAutarchy(
		request.Title,
		url,
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

	mail := smtp.New(
		email.GetValue(),
		"System Autarchy Password",
		fmt.Sprintf("Secret: %s", password.GetValue()),
	)

	sendMail := func() {
		if err := mail.Send(); err != nil {
			log.Print(err)
		}
	}

	go sendMail()
	log.Printf("Autarchy Secret Generated %s", password)

	return &contracts.AutarchyActionResponse{
		AutarchyId: autarchy.ID,
	}, nil
}
