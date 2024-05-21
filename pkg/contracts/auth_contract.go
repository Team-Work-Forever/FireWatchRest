package contracts

import "mime/multipart"

type (
	DefaultResponse struct {
		Code  int
		Title string
	}

	AuthResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignUpRequest struct {
		Email       string                `form:"email" binding:"required"`
		NIF         string                `form:"nif" binding:"required"`
		Password    string                `form:"password" binding:"required"`
		FirstName   string                `form:"first_name" binding:"required"`
		LastName    string                `form:"last_name" binding:"required"`
		PhoneCode   string                `form:"phone_code" binding:"required"`
		PhoneNumber string                `form:"phone_number" binding:"required"`
		Street      string                `form:"street" binding:"required"`
		StreetPort  int                   `form:"street_port" binding:"required"`
		ZipCode     string                `form:"zip_code" binding:"required"`
		City        string                `form:"city" binding:"required"`
		File        *multipart.FileHeader `form:"avatar" binding:"required" swaggerignore:"true"`
	}

	ForgotPasswordRequest struct {
		Email string
	}

	ResetPasswordRequest struct {
		ForgotToken     string `swaggerignore:"true"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
	}

	RefreshTokensRequest struct {
		RefreshToken string
	}
)
