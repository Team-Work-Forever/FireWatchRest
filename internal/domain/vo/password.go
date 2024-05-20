package vo

import (
	"errors"
	"unicode"
)

type Password struct {
	Value string `gorm:"column:password"`
}

func containsNumber(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}

	return false
}

func containsCapitalLetter(s string) bool {
	for _, char := range s {
		if unicode.IsUpper(char) {
			return true
		}
	}

	return false
}

func containsNonCapitalLetter(s string) bool {
	for _, char := range s {
		if unicode.IsLower(char) {
			return true
		}
	}

	return false
}
func NewPassword(password string) (*Password, error) {
	if password == "" {
		return nil, errors.New("empty password")
	}

	if len(password) < 6 || len(password) > 16 {
		return nil, errors.New("password must be between 6 and 16 characters")
	}

	if !containsNumber(password) {
		return nil, errors.New("password must contain at least one number")
	}

	if !containsCapitalLetter(password) {
		return nil, errors.New("password must contain at least one capital letter")
	}

	if !containsNonCapitalLetter(password) {
		return nil, errors.New("password must contain at least one non-capital letter")
	}

	return &Password{
		Value: password,
	}, nil
}

func NewHash(password string) Password {
	return Password{
		Value: password,
	}
}

func (e *Password) GetValue() string {
	return e.Value
}
