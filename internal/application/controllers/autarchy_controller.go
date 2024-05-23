package controllers

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/application/middlewares"
	usescases "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/autarchy"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	"github.com/gofiber/fiber/v2"
)

type AutarchyController struct {
	autarchyCreateUc *usescases.CreateAutarchyUseCase
}

func NewAutarchyController(autarchyCreateUc *usescases.CreateAutarchyUseCase) *AutarchyController {
	return &AutarchyController{
		autarchyCreateUc: autarchyCreateUc,
	}
}

func (c *AutarchyController) Route(router fiber.Router) {
	autarchies := router.Group("autarchies")
	autarchies.Post("", middlewares.ShouldAcceptMultiPart, c.CreateAutarchy)
}

// // ShowAccount godoc
//
//	@Summary	Create an Account
//	@Tags		Autarchy
//	@Accept		multipart/form-data
//	@Produce	json
//
//	@Param		accept-language	header		string					false	"some description"
//
//	@Param		data			formData	contracts.SignUpRequest	true	"Form data"
//	@Param		avatar			formData	file					true	"User avatar"
//	@Success	201				{object}	contracts.AuthResponse
//	@Router		/autarchy [post]
func (c *AutarchyController) CreateAutarchy(ctx *fiber.Ctx) error {
	var createRequest contracts.CreateAutarchyRequest

	if err := ctx.BodyParser(&createRequest); err != nil {
		return err
	}

	result, err := c.autarchyCreateUc.Handle(createRequest)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(result)
}
