package contracts

import "mime/multipart"

type DefaultResponse struct {
	Code  int
	Title string
}

type LoginRequest struct {
	Email    string
	Password string
}

type SignUpRequest struct {
	Email     string                `form:"email" binding:"required"`
	Password  string                `form:"password" binding:"required"`
	FirstName string                `form:"first_name" binding:"required"`
	LastName  string                `form:"last_name" binding:"required"`
	File      *multipart.FileHeader `form:"file" binding:"required" swaggerignore:"true"`
}
