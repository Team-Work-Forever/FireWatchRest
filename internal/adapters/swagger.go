package adapters

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	docs "github.com/Team-Work-Forever/FireWatchRest/docs/fireWatch"
)

func UseSwagger(app *fiber.App, port string) {
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", port)
	docs.SwaggerInfo.Schemes = []string{"http"}

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New(swagger.Config{}))
}
