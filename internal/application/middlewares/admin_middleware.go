package middlewares

import (
	"errors"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/shared"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrUnAuthorized error = errors.New("don't have access to this action")
)

func AdminMiddleware(ctx *fiber.Ctx) error {
	role := shared.GetRole(ctx)

	roleId, ok := vo.GetUserTypeKey(role)

	if !ok {
		return shared.WriteProblemDetails(ctx, exec.Error{
			Title:  "Authorization Failed",
			Detail: ErrUnAuthorized.Error(),
			Status: fiber.StatusUnauthorized,
		})
	}

	if roleId != uint16(vo.Admin) {
		return shared.WriteProblemDetails(ctx, exec.Error{
			Title:  "Authorization Failed",
			Detail: ErrUnAuthorized.Error(),
			Status: fiber.StatusUnauthorized,
		})
	}

	return nil
}
