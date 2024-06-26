package controllers

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/application/middlewares"
	usecases "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/auth"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/services"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	loginUseCase          *usecases.LoginUseCase
	signUpUseCase         *usecases.SignUpUseCase
	forgotPasswordUseCase *usecases.ForgotPasswordUseCase
	resetPasswordUseCase  *usecases.ResetPasswordUseCase
	refreshTokensUseCase  *usecases.RefreshTokensUseCase
}

func NewAuthController(
	loginUseCase *usecases.LoginUseCase,
	signUpUseCase *usecases.SignUpUseCase,
	forgotPasswordUseCase *usecases.ForgotPasswordUseCase,
	resetPasswordUseCase *usecases.ResetPasswordUseCase,
	refreshTokensUseCase *usecases.RefreshTokensUseCase,
) *AuthController {
	return &AuthController{
		loginUseCase:          loginUseCase,
		signUpUseCase:         signUpUseCase,
		forgotPasswordUseCase: forgotPasswordUseCase,
		resetPasswordUseCase:  resetPasswordUseCase,
		refreshTokensUseCase:  refreshTokensUseCase,
	}
}

func (c *AuthController) Route(router fiber.Router) {
	authRoutes := router.Group("auth")

	authRoutes.Post("login", middlewares.ShouldAcceptJson, c.LoginRoute)
	authRoutes.Post("signUp", middlewares.ShouldAcceptMultiPart, c.SignUpRoute)
	authRoutes.Get("forgot_password", c.ForgotPasswordRoute)
	authRoutes.Post("reset_password", middlewares.ShouldAcceptJson, c.ResetPasswordRoute)
	authRoutes.Get("refresh_tokens", c.RefreshTokensRoute)
}

// // ShowAccount godoc
//
//	@Summary	Authenticate with account
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//
//	@Param		accept-language	header		string					false	"some description"
//
//	@Param		data			body		contracts.LoginRequest	true	"Login Payload"
//	@Success	200				{object}	contracts.AuthResponse
//	@Router		/auth/login [post]
func (c *AuthController) LoginRoute(ctx *fiber.Ctx) error {
	var loginRequest contracts.LoginRequest

	if err := ctx.BodyParser(&loginRequest); err != nil {
		return err
	}

	tokens, err := c.loginUseCase.Handle(loginRequest)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(tokens)
}

// // ShowAccount godoc
//
//	@Summary	Create an Account
//	@Tags		Auth
//	@Accept		multipart/form-data
//	@Produce	json
//
//	@Param		accept-language	header		string					false	"some description"
//
//	@Param		data			formData	contracts.SignUpRequest	true	"Form data"
//	@Param		avatar			formData	file					true	"User avatar"
//	@Success	201				{object}	contracts.AuthResponse
//	@Router		/auth/signUp [post]
func (c *AuthController) SignUpRoute(ctx *fiber.Ctx) error {
	var signUpRequest contracts.SignUpRequest

	if err := ctx.BodyParser(&signUpRequest); err != nil {
		return err
	}

	fileHeader, err := services.GetFile(ctx, "avatar", false)

	if err != nil {
		return err
	}

	signUpRequest.File = fileHeader
	tokens, err := c.signUpUseCase.Handle(signUpRequest)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(tokens)
}

// // ShowAccount godoc
//
//	@Summary	Request a Password Reset
//	@Tags		Auth
//	@Produce	json
//
//	@Param		accept-language	header	string	false	"some description"
//
//	@Param		email			query	string	true	"Email address associated with the account"
//	@Success	200
//	@Router		/auth/forgot_password [get]
func (c *AuthController) ForgotPasswordRoute(ctx *fiber.Ctx) error {
	email := ctx.Query("email")

	err := c.forgotPasswordUseCase.Handle(contracts.ForgotPasswordRequest{
		Email: email,
	})

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&contracts.DefaultResponse{
		Message: "an email with password recovery instructions has been sent to your email address. If you have not received the email, please try again",
	})
}

// // ShowAccount godoc
//
//	@Summary	Reset Password
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//
//	@Param		accept-language	header		string							false	"some description"
//
//	@Param		forgot_token	query		string							true	"A unique token sent to the user's email for password reset"
//	@Param		data			body		contracts.ResetPasswordRequest	true	"Reset Password Payload"
//	@Success	200				{string}	string							"Password reset successfully"
//	@Router		/auth/reset_password [post]
func (c *AuthController) ResetPasswordRoute(ctx *fiber.Ctx) error {
	var resetPasswordRequest contracts.ResetPasswordRequest
	forgot_token := ctx.Query("forgot_token")

	if err := ctx.BodyParser(&resetPasswordRequest); err != nil {
		return err
	}

	resetPasswordRequest.ForgotToken = forgot_token
	err := c.resetPasswordUseCase.Handle(resetPasswordRequest)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&contracts.DefaultResponse{
		Message: "Password was reseted",
	})
}

// // ShowAccount godoc
//
//	@Summary	Refresh Tokens
//	@Tags		Auth
//	@Accept		json
//
//	@Param		accept-language	header	string	false	"some description"
//	@Param		token			query	string	true	"JWT token to be refreshed"
//	@Produce	json
//	@Success	201	{object}	contracts.AuthResponse
//	@Router		/auth/refresh_tokens [get]
func (c *AuthController) RefreshTokensRoute(ctx *fiber.Ctx) error {
	token := ctx.Query("token")

	tokens, err := c.refreshTokensUseCase.Handle(contracts.RefreshTokensRequest{
		RefreshToken: token,
	})

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(tokens)
}
