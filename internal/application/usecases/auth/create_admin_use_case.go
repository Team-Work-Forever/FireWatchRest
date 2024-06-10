package usecases

import (
	"fmt"
	"log"
	"os"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pwd"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/smtp"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/upload"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
)

type CreateAdminUseCase struct {
	authRepo    *repositories.AuthRepository
	fileService *upload.BlobService
}

func NewCreateAdminUseCase(
	authRepo *repositories.AuthRepository,
	fileService *upload.BlobService,
) *CreateAdminUseCase {
	return &CreateAdminUseCase{
		authRepo:    authRepo,
		fileService: fileService,
	}
}

const (
	AdminEmail           = "admin@firewatch.fire"
	AdminNIF             = "000000000"
	AdminUserName        = "administrator"
	AdminFirstName       = "super"
	AdminLastName        = "administrator"
	AdminICFNphoneCode   = "+351"
	AdminICFNphoneNumber = "213507900"

	AdminICFNStreet     = "Avenida da República"
	AdminICFNPortNumber = 16
	AdminICFNZipCode    = "1050-191"
	AdminICFNCity       = "Lisboa"
)

func (uc *CreateAdminUseCase) Handle(request contracts.CreateAdminRequest) (*contracts.CreateAdminResponse, error) {
	email, err := vo.NewEmail(AdminEmail)

	if err != nil {
		return nil, err
	}

	if uc.authRepo.ExistsUserWithEmail(email) {
		return nil, nil
	}

	nif, err := vo.NewNIF(AdminNIF)

	if err != nil {
		return nil, err
	}

	phone, err := vo.NewPhone(AdminICFNphoneCode, AdminICFNphoneNumber)

	if err != nil {
		return nil, err
	}

	zipCode, err := vo.NewZipCode(AdminICFNZipCode)

	if err != nil {
		return nil, err
	}

	address, err := vo.NewAddress(
		AdminICFNStreet,
		AdminICFNPortNumber,
		*zipCode,
		AdminICFNCity,
	)

	if err != nil {
		return nil, err
	}

	password, err := pwd.GeneratePasswordFixed()

	if err != nil {
		return nil, err
	}

	// create auth
	auth := entities.NewAuth(
		*email,
		*password,
		*nif,
		int(vo.Admin),
	)

	// Fetch ADMIN AVATAR PICTURE
	file, err := os.Open(request.AvatarFilePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	url, err := uc.fileService.UploadFile(&upload.UploadFile{
		Bucket:   upload.ClientBucket,
		FileName: file.Name(),
		FileId:   auth.GetId(),
		FileBody: file,
	})

	if err != nil {
		return nil, err
	}

	user := entities.NewUser(
		url,
		AdminUserName,
		AdminFirstName,
		AdminLastName,
		*phone,
		*address,
	)

	if err := uc.authRepo.CreateAccount(auth, user); err != nil {
		return nil, err
	}

	if request.SendItByEmail {
		mail := smtp.New(
			email.GetValue(),
			"System Admin Password",
			fmt.Sprintf("Secret: %s", password.GetValue()),
		)

		sendMail := func() {
			if err := mail.Send(); err != nil {
				log.Print(err)
			}
		}

		go sendMail()
	}

	return &contracts.CreateAdminResponse{
		GeneratedPassword: password.GetValue(),
	}, nil
}
