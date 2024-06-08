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
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/key"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/upload"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/shared"
	"github.com/golang-migrate/migrate/v4"
	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) error {
	migration, err := migrate.New(
		"file://./docker/migrations",
		adapters.GetConnectionString(),
	)

	if err != nil {
		return err
	}

	return migration.Up()
}

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

	// Setup Aux
	db := adapters.GetDatabase()

	// if err := Migrations(db); err != nil {
	// 	panic(err)
	// }

	kvService := key.NewKeyValueService()
	defer kvService.Close()

	fileService := upload.NewBlobService()

	// Setup Fiber
	app := adapters.NewHttpServer(1)

	// Setup Swagger
	adapters.UseSwagger(app.Instance, env.FIRE_WATCH_API_PORT)

	// repositories
	authRepository := repositories.NewAuthRepository(db)
	tokenRepository := repositories.NewTokenRepository(kvService)
	profileRepository := repositories.NewProfileRepository(db)
	burnRepository := repositories.NewBurnRepository(db)
	autarchyRepository := repositories.NewAutarchyRepository(db)

	// use cases
	loginUseCase := uca.NewLoginUseCase(authRepository)
	signUpUseCase := uca.NewSignUpUseCase(authRepository, fileService)
	forgotPasswordUseCase := uca.NewForgotPasswordUseCase(authRepository, tokenRepository)
	resetPasswordUseCase := uca.NewResetPasswordUseCase(authRepository, tokenRepository)
	refreshTokensUseCase := uca.NewRefreshTokesUseCase(authRepository)

	whoamiUseCase := ucp.NewWhoamiUseCase(authRepository, profileRepository)
	updateProfileUseCase := ucp.NewUpdateProfileUIseCase(authRepository, profileRepository, fileService)

	createBurnUseCase := ucb.NewCreateBurnUseCase(burnRepository, autarchyRepository)
	getBurnbyIdUseCase := ucb.NewGetBurnByIdUseCase(burnRepository)
	getAllBurnsUseCase := ucb.NewGetAllBurnsUseCase(burnRepository, autarchyRepository)
	updateBurnUseCase := ucb.NewUpdateBurnUseCase(burnRepository)
	deleteBurnUseCase := ucb.NewDeleteBurnUseCase(burnRepository)

	createAutarchyUseCase := ucy.NewCreateAutarchyUseCase(autarchyRepository, authRepository, fileService)
	getAutarchyById := ucy.NewGetAutarchyByIdUseCase(autarchyRepository)
	getAllAutarchiesUseCase := ucy.NewGetAllAutarchies(autarchyRepository)
	updateAutarchyUseCase := ucy.NewUpdateAutarchyUseCase(autarchyRepository, authRepository, fileService)
	deleteAutarchyUseCase := ucy.NewDeleteAutarchyUseCase(autarchyRepository)

	// controllers
	authController := controllers.NewAuthController(loginUseCase, signUpUseCase, forgotPasswordUseCase, resetPasswordUseCase, refreshTokensUseCase)
	profileController := controllers.NewProfileController(whoamiUseCase, updateProfileUseCase)
	burnController := controllers.NewBurnController(createBurnUseCase, getBurnbyIdUseCase, getAllBurnsUseCase, updateBurnUseCase, deleteBurnUseCase)
	autarchyController := controllers.NewAutarchyController(createAutarchyUseCase, getAutarchyById, getAllAutarchiesUseCase, updateAutarchyUseCase, deleteAutarchyUseCase, getAllBurnsUseCase)

	// Serve application
	app.AddControllers([]shared.Controller{
		authController,
		profileController,
		burnController,
		autarchyController,
	})

	app.Serve()
}
