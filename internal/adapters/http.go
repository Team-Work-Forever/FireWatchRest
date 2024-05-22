package adapters

import (
	"fmt"
	"log"

	"github.com/Team-Work-Forever/FireWatchRest/config"
	"github.com/Team-Work-Forever/FireWatchRest/internal/application/middlewares"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/locales"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type HttpServer struct {
	Instance *fiber.App
	Version  string
}

func NewHttpServer(version int) *HttpServer {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})

	app.Use(logger.New())
	app.Use(locales.New())

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
