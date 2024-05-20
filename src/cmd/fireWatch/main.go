package main

import (
	"github.com/Team-Work-Forever/FireWatchRest/config"
	"github.com/Team-Work-Forever/FireWatchRest/internal/adapters"
	"github.com/Team-Work-Forever/FireWatchRest/internal/application/controllers"
	usecases "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/auth"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
)

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

	// Setup Fiber
	app := adapters.NewHttpServer()
	version := app.GetVersion("v1")

	// Setup Swagger
	adapters.UseSwagger(app.Instance, env.FIRE_WATCH_PORT)

	// repositories
	authRepository := repositories.NewAuthRepository()

	// use cases
	loginUseCase := usecases.NewLoginUseCase(authRepository)

	// controllers
	authController := controllers.NewAuthController(loginUseCase)
	authController.Route(version)

	// Serve application
	app.Serve()
}
