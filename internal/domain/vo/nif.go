package vo

import (
	"regexp"

	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type NIF struct {
	Value string `gorm:"column:nif"`
}

const NIF_REGEX = `(?:\b[0-9]{9}\b)`

func NewNIF(nif string) (*NIF, error) {
	nifRegex := regexp.MustCompile(NIF_REGEX)

	if nif == "" {
		return nil, exec.NIF_PROVIDE
	}

	if !nifRegex.MatchString(nif) {
		return nil, exec.NIF_PROVIDE_AN_VALID
	}

	return &NIF{
		Value: nif,
	}, nil
}

func (e *NIF) GetValue() string {
	return e.Value
}
