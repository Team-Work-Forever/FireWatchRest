package controllers

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/application/middlewares"
	usescases "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/autarchy"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	"github.com/gofiber/fiber/v2"
)

type AutarchyController struct {
	autarchyCreateUc  *usescases.CreateAutarchyUseCase
	autarchyGetByIdUc *usescases.GetAutarchyByIdUseCase
}

func NewAutarchyController(
	autarchyCreateUc *usescases.CreateAutarchyUseCase,
	autarchyGetByIdUc *usescases.GetAutarchyByIdUseCase,
) *AutarchyController {
	return &AutarchyController{
		autarchyCreateUc:  autarchyCreateUc,
		autarchyGetByIdUc: autarchyGetByIdUc,
	}
}

func (c *AutarchyController) Route(router fiber.Router) {
	autarchies := router.Group("autarchies", middlewares.AuthorizationMiddleware)

	autarchies.Post("", middlewares.ShouldAcceptMultiPart, c.CreateAutarchy)
	autarchies.Get(":id", c.GetAutarchyById)
}

// // ShowAccount godoc
//
//	@Summary	Create an Account for an Autarchy
//	@Tags		Autarchy
//	@Accept		multipart/form-data
//	@Produce	json
//
//	@Param		accept-language	header		string					false	"some description"
//
//	@Param		data			formData	contracts.CreateAutarchyRequest	true	"Form data"
//	@Param		avatar			formData	file					true	"User avatar"
//	@Success	201				{object}	contracts.AutarchyActionResponse
//	@security	Bearer
//
// @Router		/autarchies [post]
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

// // ShowAccount godoc
//
//	@Summary	Get Autarchy By Id
//	@Tags		Autarchy
//	@Produce	json
//
//	@Param		accept-language	header		string					false	"some description"
//	@Param		id			path		string	true	"Fetch the autarchy by id"
//
//	@Success	200				{object}	contracts.AutarchyActionResponse
//	@security	Bearer
//
// @Router		/autarchies/{id} [get]
func (c *AutarchyController) GetAutarchyById(ctx *fiber.Ctx) error {
	autarchyId := ctx.Params("id", "")

	result, err := c.autarchyGetByIdUc.Handle(contracts.GetAutarchyRequest{
		AutarchyId: autarchyId,
	})

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
