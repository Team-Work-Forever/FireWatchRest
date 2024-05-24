package main

import (
	"github.com/Team-Work-Forever/FireWatchRest/config"
	"github.com/Team-Work-Forever/FireWatchRest/internal/adapters"
	"github.com/Team-Work-Forever/FireWatchRest/internal/application/controllers"
	ucy "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/autarchy"
	uca "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/auth"
	ucb "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/burn"
	ucp "github.com/Team-Work-Forever/FireWatchRest/internal/application/usecases/profile"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/shared"
)

//	@title						FireWatch API
//	@version					1.0
//	@description				This is the api for Fire Watch Mobile Application
//	@termsOfService				http://swagger.io/terms/
//	@contact.name				API Support
//	@contact.email				fiber@swagger.io
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@BasePath					/api/v1
//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//
// @ Schemas http
func main() {
	config.LoadEnv(".env")
	env := config.GetCofig()

	// setup database
	db, err := adapters.NewDatabaseGorm()

	if err != nil {
		panic(err)
	}

	// Setup Fiber
	app := adapters.NewHttpServer(1)

	// Setup Swagger
	adapters.UseSwagger(app.Instance, env.FIRE_WATCH_API_PORT)

	// repositories
	authRepository := repositories.NewAuthRepository(db)
	tokenRepository := repositories.NewTokenRepository(db)
	profileRepository := repositories.NewProfileRepository(db)
	burnRepository := repositories.NewBurnRepository(db)
	autarchyRepository := repositories.NewAutarchyRepository(db)

	// use cases
	loginUseCase := uca.NewLoginUseCase(authRepository)
	signUpUseCase := uca.NewSignUpUseCase(authRepository)
	forgotPasswordUseCase := uca.NewForgotPasswordUseCase(authRepository, tokenRepository)
	resetPasswordUseCase := uca.NewResetPasswordUseCase(authRepository, tokenRepository)
	refreshTokensUseCase := uca.NewRefreshTokesUseCase(authRepository)

	whoamiUseCase := ucp.NewWhoamiUseCase(authRepository, profileRepository)
	updateProfileUseCase := ucp.NewUpdateProfileUIseCase(authRepository, profileRepository)

	createBurnUseCase := ucb.NewCreateBurnUseCase(burnRepository)
	getBurnbyIdUseCase := ucb.NewGetBurnByIdUseCase(burnRepository)
	getAllBurnsUseCase := ucb.NewGetAllBurnsUseCase(burnRepository)
	updateBurnUseCase := ucb.NewUpdateBurnUseCase(burnRepository)
	deleteBurnUseCase := ucb.NewDeleteBurnUseCase(burnRepository)

	createAutarchyUseCase := ucy.NewCreateAutarchyUseCase(autarchyRepository, authRepository)
	getAutarchyById := ucy.NewGetAutarchyByIdUseCase(autarchyRepository)
	getAllAutarchiesUseCase := ucy.NewGetAllAutarchies(autarchyRepository)

	// controllers
	authController := controllers.NewAuthController(loginUseCase, signUpUseCase, forgotPasswordUseCase, resetPasswordUseCase, refreshTokensUseCase)
	profileController := controllers.NewProfileController(whoamiUseCase, updateProfileUseCase)
	burnController := controllers.NewBurnController(createBurnUseCase, getBurnbyIdUseCase, getAllBurnsUseCase, updateBurnUseCase, deleteBurnUseCase)
	autarchyController := controllers.NewAutarchyController(createAutarchyUseCase, getAutarchyById, getAllAutarchiesUseCase)

	// Serve application
	app.AddControllers([]shared.Controller{
		authController,
		profileController,
		burnController,
		autarchyController,
	})

	app.Serve()
}
