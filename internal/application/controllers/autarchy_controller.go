package controllers

import (
	"strconv"

	"github.com/Team-Work-Forever/FireWatchRest/internal/application/middlewares"
	usescases "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/autarchy"
	butnUsescases "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/burn"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pagination"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/shared"
	"github.com/gofiber/fiber/v2"
)

type AutarchyController struct {
	autarchyCreateUc  *usescases.CreateAutarchyUseCase
	autarchyGetByIdUc *usescases.GetAutarchyByIdUseCase
	autarchyGetAllUc  *usescases.GetAllAutarchies
	updateAutarchyUc  *usescases.UpdateAutarchyUseCase
	deleteAutarchyUc  *usescases.DeleteAutarchyUseCase
	burnGetAll        *butnUsescases.GetAllBurnsUseCase
}

func NewAutarchyController(
	autarchyCreateUc *usescases.CreateAutarchyUseCase,
	autarchyGetByIdUc *usescases.GetAutarchyByIdUseCase,
	autarchyGetAllUc *usescases.GetAllAutarchies,
	updateAutarchyUc *usescases.UpdateAutarchyUseCase,
	deleteAutarchyUc *usescases.DeleteAutarchyUseCase,
	burnGetalluc *butnUsescases.GetAllBurnsUseCase,
) *AutarchyController {
	return &AutarchyController{
		autarchyCreateUc:  autarchyCreateUc,
		autarchyGetByIdUc: autarchyGetByIdUc,
		autarchyGetAllUc:  autarchyGetAllUc,
		updateAutarchyUc:  updateAutarchyUc,
		deleteAutarchyUc:  deleteAutarchyUc,
		burnGetAll:        burnGetalluc,
	}
}

func (c *AutarchyController) Route(router fiber.Router) {
	autarchies := router.Group("autarchies", middlewares.AuthorizationMiddleware)

	autarchies.Post("", middlewares.AdminMiddleware, middlewares.ShouldAcceptMultiPart, c.CreateAutarchy)
	autarchies.Put(":id", middlewares.AdminMiddleware, middlewares.ShouldAcceptMultiPart, c.UpdateAutarchy)

	autarchies.Get("", c.GetAllAutarchies)
	autarchies.Get(":id", c.GetAutarchyById)
	autarchies.Get(":id/burns", c.GetAutarchyBurns)

	autarchies.Delete(":id", middlewares.AdminMiddleware, c.DeleteAutarchy)
}

// // ShowAccount godoc
//
//	@Summary	Create an Account for an Autarchy
//	@Tags		Autarchy
//	@Accept		multipart/form-data
//	@Produce	json
//
//	@Param		accept-language	header		string							false	"some description"
//
//	@Param		data			formData	contracts.CreateAutarchyRequest	true	"Form data"
//	@Param		avatar			formData	file							true	"User avatar"
//	@Success	201				{object}	contracts.AutarchyActionResponse
//	@security	Bearer
//
//	@Router		/autarchies [post]
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
//	@Summary	Update an Burn Request
//	@Tags		Autarchy
//	@Accept		multipart/form-data
//	@Produce	json
//
//	@Param		accept-language	header		string							false	"some description"
//
//	@Param		id				path		string							true	"Fetch the autarchy by id"
//	@Param		data			formData	contracts.UpdateAutarchyRequest	true	"Form data"
//
//	@Success	202				{object}	geojson.GeoJsonFeature
//
//	@security	Bearer
//
//	@Router		/autarchies/{id} [put]
func (c *AutarchyController) UpdateAutarchy(ctx *fiber.Ctx) error {
	var updateAutarchyRequest contracts.UpdateAutarchyRequest
	autarchyId := ctx.Params("id", "")

	updateAutarchyRequest.AutarchyId = autarchyId
	if err := ctx.BodyParser(&updateAutarchyRequest); err != nil {
		return err
	}

	result, err := c.updateAutarchyUc.Handle(updateAutarchyRequest)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusAccepted).JSON(result)
}

