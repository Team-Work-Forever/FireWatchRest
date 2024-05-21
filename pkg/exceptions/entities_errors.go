package exec

import (
	"github.com/gofiber/fiber/v2"
)

var (
	USER_ALREADY_EXISTS          = fiber.NewError(fiber.StatusConflict, "There is already an user with that email")
	USER_ALREADY_EXISTS_NIF      = fiber.NewError(fiber.StatusConflict, "There is already an user with that NIF")
	USER_NOT_FOUND               = fiber.NewError(fiber.StatusConflict, "No user was found")
	PASSWORD_WRONG               = fiber.NewError(fiber.StatusConflict, "The email or password is wrong")
	PASSWORDS_DONT_MATCH         = fiber.NewError(fiber.StatusConflict, "The password don't match")
	CANNOT_CHANGE_PASSWORD_AGAIN = fiber.NewError(fiber.StatusConflict, "It's not possible to reset to the same password")

	TOKEN_NOT_FOUND       = fiber.NewError(fiber.StatusConflict, "No token was found")
	TOKEN_ISNT_VALD       = fiber.NewError(fiber.StatusConflict, "Token is not valid")
	FAILED_FETCHING_TOKEN = fiber.NewError(fiber.StatusConflict, "Failed fetching token")
)
