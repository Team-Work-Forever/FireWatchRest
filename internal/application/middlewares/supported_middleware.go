package middlewares

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func shouldAccept(ctx *fiber.Ctx, contentTypeMime string) error {
	contentType := strings.Split(ctx.Get("Content-Type"), ";")[0]

	if contentType != contentTypeMime {
		return fiber.NewError(fiber.StatusUnsupportedMediaType, fmt.Sprintf("Content-Type must be %s", contentTypeMime))
	}

	return ctx.Next()
}

func ShouldAcceptJson(ctx *fiber.Ctx) error {
	return shouldAccept(ctx, fiber.MIMEApplicationJSON)
}

func ShouldAcceptMultiPart(ctx *fiber.Ctx) error {
	return shouldAccept(ctx, fiber.MIMEMultipartForm)
}