// // ShowAccount godoc
//
//	@Summary	Get Autarchy By Id
//	@Tags		Autarchy
//	@Produce	json
//
//	@Param		accept-language	header		string	false	"some description"
//	@Param		id				path		string	true	"Fetch the autarchy by id"
//
//	@Success	200				{object}	contracts.AutarchyActionResponse
//	@security	Bearer
//
//	@Router		/autarchies/{id} [get]
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

// // ShowAccount godoc
//
//	@Summary	Get Burns of Autarchy
//	@Tags		Autarchy
//	@Produce	json
//
//	@Param		accept-language	header		string	false	"some description"
//	@Param		id				path		string	true	"Fetch the autarchy by id"
//
//	@Param		search			query		string	false	"search burn title"
//	@Param		state			query		string	false	"search by burn state"
//	@Param		start_date		query		string	false	"search by an inital date"
//	@Param		end_date		query		string	false	"search by an end date"
//	@Param		page			query		int		false	"view page"						default(1)
//	@Param		page_size		query		int		false	"number of returned elements"	default(10)
//
//	@Success	200				{object}	geojson.GeoJsonFeature
//	@security	Bearer
//
//	@Router		/autarchies/{id}/burns [get]
func (c *AutarchyController) GetAutarchyBurns(ctx *fiber.Ctx) error {
	autarchyId := ctx.Params("id", "")
	userId := shared.GetUserId(ctx)
	search := ctx.Query("search", "")
	state := ctx.Query("state", "")
	startDate := ctx.Query("start_date", "")
	endDate := ctx.Query("end_date", "")
	pageString := ctx.Query("page", "1")
	pageSizeString := ctx.Query("page_size", "10")

	page, err := pagination.New(pageString, pageSizeString)

	if err != nil {
		return err
	}

	result, err := c.burnGetAll.Handle(contracts.GetAllBurnsRequest{
		AutarchyId: autarchyId,
		AuthId:     userId,
		Search:     search,
		State:      state,
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
//	@Summary	Fetch Autarchies
//	@Tags		Autarchy
//	@Produce	json
//
//	@Param		accept-language	header		string	false	"some description"
//
//	@Param		search			query		string	false	"search burn title"
//	@Param		page			query		int		false	"view page"						default(1)
//	@Param		page_size		query		int		false	"number of returned elements"	default(10)
//
//	@Success	200				{object}	geojson.GeoJsonFeature
//
//	@security	Bearer
//
//	@Router		/autarchies [get]
func (c *AutarchyController) GetAllAutarchies(ctx *fiber.Ctx) error {
	search := ctx.Query("search", "")
	pageString := ctx.Query("page", "1")
	pageSizeString := ctx.Query("page_size", "10")

	page, err := strconv.ParseUint(pageString, 10, 64)

	if err != nil {
		return err
	}

	pageSize, err := strconv.ParseUint(pageSizeString, 10, 64)

	if err != nil {
		return err
	}

	result, err := c.autarchyGetAllUc.Handle(contracts.GetAllAutarchiesRequest{
		Search:   search,
		Page:     page,
		PageSize: pageSize,
	})

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// // ShowAccount godoc
//
//	@Summary	Delete an Autarchy
//	@Tags		Autarchy
//	@Produce	json
//
//	@Param		accept-language	header		string	false	"some description"
//
//	@Param		id				path		string	true	"Delete the autarchy by id"
//
//	@Success	202				{object}	contracts.AutarchyActionResponse
//
//	@security	Bearer
//
//	@Router		/autarchies/{id} [delete]
func (c *AutarchyController) DeleteAutarchy(ctx *fiber.Ctx) error {
	userId := shared.GetUserId(ctx)
	autarchyId := ctx.Params("id", "")

	result, err := c.deleteAutarchyUc.Handle(contracts.DeleteAutarchyRequest{
		UserId:     userId,
		AutarchyId: autarchyId,
	})

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusAccepted).JSON(result)
}
