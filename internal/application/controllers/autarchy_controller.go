package controllers

import (
	"strconv"

	"github.com/Team-Work-Forever/FireWatchRest/internal/application/middlewares"
	usescases "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/autarchy"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	"github.com/gofiber/fiber/v2"
)

type AutarchyController struct {
	autarchyCreateUc  *usescases.CreateAutarchyUseCase
	autarchyGetByIdUc *usescases.GetAutarchyByIdUseCase
	autarchyGetAllUc  *usescases.GetAllAutarchies
}

func NewAutarchyController(
	autarchyCreateUc *usescases.CreateAutarchyUseCase,
	autarchyGetByIdUc *usescases.GetAutarchyByIdUseCase,
	autarchyGetAllUc *usescases.GetAllAutarchies,
) *AutarchyController {
	return &AutarchyController{
		autarchyCreateUc:  autarchyCreateUc,
		autarchyGetByIdUc: autarchyGetByIdUc,
		autarchyGetAllUc:  autarchyGetAllUc,
	}
}

func (c *AutarchyController) Route(router fiber.Router) {
	autarchies := router.Group("autarchies", middlewares.AuthorizationMiddleware)

	autarchies.Post("", middlewares.ShouldAcceptMultiPart, c.CreateAutarchy)
	autarchies.Get(":id", c.GetAutarchyById)
	autarchies.Get("", c.GetAllAutarchies)
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
