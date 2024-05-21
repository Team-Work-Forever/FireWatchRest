package vo

import (
	"regexp"

	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

const EMAIL_REGEX = `^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$`

type Email struct {
	Value string `gorm:"column:email"`
}

func NewEmail(email string) (*Email, error) {
	emailRegex := regexp.MustCompile(EMAIL_REGEX)

	if email == "" {
		return nil, exec.EMAIL_PROVIDE
	}

	if !emailRegex.MatchString(email) {
		return nil, exec.EMAIL_PROVIDE_AN_VALID
	}

	return &Email{
		Value: email,
	}, nil
}

func (e *Email) GetValue() string {
	return e.Value
}
