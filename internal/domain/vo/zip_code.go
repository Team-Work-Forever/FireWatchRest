package vo

import (
	"regexp"

	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

const ZIP_CODE_REGEX = `^\d{4}(-\d{3})?$`

type ZipCode struct {
	Value string
}

func NewZipCode(value string) (*ZipCode, error) {
	zipCodeRegex := regexp.MustCompile(ZIP_CODE_REGEX)

	if value == "" {
		return nil, exec.ZIP_CODE_PROVIDE
	}

	if !zipCodeRegex.MatchString(value) {
		return nil, exec.ZIP_CODE_INVALID
	}

	return &ZipCode{
		Value: value,
	}, nil
}

func (z *ZipCode) GetValue() string {
	return z.Value
}
