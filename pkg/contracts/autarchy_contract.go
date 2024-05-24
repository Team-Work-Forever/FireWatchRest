package contracts

import "mime/multipart"

type (
	CreateAutarchyRequest struct {
		Title       string                `form:"title" binding:"required"`
		Email       string                `form:"email" binding:"required"`
		NIF         string                `form:"nif" binding:"required"`
		Password    string                `form:"password" binding:"required"`
		PhoneCode   string                `form:"phone_code" binding:"required"`
		PhoneNumber string                `form:"phone_number" binding:"required"`
		Street      string                `form:"street" binding:"required"`
		StreetPort  int                   `form:"street_port" binding:"required"`
		ZipCode     string                `form:"zip_code" binding:"required"`
		City        string                `form:"city" binding:"required"`
		Lon         float32               `form:"lon" binding:"required"`
		Lat         float32               `form:"lat" binding:"required"`
		Avatar      *multipart.FileHeader `form:"avatar" binding:"required" swaggerignore:"true"`
	}

	AutarchyResponse struct {
		Title       string  `json:"title"`
		Email       string  `json:"email"`
		PhoneCode   string  `json:"phone_code"`
		PhoneNumber string  `json:"phone_number"`
		Street      string  `json:"street"`
		StreetPort  int     `json:"street_port"`
		ZipCode     string  `json:"zip_code"`
		City        string  `json:"city"`
		Lon         float32 `json:"lon"`
		Lat         float32 `json:"lat"`
		Avatar      string  `json:"avatar"`
	}

	GetAutarchyRequest struct {
		AutarchyId string
	}

	AutarchyActionResponse struct {
		AutarchyId string `json:"autarchy_id"`
	}
)
