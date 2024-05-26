package middlewares

import (
	"github.com/Team-Work-Forever/FireWatchRest/config"
	"github.com/Team-Work-Forever/FireWatchRest/internal/adapters"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	jwtService "github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/jwt"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/shared"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizationMiddleware(ctx *fiber.Ctx) error {
	env := config.GetCofig()

	return jwtware.New(jwtware.Config{
		Claims: &jwtService.AuthClaims{},
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwt.SigningMethodHS256.Name,
			Key:    []byte(env.JWT_SECRET),
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			authRepository := repositories.NewAuthRepository(adapters.GetDatabase())

			if ok := authRepository.ExistsAuthById(shared.GetUserId(ctx)); !ok {
				return shared.WriteProblemDetails(ctx, exec.Error{
					Title:  "Authorization Failed",
					Detail: "you don't have access to this resource",
					Status: fiber.StatusUnauthorized,
				})
			}

			return ctx.Next()
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return shared.WriteProblemDetails(ctx, exec.Error{
				Title:  "Authorization Failed",
				Detail: err.Error(),
				Status: fiber.StatusUnauthorized,
			})
		},
	})(ctx)
}
