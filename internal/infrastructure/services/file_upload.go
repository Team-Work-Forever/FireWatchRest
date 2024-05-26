package services

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrNoFileProvided  error = errors.New("no file provided")
	ErrInvalidFileType error = errors.New("file type not allowed")

	allowedFileTypes = map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
	}
)

func isValidType(filePath string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))

	return allowedFileTypes[ext]
}

func GetFile(ctx *fiber.Ctx, field string, isOptional bool) (*multipart.FileHeader, error) {
	files, err := GetFiles(ctx, field, isOptional)

	if err != nil {
		return nil, err
	}

	if len(files) == 0 && isOptional {
		return nil, nil
	}

	return files[0], nil
}

func GetFiles(ctx *fiber.Ctx, field string, isOptional bool) ([]*multipart.FileHeader, error) {
	var files []*multipart.FileHeader

	form, err := ctx.MultipartForm()

	if err != nil {
		return nil, err
	}

	for _, file := range form.File[field] {
		if !isValidType(file.Filename) {
			return nil, ErrInvalidFileType
		}

		files = append(files, file)
	}

	if len(files) == 0 && !isOptional {
		return nil, ErrNoFileProvided
	}

	return files, nil
}
