package controllers

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/application/middlewares"
	usecases "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/profile"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/shared"
	"github.com/gofiber/fiber/v2"
)

type ProfileController struct {
	whoamiUseCase *usecases.WhoamiUseCase
}

func NewProfileController(whoamiUseCase *usecases.WhoamiUseCase) *ProfileController {
	return &ProfileController{
		whoamiUseCase: whoamiUseCase,
	}
}

func (c *ProfileController) Route(router fiber.Router) {
	auth := router.Group("", middlewares.AuthorizationMiddleware)
	auth.Get("whoami", c.WhoamiRoute)
}

// // ShowAccount godoc
//
//	@Summary	Fetch Profile Information
//	@Tags		Profile
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	contracts.ProfileResponse
//
//	@security	Bearer
//
//	@Router		/whoami [get]
func (c *ProfileController) WhoamiRoute(ctx *fiber.Ctx) error {
	userId := shared.GetUserId(ctx)

	result, err := c.whoamiUseCase.Handle(contracts.WhoamiRequest{
		UserId: userId,
	})

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
