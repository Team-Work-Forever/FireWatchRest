package exec

import "github.com/gofiber/fiber/v2"

var (
	ADDRESS_PROVIDE_STREET          = fiber.NewError(fiber.StatusBadRequest, "Provide street field")
	ADDRESS_PROVIDE_NUMBER          = fiber.NewError(fiber.StatusBadRequest, "Provide street number field")
	ADDRESS_PROVIDE_AN_VALID_NUMBER = fiber.NewError(fiber.StatusBadRequest, "Provide an valid street number field")
	ADDRESS_PROVIDE_CITY            = fiber.NewError(fiber.StatusBadRequest, "Provide city field")

	EMAIL_PROVIDE          = fiber.NewError(fiber.StatusBadRequest, "Provide email field")
	EMAIL_PROVIDE_AN_VALID = fiber.NewError(fiber.StatusBadRequest, "Provide an valid email field")

	NIF_PROVIDE          = fiber.NewError(fiber.StatusBadRequest, "Provide NIF field")
	NIF_PROVIDE_AN_VALID = fiber.NewError(fiber.StatusBadRequest, "Provide an valid NIF field")

	PASSWORD_PROVIDE                  = fiber.NewError(fiber.StatusBadRequest, "Provide password field")
	PASSWORD_BTW_6_16                 = fiber.NewError(fiber.StatusBadRequest, "Password must be between 6 and 16 characters")
	PASSWORD_MUST_CONTAIN_ONE_NUMBER  = fiber.NewError(fiber.StatusBadRequest, "Password must contain at least one number")
	PASSWORD_MUST_CONTAIN_ONE_CAPITAL = fiber.NewError(fiber.StatusBadRequest, "Password must contain at least one capital letter")
	PASSWORD_MUST_CONTAIN_NON_CAPITAL = fiber.NewError(fiber.StatusBadRequest, "Password must contain at least one non-capital letter")

	PHONE_PROVIDE                = fiber.NewError(fiber.StatusBadRequest, "Provide phone field")
	PHONE_INVALID_COUNTRY_NUMBER = fiber.NewError(fiber.StatusBadRequest, "Country Code is Invalid")
	PHONE_MUST_BE_NINE_DIGITS    = fiber.NewError(fiber.StatusBadRequest, "Phone number must be nine digits long")

	ZIP_CODE_PROVIDE = fiber.NewError(fiber.StatusBadRequest, "Provide phone field")
	ZIP_CODE_INVALID = fiber.NewError(fiber.StatusBadRequest, "Invalid Zip Code: format 4444-444")

	FIRST_NAME_PROVIDE = fiber.NewError(fiber.StatusBadRequest, "Provide first name field")
	LAST_NAME_PROVIDE  = fiber.NewError(fiber.StatusBadRequest, "Provide last name field")
)
