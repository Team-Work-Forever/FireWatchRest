package usecases

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/jwt"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type ResetPasswordUseCase struct {
	authRepository  *repositories.AuthRepository
	tokenRepository *repositories.TokenRepostory
}

func NewResetPasswordUseCase(
	authRepository *repositories.AuthRepository,
	tokenRepository *repositories.TokenRepostory,
) *ResetPasswordUseCase {
	return &ResetPasswordUseCase{
		authRepository:  authRepository,
		tokenRepository: tokenRepository,
	}
}

func (r *ResetPasswordUseCase) Handle(request contracts.ResetPasswordRequest) error {
	token, err := r.tokenRepository.GetByToken(request.ForgotToken, entities.ForgotToken)

	if err != nil {
		return exec.TOKEN_NOT_FOUND
	}

	if !jwt.ValidateToken(request.ForgotToken) {
		r.tokenRepository.Delete(token)

		return exec.TOKEN_ISNT_VALD
	}

	claims, err := jwt.GetClaims(token.Token, &jwt.AuthClaims{})

	if err != nil {
		r.tokenRepository.Delete(token)

		return exec.FAILED_FETCHING_TOKEN
	}

	authClaims, ok := claims.(*jwt.AuthClaims)

	if !ok {
		r.tokenRepository.Delete(token)

		return exec.FAILED_FETCHING_TOKEN
	}

	email, err := vo.NewEmail(authClaims.Email)

	if err != nil {
		return err
	}

	foundAuth, err := r.authRepository.GetAuthByEmail(email)

	if err != nil {
		return exec.USER_NOT_FOUND
	}

	password, err := vo.NewPassword(request.Password)

	if err != nil {
		return err
	}

	confirmPassword, err := vo.NewPassword(request.ConfirmPassword)

	if err != nil {
		return err
	}

	if password.GetValue() != confirmPassword.GetValue() {
		return exec.PASSWORDS_DONT_MATCH
	}

	if err := foundAuth.ChangePassword(password); err != nil {
		return err
	}

	if err := r.authRepository.Update(foundAuth); err != nil {
		return err
	}

	r.tokenRepository.Delete(token)
	return nil
}
