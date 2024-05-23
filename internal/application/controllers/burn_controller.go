package controllers

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/application/middlewares"
	usecases "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/burn"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/shared"
	"github.com/gofiber/fiber/v2"
)

type BurnController struct {
	createBurnUc  *usecases.CreateBurnUseCase
	getBurnByIdUc *usecases.GetBurnByIdUseCase
}

func NewBurnController(
	createBurnUc *usecases.CreateBurnUseCase,
	getBurnByIdUc *usecases.GetBurnByIdUseCase,
) *BurnController {
	return &BurnController{
		createBurnUc:  createBurnUc,
		getBurnByIdUc: getBurnByIdUc,
	}
}

func (c *BurnController) Route(router fiber.Router) {
	burn := router.Group("burn", middlewares.AuthorizationMiddleware)

	burn.Post("", middlewares.ShouldAcceptMultiPart, c.CreateBurn)
	burn.Get(":id", c.GetBurnById)

	burn.Get("types", c.GetBurnTypes)
	burn.Get("reasons", c.GetBurnReasons)
	burn.Get("states", c.GetBurnStates)
}

// // ShowAccount godoc
//
//	@Summary	Create an Burn Request
//	@Tags		Burn
//	@Accept		multipart/form-data
//	@Produce	json
//
// @Param   accept-language  header     string     false  "some description"
//
//	@Param		data	formData	contracts.CreateBurnRequest	true	"Form data"
//
//	@Success	201	{object}	contracts.ProfileResponse
//
//	@security	Bearer
//
//	@Router		/burn [post]
func (c *BurnController) CreateBurn(ctx *fiber.Ctx) error {
	var createBurnRequest contracts.CreateBurnRequest

	userId := shared.GetUserId(ctx)

	createBurnRequest.UserId = userId
	if err := ctx.BodyParser(&createBurnRequest); err != nil {
		return err
	}

	result, err := c.createBurnUc.Handler(createBurnRequest)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(result)
}

// // ShowAccount godoc
//
//	@Summary	Fetch Burn By Id
//	@Tags		Burn
//	@Produce	json
//
// @Param   accept-language  header     string     false  "some description"
//
//	@Param		burnId	query	string	true	"Fetch the burn by id"
//
//	@Success	200	{object}	geojson.GeoJsonFeature
//
//	@security	Bearer
//
//	@Router		/burn [get]
func (c *BurnController) GetBurnById(ctx *fiber.Ctx) error {
	burnId := ctx.Params("id", "")
	userId := shared.GetUserId(ctx)

	result, err := c.getBurnByIdUc.Handle(contracts.GetBurnRequest{
		AuthId: userId,
		BurnId: burnId,
	})

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(result)
}

// // ShowAccount godoc
//
//	@Summary	Burn Available Types
//	@Tags		Burn
//	@Produce	json
//
// @Param   accept-language  header     string     false  "some description"
//
//	@Success	200	{array} string
//
//	@security	Bearer
//
//	@Router		/burn/types [get]
func (c *BurnController) GetBurnTypes(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(vo.GetAllBurnTypes())
}

// // ShowAccount godoc
//
//	@Summary	Burn Available Reasons
//	@Tags		Burn
//	@Produce	json
//
// @Param   accept-language  header     string     false  "some description"
//
//	@Success	200	{array} string
//
//	@security	Bearer
//
//	@Router		/burn/reasons [get]
func (c *BurnController) GetBurnReasons(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(vo.GetAllBurnReasons())
}

// // ShowAccount godoc
//
//	@Summary	Burn Available States
//	@Tags		Burn
//	@Produce	json
//
// @Param   accept-language  header     string     false  "some description"
//
//	@Success	200	{array} string
//
//	@security	Bearer
//
//	@Router		/burn/states [get]
func (c *BurnController) GetBurnStates(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(vo.GetAllBurnStates())
}
