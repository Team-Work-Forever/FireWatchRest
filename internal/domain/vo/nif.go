package vo

import (
	"errors"
	"regexp"
)

type NIF struct {
	Value string `gorm:"column:nif"`
}

const NIF_REGEX = `(?:\b[0-9]{9}\b)`

func NewNIF(nif string) (*NIF, error) {
	nifRegex := regexp.MustCompile(NIF_REGEX)

	if nif == "" {
		return nil, errors.New("empty nif")
	}

	if !nifRegex.MatchString(nif) {
		return nil, errors.New("provide an valid nif")
	}

	return &NIF{
		Value: nif,
	}, nil
}

func (e *NIF) GetValue() string {
	return e.Value
}
