package exec

import (
	"github.com/gofiber/fiber/v2"
)

var (
	USER_ALREADY_EXISTS     = NewError("user-ae", fiber.StatusNotFound, "User Already Exists", "There is already an user with that email")
	USER_ALREADY_EXISTS_NIF = NewError("user-an", fiber.StatusConflict, "User Already Exists", "There is already an user with that NIF")

	USER_NOT_FOUND = NewError("user-nf", fiber.StatusNotFound, "User Not Found", "User not found")

	PASSWORD_WRONG       = NewError("password-wrg", fiber.StatusConflict, "Authentication", "The email or password is wrong")
	PASSWORDS_DONT_MATCH = NewError("password-d-mtch", fiber.StatusConflict, "Authentication", "The password don't match")

	CANNOT_CHANGE_PASSWORD_AGAIN = NewError("password-ch", fiber.StatusConflict, "Reset Password", "It's not possible to reset to the same password")

	TOKEN_NOT_FOUND = NewError("token-nf", fiber.StatusNotFound, "Token Not Found", "No token was found")

	TOKEN_ISNT_VALD       = NewError("token-isn-v", fiber.StatusForbidden, "Token Invalid", "Token is not valid")
	FAILED_FETCHING_TOKEN = NewError("token-f", fiber.StatusConflict, "Fetch Failed", "Failed fetching token")
)
