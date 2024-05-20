package vo

import (
	"errors"
	"regexp"
)

const ZIP_CODE_REGEX = `^\d{4}(-\d{3})?$`

type ZipCode struct {
	Value string
}

func NewZipCode(value string) (*ZipCode, error) {
	zipCodeRegex := regexp.MustCompile(ZIP_CODE_REGEX)

	if value == "" {
		return nil, errors.New("empty postal code")
	}

	if !zipCodeRegex.MatchString(value) {
		return nil, errors.New("invalid postal code")
	}

	return &ZipCode{
		Value: value,
	}, nil
}

func (z *ZipCode) GetValue() string {
	return z.Value
}
