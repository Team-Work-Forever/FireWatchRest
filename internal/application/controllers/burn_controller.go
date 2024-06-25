package controllers

import (
	"strconv"

	"github.com/Team-Work-Forever/FireWatchRest/internal/application/middlewares"
	usecases "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/burn"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pagination"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/services"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/shared"
	"github.com/gofiber/fiber/v2"
)

type BurnController struct {
	createBurnUc  *usecases.CreateBurnUseCase
	getBurnByIdUc *usecases.GetBurnByIdUseCase
	getAllBurnsUc *usecases.GetAllBurnsUseCase
	updateBurnUc  *usecases.UpdateBurnUseCase
	deleteBurnUc  *usecases.DeleteBurnUseCase
	terminateUc   *usecases.TerminateBurnUseCase
	startUc       *usecases.StartBurnUseCase
}

func NewBurnController(
	createBurnUc *usecases.CreateBurnUseCase,
	getBurnByIdUc *usecases.GetBurnByIdUseCase,
	getAllBurnsUc *usecases.GetAllBurnsUseCase,
	updateBurnUc *usecases.UpdateBurnUseCase,
	deleteBurnUc *usecases.DeleteBurnUseCase,
	terminateUc *usecases.TerminateBurnUseCase,
	startUc *usecases.StartBurnUseCase,
) *BurnController {
	return &BurnController{
		createBurnUc:  createBurnUc,
		getBurnByIdUc: getBurnByIdUc,
		getAllBurnsUc: getAllBurnsUc,
		updateBurnUc:  updateBurnUc,
		deleteBurnUc:  deleteBurnUc,
		terminateUc:   terminateUc,
		startUc:       startUc,
	}
}

func (c *BurnController) Route(router fiber.Router) {
	burn := router.Group("burns", middlewares.AuthorizationMiddleware)

	burn.Get("types", c.GetBurnTypes)
	burn.Get("reasons", c.GetBurnReasons)
	burn.Get("states", c.GetBurnStates)
	burn.Get("availability/:lat/:lon", c.GetAvailabilty)

	burn.Post("", middlewares.ShouldAcceptMultiPart, c.CreateBurn)

	burn.Get("", c.GetAllBurns)
	burn.Get(":id", c.GetBurnById)

	burn.Put(":id", middlewares.ShouldAcceptMultiPart, c.UpdateBurn)
	burn.Put(":id/terminate", c.TerminateBurn)
	burn.Put(":id/start", c.StartBurn)
	burn.Delete(":id", c.DeleteBurn)
}

// // ShowAccount godoc
//
//	@Summary	Create an Burn Request
//	@Tags		Burn
//	@Accept		multipart/form-data
//	@Produce	json
//
//	@Param		accept-language	header		string						false	"some description"
//
//	@Param		data			formData	contracts.CreateBurnRequest	true	"Form data"
//
//	@Success	201				{object}	contracts.BurnActionResponse
//	@security	Bearer
//
//	@Router		/burns [post]
func (c *BurnController) CreateBurn(ctx *fiber.Ctx) error {
	var createBurnRequest contracts.CreateBurnRequest

	userId := shared.GetUserId(ctx)

	createBurnRequest.UserId = userId
	if err := ctx.BodyParser(&createBurnRequest); err != nil {
		return err
	}

	fileHeader, err := services.GetFile(ctx, "avatar", true)

	if err != nil {
		return err
	}

	createBurnRequest.MapPicture = fileHeader
	result, err := c.createBurnUc.Handler(createBurnRequest)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(result)
}

// // ShowAccount godoc
//
//	@Summary	Delete an Record if it is Scheduale
//	@Tags		Burn
//	@Produce	json
//
//	@Param		accept-language	header		string	false	"some description"
//
//	@Param		id				path		string	true	"Delete the burn by id"
//
//	@Success	202				{object}	contracts.BurnActionResponse
//
//	@security	Bearer
//
//	@Router		/burns/{id} [delete]
func (c *BurnController) DeleteBurn(ctx *fiber.Ctx) error {
	userId := shared.GetUserId(ctx)
	burnId := ctx.Params("id", "")

	result, err := c.deleteBurnUc.Handle(contracts.DeleteBurnRequest{
		UserId: userId,
		BurnId: burnId,
	})

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusAccepted).JSON(result)
}

// // ShowAccount godoc
//
//	@Summary	Create an Burn Request
//	@Tags		Burn
//	@Accept		multipart/form-data
//	@Produce	json
//
//	@Param		accept-language	header		string						false	"some description"
//
//	@Param		id				path		string						true	"Fetch the burn by id"
//	@Param		data			formData	contracts.CreateBurnRequest	true	"Form data"
//
//	@Success	202				{object}	contracts.BurnActionResponse
//
//	@security	Bearer
//
//	@Router		/burns/{id} [put]
func (c *BurnController) UpdateBurn(ctx *fiber.Ctx) error {
	var updateBurnRequest contracts.UpdateBurnRequest
	userId := shared.GetUserId(ctx)
	burnId := ctx.Params("id", "")

	updateBurnRequest.UserId = userId
	updateBurnRequest.BurnId = burnId
	if err := ctx.BodyParser(&updateBurnRequest); err != nil {
		return err
	}

	result, err := c.updateBurnUc.Handle(updateBurnRequest)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusAccepted).JSON(result)
}

