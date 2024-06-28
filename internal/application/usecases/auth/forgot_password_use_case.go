package usecases

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/smtp"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

const (
	MAX_TRIES = 5
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

func generateCode(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func (f *ForgotPasswordUseCase) Handle(request contracts.ForgotPasswordRequest) error {
	email, err := vo.NewEmail(request.Email)

	if err != nil {
		return err
	}

	// find user by email
	if !f.authRepository.ExistsUserWithEmail(email) {
		return exec.USER_NOT_FOUND
	}

	var code string

	for i := 0; i < MAX_TRIES; i++ {
		if i == MAX_TRIES-1 {
			return exec.TRY_AGAIN
		}

		code = strconv.Itoa(generateCode(1000, 9000))

		foundToken, err := f.tokenRepository.GetByToken(code, entities.ForgotToken)

		if foundToken == nil && err != nil {
			break
		}
	}

	// store token
	if err := f.tokenRepository.Create(entities.NewToken(
		code,
		email.Value,
		entities.ForgotToken,
	)); err != nil {
		return err
	}

	// send the token by email
	mail := smtp.New(
		email.GetValue(),
		"Forgot Password Request",
		fmt.Sprintf("token: %s", code),
	)

	sendMail := func() {
		if err := mail.Send(); err != nil {
			log.Print(err)
		}
	}

	go sendMail()

	return nil
}
