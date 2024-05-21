package middlewares

import (
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/shared"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	problem, ok := err.(*exec.Error)

	if !ok {
		return shared.WriteProblemDetails(
			ctx,
			exec.Error{
				Title:  "Internal Server Error",
				Status: fiber.StatusInternalServerError,
				Detail: err.Error(),
			},
		)
	}

	return shared.WriteProblemDetails(ctx, *problem)
}