// // ShowAccount godoc
//
//	@Summary	Terminate Burn
//	@Tags		Burn
//	@Accept		application/json
//	@Produce	json
//
//	@Param		accept-language	header		string						false	"some description"
//
//	@Param		id				path		string						true	"Fetch the burn by id"
//	@Param		data			formData	contracts.CreateBurnRequest	true	"Form data"
//
//	@Success	200				{object}	contracts.BurnActionResponse
//
//	@security	Bearer
//
//	@Router		/burns/{id} [put]
func (c *BurnController) TerminateBurn(ctx *fiber.Ctx) error {
	burnId := ctx.Params("id", "")
	userId := shared.GetUserId(ctx)

	result, err := c.terminateUc.Handle(contracts.TerminateBurnRequest{
		BurnId: burnId,
		UserId: userId,
	})

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusAccepted).JSON(result)
}

// // ShowAccount godoc
//
//	@Summary	Start Burn
//	@Tags		Burn
//	@Accept		application/json
//	@Produce	json
//
//	@Param		accept-language	header		string						false	"some description"
//
//	@Param		id				path		string						true	"Fetch the burn by id"
//	@Param		data			formData	contracts.CreateBurnRequest	true	"Form data"
//
//	@Success	200				{object}	contracts.BurnActionResponse
//
//	@security	Bearer
//
//	@Router		/burns/{id} [put]
func (c *BurnController) StartBurn(ctx *fiber.Ctx) error {
	burnId := ctx.Params("id", "")
	userId := shared.GetUserId(ctx)

	result, err := c.startUc.Handle(contracts.StartBurnRequest{
		BurnId: burnId,
		UserId: userId,
	})

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusAccepted).JSON(result)
}

// // ShowAccount godoc
//
//	@Summary	Fetch Burn By Id
//	@Tags		Burn
//	@Produce	json
//
//	@Param		accept-language	header		string	false	"some description"
//
//	@Param		id				path		string	true	"Fetch the burn by id"
//
//	@Success	200				{object}	geojson.GeoJsonFeature
//
//	@security	Bearer
//
//	@Router		/burns/{id} [get]
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

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// // ShowAccount godoc
//
//	@Summary	Fetch burns
//	@Tags		Burn
//	@Produce	json
//
//	@Param		accept-language	header		string	false	"some description"
//
//	@Param		search			query		string	false	"search burn title"
//	@Param		sort			query		string	false	"sort asc | desc"
//	@Param		state			query		string	false	"search by burn state"
//	@Param		start_date		query		string	false	"search by an inital date"
//	@Param		end_date		query		string	false	"search by an end date"
//	@Param		page			query		int		false	"view page"						default(1)
//	@Param		page_size		query		int		false	"number of returned elements"	default(10)
//
//	@Success	200				{object}	geojson.GeoJsonFeature
//
//	@security	Bearer
//
//	@Router		/burns [get]
func (c *BurnController) GetAllBurns(ctx *fiber.Ctx) error {
	userId := shared.GetUserId(ctx)
	search := ctx.Query("search", "")
	state := ctx.Query("state", "")
	sort := ctx.Query("sort", "asc")
	startDate := ctx.Query("start_date", "")
	endDate := ctx.Query("end_date", "")
	pageString := ctx.Query("page", "1")
	pageSizeString := ctx.Query("page_size", "10")

	page, err := pagination.New(pageString, pageSizeString)

	if err != nil {
		return err
	}

	result, err := c.getAllBurnsUc.Handle(contracts.GetAllBurnsRequest{
		AuthId:     userId,
		Search:     search,
		State:      state,
		Sort:       sort,
		StartDate:  startDate,
		EndDate:    endDate,
		Pagination: page,
	})

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// // ShowAccount godoc
//
//	@Summary	Burn Available Types
//	@Tags		Burn
//	@Produce	json
//
//	@Param		accept-language	header	string	false	"some description"
//
//	@Success	200				{array}	string
//
//	@security	Bearer
//
//	@Router		/burns/types [get]
func (c *BurnController) GetBurnTypes(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(vo.GetAllBurnTypes())
}

// // ShowAccount godoc
//
//	@Summary	Burn Available Reasons
//	@Tags		Burn
//	@Produce	json
//
//	@Param		accept-language	header	string	false	"some description"
//
//	@Success	200				{array}	string
//
//	@security	Bearer
//
//	@Router		/burns/reasons [get]
func (c *BurnController) GetBurnReasons(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(vo.GetAllBurnReasons())
}

// // ShowAccount godoc
//
//	@Summary	Burn Available States
//	@Tags		Burn
//	@Produce	json
//
//	@Param		accept-language	header	string	false	"some description"
//
//	@Success	200				{array}	string
//
//	@security	Bearer
//
//	@Router		/burns/states [get]
func (c *BurnController) GetBurnStates(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(vo.GetAllBurnStates())
}

// // ShowAccount godoc
//
//	@Summary	Check the availability to make burn
//	@Tags		Burn
//	@Produce	json
//
//	@Param		accept-language	header	string	false	"some description"
//
//	@Success	200				{array}	string
//
//	@security	Bearer
//
// @Router /burns/availability [get]
func (c *BurnController) GetAvailabilty(ctx *fiber.Ctx) error {
	lat := ctx.Params("lat", "0")
	lon := ctx.Params("lon", "0")
	hasAidTeam := ctx.Params("hasIdTeam", "false")
	ignore := ctx.Query("ignore", "false")

	ignoreValue, err := strconv.ParseBool(ignore)

	if err != nil {
		return err
	}

	latValue, err := strconv.ParseFloat(lat, 64)

	if err != nil {
		return err
	}

	lonValue, err := strconv.ParseFloat(lon, 64)

	if err != nil {
		return err
	}

	hasAidTeamValue, err := strconv.ParseBool(hasAidTeam)

	if err != nil {
		return err
	}

	result := services.CheckICFNIndex(latValue, lonValue, hasAidTeamValue, ignoreValue)

	return ctx.Status(fiber.StatusAccepted).JSON(contracts.AvailabilityResponse{
		Result: result,
	})
}
