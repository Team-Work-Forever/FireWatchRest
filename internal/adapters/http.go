package adapters

import (
	"fmt"
	"log"

	"github.com/Team-Work-Forever/FireWatchRest/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type HttpServer struct {
	Instance *fiber.App
}

func NewHttpServer() *HttpServer {
	app := fiber.New()

	app.Use(logger.New())

	return &HttpServer{
		Instance: app,
	}
}

func (hs *HttpServer) GetVersion(version string) fiber.Router {
	return hs.Instance.Group(fmt.Sprintf("api/%s", version))
}

func (hs *HttpServer) Serve() {
	env := config.GetCofig()

	log.Fatal(hs.Instance.Listen(fmt.Sprintf(":%s", env.FIRE_WATCH_API_PORT)))
}
