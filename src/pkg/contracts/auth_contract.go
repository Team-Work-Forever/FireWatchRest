package contracts

import "mime/multipart"

type DefaultResponse struct {
	Code  int
	Title string
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Email     string                `form:"email" binding:"required"`
	Password  string                `form:"password" binding:"required"`
	FirstName string                `form:"first_name" binding:"required"`
	LastName  string                `form:"last_name" binding:"required"`
	File      *multipart.FileHeader `form:"file" binding:"required" swaggerignore:"true"`
}
