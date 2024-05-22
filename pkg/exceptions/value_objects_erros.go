package exec

import "github.com/gofiber/fiber/v2"

var (
	ADDRESS_PROVIDE_STREET          = NewError("address-pv-street", fiber.StatusBadRequest, "Adress Validation", "Provide street field")
	ADDRESS_PROVIDE_NUMBER          = NewError("address-pv-number", fiber.StatusBadRequest, "Adress Validation", "Provide street number field")
	ADDRESS_PROVIDE_AN_VALID_NUMBER = NewError("address-pv-number-v", fiber.StatusBadRequest, "Adress Validation", "Provide an valid street number field")
	ADDRESS_PROVIDE_CITY            = NewError("address-pv-city", fiber.StatusBadRequest, "Adress Validation", "Provide city field")

	EMAIL_PROVIDE          = NewError("email-pv", fiber.StatusBadRequest, "Email Validation", "Provide email field")
	EMAIL_PROVIDE_AN_VALID = NewError("email-pv-v", fiber.StatusBadRequest, "Email Validation", "Provide an valid email field")

	NIF_PROVIDE          = NewError("nif-pv", fiber.StatusBadRequest, "NIF Validation", "Provide NIF field")
	NIF_PROVIDE_AN_VALID = NewError("nif-pv-v", fiber.StatusBadRequest, "NIF Validation", "Provide an valid NIF field")

	PASSWORD_PROVIDE                  = NewError("password-pv", fiber.StatusBadRequest, "Password Validation", "Provide password field")
	PASSWORD_BTW_6_16                 = NewError("password-btw-6-16", fiber.StatusBadRequest, "Password Validation", "Password must be between 6 and 16 characters")
	PASSWORD_MUST_CONTAIN_ONE_NUMBER  = NewError("password-mst-number", fiber.StatusBadRequest, "Password Validation", "Password must contain at least one number")
	PASSWORD_MUST_CONTAIN_ONE_CAPITAL = NewError("password-mst-capital", fiber.StatusBadRequest, "Password Validation", "Password must contain at least one capital letter")
	PASSWORD_MUST_CONTAIN_NON_CAPITAL = NewError("password-mst-n-capital", fiber.StatusBadRequest, "Password Validation", "Password must contain at least one non-capital letter")

	PHONE_PROVIDE                = NewError("phone-pv", fiber.StatusBadRequest, "Phone Validation", "Provide phone number and country code field")
	PHONE_INVALID_COUNTRY_NUMBER = NewError("phone-i-country", fiber.StatusBadRequest, "Phone Validation", "Country Code is Invalid")
	PHONE_MUST_BE_NINE_DIGITS    = NewError("phone-mst-9", fiber.StatusBadRequest, "Phone Validation", "Phone number must be nine digits long")

	ZIP_CODE_PROVIDE = NewError("zip-code-pv", fiber.StatusBadRequest, "Zip Code Validation", "Provide zip code field")
	ZIP_CODE_INVALID = NewError("zip-code-i", fiber.StatusBadRequest, "Zip Code Validation", "Invalid Zip Code: format 4444-444")

	USER_NAME_PROVIDE  = NewError("usern-pv", fiber.StatusBadRequest, "User Validation", "Provide user name field")
	FIRST_NAME_PROVIDE = NewError("firstn-pv", fiber.StatusBadRequest, "User Validation", "Provide first name field")
	LAST_NAME_PROVIDE  = NewError("lastn-pv", fiber.StatusBadRequest, "User Validation", "Provide last name field")
)
