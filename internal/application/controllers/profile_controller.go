package controllers

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/application/middlewares"
	usecases "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/profile"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/services"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/shared"
	"github.com/gofiber/fiber/v2"
)

type ProfileController struct {
	whoamiUseCase        *usecases.WhoamiUseCase
	updateProfileUseCase *usecases.UpdateProfileUseCase
}

func NewProfileController(
	whoamiUseCase *usecases.WhoamiUseCase,
	updateProfileUseCase *usecases.UpdateProfileUseCase,
) *ProfileController {
	return &ProfileController{
		whoamiUseCase:        whoamiUseCase,
		updateProfileUseCase: updateProfileUseCase,
	}
}

func (c *ProfileController) Route(router fiber.Router) {
	auth := router.Group("", middlewares.AuthorizationMiddleware)
	auth.Get("whoami", c.WhoamiRoute)

	profile := auth.Group("profile")
	profile.Put("", middlewares.ShouldAcceptMultiPart, c.UpdateProfile)

	profile.Get("locale", c.Locale)
}

// // ShowAccount godoc
//
//	@Summary	Fetch Profile Information
//	@Tags		Profile
//	@Accept		multipart/form-data
//	@Produce	json
//
//	@Param		accept-language	header		string	false	"some description"
//
//	@Success	200				{object}	contracts.ProfileResponse
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

// // ShowAccount godoc
//
//	@Summary	Update Profile Information
//	@Tags		Profile
//	@Accept		multipart/form-data
//	@Produce	json
//
//	@Param		accept-language	header		string							false	"some description"
//
//	@Param		data			formData	contracts.UpdateProfileResponse	true	"Form data"
//	@Param		avatar			formData	file							true	"User avatar"
//
//	@Success	202				{object}	contracts.ProfileResponse
//
//	@security	Bearer
//
//	@Router		/profile [put]
func (c *ProfileController) UpdateProfile(ctx *fiber.Ctx) error {
	var updateRequest contracts.UpdateProfileResponse

	userId := shared.GetUserId(ctx)

	if err := ctx.BodyParser(&updateRequest); err != nil {
		return err
	}

	updateRequest.UserId = userId
	result, err := c.updateProfileUseCase.Handle(updateRequest)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusAccepted).JSON(result)
}

type Response struct {
	Lat float32
	Lon float32
}

func (c *ProfileController) Locale(ctx *fiber.Ctx) error {
	lat := ctx.Query("lat", "")
	lon := ctx.Query("lon", "")

	latResult, lonResult, err := services.GetCoordinates(lat, lon)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&Response{
		Lat: latResult,
		Lon: lonResult,
	})
}
