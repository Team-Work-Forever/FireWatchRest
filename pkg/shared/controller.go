package shared

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	jwtService "github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/jwt"
)

type Controller interface {
	Route(router fiber.Router)
}

func GetUserId(ctx *fiber.Ctx) string {
	user := ctx.Locals("user").(*jwt.Token)
	claims, _ := user.Claims.(*jwtService.AuthClaims)

	return claims.Subject
}
