package usecases

import (
	"fmt"
	"log"
	"time"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/jwt"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/smtp"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
)

type ForgotPasswordUseCase struct {
	authRepository  *repositories.AuthRepository
	tokenRepository *repositories.TokenRepostory
}

func NewForgotPasswordUseCase(
	authRepository *repositories.AuthRepository,
	tokenRepository *repositories.TokenRepostory,
) *ForgotPasswordUseCase {
	return &ForgotPasswordUseCase{
		authRepository:  authRepository,
		tokenRepository: tokenRepository,
	}
}

func (f *ForgotPasswordUseCase) Handle(request contracts.ForgotPasswordRequest) error {
	email, err := vo.NewEmail(request.Email)

	if err != nil {
		return err
	}

	// find user by email
	foundAuth, err := f.authRepository.GetAuthByEmail(email)

	if err != nil {
		return err
	}

	// generate an token to safe guard the request
	expire_at := time.Now().Add(time.Duration(5) * 24 * time.Minute)
	forgotToken, err := jwt.CreateJwtToken(jwt.TokenPayload{
		UserId:   foundAuth.ID,
		Email:    foundAuth.Email.GetValue(),
		Role:     "admin",
		Duration: expire_at,
	})

	if err != nil {
		return err
	}

	// store token
	if err := f.tokenRepository.Create(entities.NewToken(
		forgotToken,
		"forgot_token",
		expire_at,
	)); err != nil {
		return err
	}

	// send the token by email
	mail := smtp.New(
		email.GetValue(),
		"Forgot Password Request",
		fmt.Sprintf("token: %s", forgotToken),
	)

	sendMail := func() {
		if err := mail.Send(); err != nil {
			log.Print(err)
		}
	}

	go sendMail()

	return nil
}
