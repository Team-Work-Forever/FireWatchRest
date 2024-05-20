package vo

import (
	"errors"
	"strings"
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
		return nil, errors.New("empty phone")
	}

	if !validateCountryCode(code) {
		return nil, errors.New("invalid country code")
	}

	if shouldHaveLength(value, 9) {
		return nil, errors.New("phone number must be nine digits long")
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
