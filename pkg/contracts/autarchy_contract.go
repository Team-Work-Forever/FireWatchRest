package contracts

import "mime/multipart"

type (
	CreateAutarchyRequest struct {
		Title string `form:"title" binding:"required"`
		Email string `form:"email" binding:"required"`
		NIF   string `form:"nif" binding:"required"`
		// Password    string                `form:"password" binding:"required"`
		PhoneCode   string                `form:"phone_code" binding:"required"`
		PhoneNumber string                `form:"phone_number" binding:"required"`
		Street      string                `form:"street" binding:"required"`
		StreetPort  int                   `form:"street_port" binding:"required"`
		ZipCode     string                `form:"zip_code" binding:"required"`
		City        string                `form:"city" binding:"required"`
		Lon         float64               `form:"lon" binding:"required"`
		Lat         float64               `form:"lat" binding:"required"`
		Avatar      *multipart.FileHeader `form:"avatar" binding:"required" swaggerignore:"true"`
	}

	AutarchyResponse struct {
		Id      string          `json:"id"`
		Title   string          `json:"title"`
		Email   string          `json:"email"`
		Phone   PhoneResponse   `json:"phone"`
		Address AddressResponse `json:"address"`
		Avatar  string          `json:"avatar"`
	}

	GetAutarchyRequest struct {
		AutarchyId string
	}

	GetAllAutarchiesRequest struct {
		Search   string
		PageSize uint64
		Page     uint64
	}

	DeleteAutarchyRequest struct {
		UserId     string
		AutarchyId string
	}

	UpdateAutarchyRequest struct {
		UserId      string                `swaggerignore:"true"`
		AutarchyId  string                `swaggerignore:"true"`
		Title       string                `form:"title" binding:"required"`
		Lat         string                `form:"lat" binding:"required"`
		Lon         string                `form:"lon" binding:"required"`
		Email       string                `form:"email" binding:"required"`
		PhoneCode   string                `form:"phone_code" binding:"required"`
		PhoneNumber string                `form:"phone_number" binding:"required"`
		Street      string                `form:"street" binding:"required"`
		StreetPort  *int                  `form:"street_port" binding:"required"`
		ZipCode     string                `form:"zip_code" binding:"required"`
		City        string                `form:"city" binding:"required"`
		Avatar      *multipart.FileHeader `form:"avatar" binding:"required" swaggerignore:"true"`
	}

	AutarchyActionResponse struct {
		AutarchyId string `json:"autarchy_id"`
	}
)
