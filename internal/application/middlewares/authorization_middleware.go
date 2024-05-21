package middlewares

import (
	"github.com/Team-Work-Forever/FireWatchRest/config"
	jwtService "github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/jwt"
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
	})(ctx)
}
