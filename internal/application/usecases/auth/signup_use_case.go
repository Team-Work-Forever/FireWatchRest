package usecases

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/jwt"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/upload"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type SignUpUseCase struct {
	authRepository *repositories.AuthRepository
	fileService    *upload.BlobService
}

func NewSignUpUseCase(
	authRepository *repositories.AuthRepository,
	fileService *upload.BlobService,
) *SignUpUseCase {
	return &SignUpUseCase{
		authRepository: authRepository,
		fileService:    fileService,
	}
}

func (uc *SignUpUseCase) Handle(request contracts.SignUpRequest) (*contracts.AuthResponse, error) {
	email, err := vo.NewEmail(request.Email)

	if err != nil {
		return nil, err
	}

	password, err := vo.NewPassword(request.Password)

	if err != nil {
		return nil, err
	}

	nif, err := vo.NewNIF(request.NIF)

	if err != nil {
		return nil, err
	}

	if request.UserName == "" {
		return nil, exec.USER_NAME_PROVIDE
	}

	if request.FirstName == "" {
		return nil, exec.FIRST_NAME_PROVIDE
	}

	if request.LastName == "" {
		return nil, exec.LAST_NAME_PROVIDE
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

	// _, err = services.GetAutarchy(*address)

	// if err != nil {
	// 	return nil, err
	// }

	if ok := uc.authRepository.ExistsUserWithEmail(email); ok {
		return nil, exec.USER_ALREADY_EXISTS
	}

	if ok := uc.authRepository.ExistsUserWithNif(nif); ok {
		return nil, exec.USER_ALREADY_EXISTS_NIF
	}

	auth := entities.NewAuth(
		*email,
		*password,
		*nif,
		int(vo.User),
	)

	file, err := request.File.Open()

	if err != nil {
		return nil, err
	}

	defer file.Close()

	url, err := uc.fileService.UploadFile(&upload.UploadFile{
		Bucket:   upload.ClientBucket,
		FileName: request.File.Filename,
		FileId:   auth.GetId(),
		FileBody: file,
	})

	if err != nil {
		return nil, err
	}

	user := entities.NewUser(
		url,
		request.UserName,
		request.FirstName,
		request.LastName,
		*phone,
		*address,
	)

	if err := uc.authRepository.CreateAccount(auth, user); err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := jwt.CreateAuthTokens(jwt.AuthTokenPayload{
		Email:  auth.Email.GetValue(),
		UserId: auth.ID,
		Role:   auth.GetRole(),
	})

	if err != nil {
		return nil, err
	}

	return &contracts.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
