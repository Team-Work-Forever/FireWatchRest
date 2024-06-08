package contracts

import "mime/multipart"

type (
	ProfileResponse struct {
		Id        string          `json:"id"`
		Email     string          `json:"email"`
		Avatar    string          `json:"avatar"`
		UserName  string          `json:"user_name"`
		FirstName string          `json:"first_name"`
		LastName  string          `json:"last_name"`
		Phone     PhoneResponse   `json:"phone"`
		Address   AddressResponse `json:"address"`
		UserType  string          `json:"user_type"`
	}

	WhoamiRequest struct {
		UserId string
	}

	UpdateProfileResponse struct {
		UserId      string                `swaggerignore:"true"`
		Email       string                `form:"email" binding:"required"`
		UserName    string                `form:"user_name" binding:"required"`
		PhoneCode   string                `form:"phone_code" binding:"required"`
		PhoneNumber string                `form:"phone_number" binding:"required"`
		Street      string                `form:"street" binding:"required"`
		StreetPort  *int                  `form:"street_port" binding:"required"`
		ZipCode     string                `form:"zip_code" binding:"required"`
		City        string                `form:"city" binding:"required"`
		Avatar      *multipart.FileHeader `form:"avatar" binding:"required" swaggerignore:"true"`
	}
)
