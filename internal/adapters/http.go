package adapters

import (
	"fmt"
	"log"

	"github.com/Team-Work-Forever/FireWatchRest/config"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/locales"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type HttpServer struct {
	Instance *fiber.App
	Version  string
}

func NewHttpServer(version int) *HttpServer {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	app.Use(logger.New())
	app.Use(locales.New())
	app.Use(cors.New())

	return &HttpServer{
		Instance: app,
		Version:  fmt.Sprintf("v%d", version),
	}
}

func (hs *HttpServer) GetVersion() fiber.Router {
	return hs.Instance.Group(fmt.Sprintf("api/%s", hs.Version))
}

func (hs *HttpServer) AddControllers(controllers []shared.Controller) {
	for i := 0; i < len(controllers); i++ {
		controllers[i].Route(hs.GetVersion())
	}
}

func (hs *HttpServer) Serve() {
	env := config.GetCofig()

	log.Fatal(hs.Instance.Listen(fmt.Sprintf(":%s", env.FIRE_WATCH_API_PORT)))
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case *exec.Error:
		return shared.WriteProblemDetails(ctx, *e)
	case *fiber.Error:
		return shared.WriteProblemDetails(
			ctx,
			exec.Error{
				Title:  "Bad Input",
				Status: e.Code,
				Detail: e.Message,
			},
		)
	default:
		return shared.WriteProblemDetails(
			ctx,
			exec.Error{
				Title:  "Internal Server Error",
				Status: fiber.StatusInternalServerError,
				Detail: err.Error(),
			},
		)
	}
}
