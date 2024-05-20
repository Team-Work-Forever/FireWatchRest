package contracts

import "github.com/gofiber/fiber/v2"

type HttpInternalError struct {
	Code  int
	Title string
}

func (err *HttpInternalError) NewHttpInternalError(title string) *HttpInternalError {
	return &HttpInternalError{
		Code:  fiber.StatusInternalServerError,
		Title: title,
	}
}
