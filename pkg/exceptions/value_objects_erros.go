package exec

import "github.com/gofiber/fiber/v2"

var (
	BOOLEAN_PROVIDE_AN_VALID    = NewError("boolean-pv", fiber.StatusBadRequest, "Boolean Validation", "Provide an valid boolean")
	DATE_PROVIDE_AN_VALID       = NewError("date-pv", fiber.StatusBadRequest, "Date Validation", "Provide an valid date yyyy-mm-dd")
	START_DATE_PROVIDE_AN_VALID = NewError("start-date-pv", fiber.StatusBadRequest, "Start Date Validation", "Provide an valid start date")
	TITLE_PROVIDE               = NewError("title-pv", fiber.StatusBadRequest, "Title Validation", "Provide title field")

	COORDINATES_NOT_FOUND = NewError("coordinates-n-notfound", fiber.StatusBadRequest, "Local Not Found with Coordinates", "There ins't any local with those coordinates")
	COORDINATES_INVALID   = NewError("coordinates-i", fiber.StatusBadRequest, "Local Validation", "Coordinate is Invalid")

	ADDRESS_PROVIDE_STREET          = NewError("address-pv-street", fiber.StatusBadRequest, "Address Validation", "Provide street field")
	ADDRESS_PROVIDE_NUMBER          = NewError("address-pv-number", fiber.StatusBadRequest, "Address Validation", "Provide street number field")
	ADDRESS_PROVIDE_AN_VALID_NUMBER = NewError("address-pv-number-v", fiber.StatusBadRequest, "Address Validation", "Provide an valid street number field")
	ADDRESS_PROVIDE_CITY            = NewError("address-pv-city", fiber.StatusBadRequest, "Address Validation", "Provide city field")
	ADDRESS_INVALID                 = NewError("address-i", fiber.StatusBadRequest, "Address Validation", "Address is Invalid")

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

	ZIP_CODE_PROVIDE   = NewError("zip-code-pv", fiber.StatusBadRequest, "Zip Code Validation", "Provide zip code field")
	ZIP_CODE_INVALID   = NewError("zip-code-i", fiber.StatusBadRequest, "Zip Code Validation", "Invalid Zip Code: format 4444-444")
	ZIP_CODE_NOT_FOUND = NewError("zip-code-n-notfound", fiber.StatusBadRequest, "Zip Code Not Found", "There ins't any local with that zip code")

	USER_NAME_PROVIDE  = NewError("usern-pv", fiber.StatusBadRequest, "User Validation", "Provide user name field")
	FIRST_NAME_PROVIDE = NewError("firstn-pv", fiber.StatusBadRequest, "User Validation", "Provide first name field")
	LAST_NAME_PROVIDE  = NewError("lastn-pv", fiber.StatusBadRequest, "User Validation", "Provide last name field")

	DENIAL_OF_ACCESS = NewError("doa", fiber.StatusBadRequest, "Access Denied", "Don't have access to this action")

	TRY_AGAIN = NewError("ta", fiber.StatusBadRequest, "Try Again", "Please try again")

	STATE_NOT_VALID = NewError("snv", fiber.StatusBadRequest, "State not valid", "State is invalid")

	NOT_ABLE_CONVERT = NewError("n-ableconvert", fiber.StatusBadRequest, "Not Able Conversion", "Cannot Convert")

	QUERY_PARAMETER_SORT_NOT_VALID = NewError("query-sort-n-valid", fiber.StatusBadRequest, "Sort query parameter invalid", "That is not an valid sort query parameter")

	AUTARCHY_PROVIDE_NAME         = NewError("autarchy-pv-name", fiber.StatusBadRequest, "Autarchy Validation", "Provide autarchy name field")
	AUTARCHY_PROVIDE_LAT          = NewError("autarchy-pv-lat", fiber.StatusBadRequest, "Autarchy Validation", "Provide autarchy a valid lat field")
	AUTARCHY_PROVIDE_LON          = NewError("autarchy-pv-lon", fiber.StatusBadRequest, "Autarchy Validation", "Provide autarchy a valid lon field")
	AUTARCHY_ALREADY_REGISTERED   = NewError("autarchy-fail-registered", fiber.StatusBadRequest, "Autarchy Already Registered", "That autarchy is already registered")
	AUTARCHY_FAILED_REMOVAL       = NewError("autarchy-fail-removal", fiber.StatusBadRequest, "Autarchy Failed Removal", "The autarchy failed to be removed")
	AUTARCHY_FAILED_DETAILS_FETCH = NewError("autarchy-fail-fetchdetais", fiber.StatusBadRequest, "Autarchy Failed Details Fetch", "Could not fetch autarchy details")
	AUTARCHY_NOT_FOUND            = NewError("autarchy-n-notfound", fiber.StatusBadRequest, "Autarchy Not Found", "Autarchy not found")
	AUTARCHY_NOT_ABLE_UPDATE      = NewError("autarchy-n-ableupdate", fiber.StatusBadRequest, "Autarchy Not Able to Update", "Autarchy could not be updated")

	BURN_PROVIDE_NOT_EXISTING_REASON = NewError("burn-pv-ne-reason", fiber.StatusBadRequest, "Burn Validation", "Provide a valid reason type field")
	BURN_PROVIDE_NOT_EXISTING_TYPE   = NewError("burn-pv-ne-type", fiber.StatusBadRequest, "Burn Validation", "Provide a valid burn type field")
	BURN_DENIAL_OF_ACCESS            = NewError("burn-doa", fiber.StatusBadRequest, "Access Denied", "Don't have access to this burn")
	BURN_FAILED_SCHEDULED_ACTION     = NewError("burn-fail-schedualaction", fiber.StatusBadRequest, "Burn Failed Schedule Action", "Burn is not scheduled, it cannot perform this action")
	BURN_FAILED_REMOVAL              = NewError("burn-fail-removal", fiber.StatusBadRequest, "Burn Failed Removal", "The burn failed to be removed")
	BURN_FAILED_FETCH                = NewError("burn-fail_fetch", fiber.StatusBadRequest, "Burn Failed Fetch", "Could not fetch the status")
	BURN_NOT_FOUND                   = NewError("burn-n-notfound", fiber.StatusBadRequest, "Burn Not Found", "Burn not found")
	BURN_NOT_ABLE_REMOVAL            = NewError("burn-n-ableremoval", fiber.StatusBadRequest, "Burn Not Able to Remove", "Burn could not be removed while in action")
	BURN_NOT_ABLE_UPDATE             = NewError("burn-n-ableupdate", fiber.StatusBadRequest, "Burn Not Able to Update", "Burn could not be updated")
	BURN_NOT_ABLE_START              = NewError("burn-n-ablestart", fiber.StatusBadRequest, "Burn Not Able to Start", "Cannot start burn")
	BURN_NOT_ABLE_COMPLETE           = NewError("burn-n-complete", fiber.StatusBadRequest, "Burn Not Able to Complete", "Cannot complete burn")
	BURN_NOT_ACTIVE                  = NewError("burn-n-active", fiber.StatusBadRequest, "Burn Not Active", "Burn is not active, it cannot perform this action")

	FILE_PROVIDE                   = NewError("file-pv", fiber.StatusBadRequest, "File Validation", "No file provided")
	FILE_PROVIDE_NOT_EXISTING_TYPE = NewError("file-pv-ne-type", fiber.StatusBadRequest, "File Validation", "Provide a valid file type field")
	FILE_NOT_ABLE_UPLOAD           = NewError("file-n-ableupload", fiber.StatusBadRequest, "File Not Able to Upload", "Error uploading file")
)
