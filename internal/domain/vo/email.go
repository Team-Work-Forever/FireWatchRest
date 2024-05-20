package vo

import (
	"errors"
	"regexp"
)

const EMAIL_REGEX = `^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$`

type Email struct {
	Value string `gorm:"column:email"`
}

func NewEmail(email string) (*Email, error) {
	emailRegex := regexp.MustCompile(EMAIL_REGEX)

	if email == "" {
		return nil, errors.New("empty email")
	}

	if !emailRegex.MatchString(email) {
		return nil, errors.New("provide an valid email")
	}

	return &Email{
		Value: email,
	}, nil
}

func (e *Email) GetValue() string {
	return e.Value
}
