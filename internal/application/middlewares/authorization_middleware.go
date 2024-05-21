package middlewares

import (
	"github.com/Team-Work-Forever/FireWatchRest/config"
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
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return shared.WriteProblemDetails(ctx, exec.Error{
				Title:  "Authorization Failed",
				Detail: err.Error(),
				Status: fiber.StatusUnauthorized,
			})
		},
	})(ctx)
}
