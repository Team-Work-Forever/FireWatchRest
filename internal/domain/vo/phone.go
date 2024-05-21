package vo

import (
	"strings"

	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type Phone struct {
	CountryCode string `gorm:"column:phone_code"`
	Number      string `gorm:"column:phone_number"`
}

func shouldNotBeNull(code string, value string) bool {
	return code == "" || value == ""
}

func validateCountryCode(code string) bool {
	return strings.Contains(code, "+") && len(code) <= 4
}

func shouldHaveLength(value string, length int) bool {
	return len(value) != length
}

func NewPhone(code string, value string) (*Phone, error) {
	if shouldNotBeNull(code, value) {
		return nil, exec.PHONE_PROVIDE
	}

	if !validateCountryCode(code) {
		return nil, exec.PHONE_INVALID_COUNTRY_NUMBER
	}

	if shouldHaveLength(value, 9) {
		return nil, exec.PHONE_MUST_BE_NINE_DIGITS
	}

	return &Phone{
		CountryCode: code,
		Number:      value,
	}, nil
}

func (p *Phone) GetCountryCode() string {
	return p.CountryCode
}

func (p *Phone) GetNumber() string {
	return p.Number
}
