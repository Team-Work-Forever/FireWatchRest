package vo

import (
	"unicode"

	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
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
		return nil, exec.PASSWORD_PROVIDE
	}

	if len(password) < 6 || len(password) > 16 {
		return nil, exec.PASSWORD_BTW_6_16
	}

	if !containsNumber(password) {
		return nil, exec.PASSWORD_MUST_CONTAIN_ONE_NUMBER
	}

	if !containsCapitalLetter(password) {
		return nil, exec.PASSWORD_MUST_CONTAIN_ONE_CAPITAL
	}

	if !containsNonCapitalLetter(password) {
		return nil, exec.PASSWORD_MUST_CONTAIN_NON_CAPITAL
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
