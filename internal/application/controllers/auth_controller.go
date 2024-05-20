package controllers

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/application/middlewares"
	usecases "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/auth"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	loginUseCase  *usecases.LoginUseCase
	signUpUseCase *usecases.SignUpUseCase
}

func NewAuthController(
	loginUseCase *usecases.LoginUseCase,
	signUpUseCase *usecases.SignUpUseCase,
) *AuthController {
	return &AuthController{
		loginUseCase:  loginUseCase,
		signUpUseCase: signUpUseCase,
	}
}

func (c *AuthController) Route(router fiber.Router) {
	authRoutes := router.Group("auth")

	authRoutes.Post("login", middlewares.ShouldAcceptJson, c.LoginRoute)
	authRoutes.Post("signUp", middlewares.ShouldAcceptMultiPart, c.SignUpRoute)
}

// // ShowAccount godoc
//
//	@Summary	Authenticate with account
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		data	body		contracts.LoginRequest	true	"Login Payload"
//	@Success	200		{object}	contracts.AuthResponse
//	@Security	BearerAuth
//	@Router		/auth/login [post]
func (c *AuthController) LoginRoute(ctx *fiber.Ctx) error {
	loginRequest := new(contracts.LoginRequest)

	if err := ctx.BodyParser(loginRequest); err != nil {
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
//	@Tags		auth
//	@Accept		multipart/form-data
//	@Produce	json
//	@Param		data formData contracts.SignUpRequest true "Form data"
//	@Param		avatar formData file true "User avatar"
//	@Success	200	{object}	contracts.AuthResponse
//	@Security	BearerAuth
//	@Router		/auth/signUp [post]
func (c *AuthController) SignUpRoute(ctx *fiber.Ctx) error {
	signUpRequest := new(contracts.SignUpRequest)

	if err := ctx.BodyParser(signUpRequest); err != nil {
		return err
	}

	tokens, err := c.signUpUseCase.Handle(signUpRequest)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(tokens)
}
