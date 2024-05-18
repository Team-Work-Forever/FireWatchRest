package main

import (
	"fmt"
	"log"

	"github.com/Team-Work-Forever/FireWatchRest/config"
	"github.com/Team-Work-Forever/FireWatchRest/internal/application/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	docs "github.com/Team-Work-Forever/FireWatchRest/docs/fireWatch"
)

func ConfigureSwagger(app *fiber.App, port string) {
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", port)
	docs.SwaggerInfo.Schemes = []string{"http"}

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New(swagger.Config{}))
}

// @title			FireWatch API
// @version		1.0
// @description	This is the api for Fire Watch Mobile Application
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	fiber@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath		/api/v1
// @ Schemas http
func main() {
	config.LoadEnv("../.env")
	env := config.GetCofig()

	app := fiber.New()
	v1 := app.Group("api/v1")

	authController := controllers.NewAuthController()
	authController.Route(v1)

	app.Get("/", controllers.LoginRoute)
	ConfigureSwagger(app, env.FIRE_WATCH_PORT)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", env.FIRE_WATCH_PORT)))
}
