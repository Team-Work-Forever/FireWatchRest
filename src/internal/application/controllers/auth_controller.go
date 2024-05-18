package controllers

import (
	usecases "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/auth"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (c *AuthController) Route(router fiber.Router) {
	authRoutes := router.Group("auth")

	authRoutes.Post("login", LoginRoute)
	authRoutes.Post("signUp", SignUpRoute)
}

// // ShowAccount godoc
//
//	@Summary	Authenticate with account
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		data	body		contracts.LoginRequest	true	"Login Payload"
//	@Success	200		{object}	contracts.DefaultResponse
//	@Security	BearerAuth
//	@Router		/api/v1/auth/login [post]
func LoginRoute(ctx *fiber.Ctx) error {
	useCase := usecases.NewLoginUseCase()
	result := useCase.Handle()

	ctx.Status(fiber.StatusOK).JSON(contracts.DefaultResponse{
		Code:  fiber.StatusOK,
		Title: result,
	})

	return nil
}

// // ShowAccount godoc
//
//	@Summary	Create an Account
//	@Tags		auth
//	@Accept		multipart/form-data
//	@Produce	json
//	@Param		data formData contracts.SignUpRequest true "Form data"
//	@Param		file formData file true "User avatar"
//	@Success	200	{object}	contracts.DefaultResponse
//	@Security	BearerAuth
//	@Router		/api/v1/auth/signUp [post]
func SignUpRoute(ctx *fiber.Ctx) error {
	useCase := usecases.NewLoginUseCase()
	result := useCase.Handle()

	ctx.Status(fiber.StatusOK).JSON(contracts.DefaultResponse{
		Code:  fiber.StatusOK,
		Title: result,
	})

	return nil
}
