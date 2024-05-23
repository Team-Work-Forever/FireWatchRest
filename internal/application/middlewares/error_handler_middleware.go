package middlewares

import (
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/shared"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case *exec.Error:
		return shared.WriteProblemDetails(ctx, *e)
	case *fiber.Error:
		return shared.WriteProblemDetails(
			ctx,
			exec.Error{
				Title:  "Bad Input",
				Status: e.Code,
				Detail: e.Message,
			},
		)
	default:
		return shared.WriteProblemDetails(
			ctx,
			exec.Error{
				Title:  "Internal Server Error",
				Status: fiber.StatusInternalServerError,
				Detail: err.Error(),
			},
		)
	}

}
