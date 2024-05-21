package exec

import "github.com/gofiber/fiber/v2"

var (
	ADDRESS_PROVIDE_STREET          = NewError(fiber.StatusBadRequest, "Adress Validation", "Provide street field")
	ADDRESS_PROVIDE_NUMBER          = NewError(fiber.StatusBadRequest, "Adress Validation", "Provide street number field")
	ADDRESS_PROVIDE_AN_VALID_NUMBER = NewError(fiber.StatusBadRequest, "Adress Validation", "Provide an valid street number field")
	ADDRESS_PROVIDE_CITY            = NewError(fiber.StatusBadRequest, "Adress Validation", "Provide city field")

	EMAIL_PROVIDE          = NewError(fiber.StatusBadRequest, "Email Validation", "Provide email field")
	EMAIL_PROVIDE_AN_VALID = NewError(fiber.StatusBadRequest, "Email Validation", "Provide an valid email field")

	NIF_PROVIDE          = NewError(fiber.StatusBadRequest, "NIF Validation", "Provide NIF field")
	NIF_PROVIDE_AN_VALID = NewError(fiber.StatusBadRequest, "NIF Validation", "Provide an valid NIF field")

	PASSWORD_PROVIDE                  = NewError(fiber.StatusBadRequest, "Password Validation", "Provide password field")
	PASSWORD_BTW_6_16                 = NewError(fiber.StatusBadRequest, "Password Validation", "Password must be between 6 and 16 characters")
	PASSWORD_MUST_CONTAIN_ONE_NUMBER  = NewError(fiber.StatusBadRequest, "Password Validation", "Password must contain at least one number")
	PASSWORD_MUST_CONTAIN_ONE_CAPITAL = NewError(fiber.StatusBadRequest, "Password Validation", "Password must contain at least one capital letter")
	PASSWORD_MUST_CONTAIN_NON_CAPITAL = NewError(fiber.StatusBadRequest, "Password Validation", "Password must contain at least one non-capital letter")

	PHONE_PROVIDE                = NewError(fiber.StatusBadRequest, "Phone Validation", "Provide phone field")
	PHONE_INVALID_COUNTRY_NUMBER = NewError(fiber.StatusBadRequest, "Phone Validation", "Country Code is Invalid")
	PHONE_MUST_BE_NINE_DIGITS    = NewError(fiber.StatusBadRequest, "Phone Validation", "Phone number must be nine digits long")

	ZIP_CODE_PROVIDE = NewError(fiber.StatusBadRequest, "Zip Code Validation", "Provide phone field")
	ZIP_CODE_INVALID = NewError(fiber.StatusBadRequest, "Zip Code Validation", "Invalid Zip Code: format 4444-444")

	USER_NAME_PROVIDE  = NewError(fiber.StatusBadRequest, "User Validation", "Provide user name field")
	FIRST_NAME_PROVIDE = NewError(fiber.StatusBadRequest, "User Validation", "Provide first name field")
	LAST_NAME_PROVIDE  = NewError(fiber.StatusBadRequest, "User Validation", "Provide last name field")
)
