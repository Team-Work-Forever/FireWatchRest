package date

import (
	"time"
)

func Now() (*time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02", time.Now().Truncate(24*time.Hour).Format("2006-01-02"))

	if err != nil {
		return nil, err
	}

	return &parsedDate, nil
}

func ParseString(date string) (*time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02", date)

	if err != nil {
		return nil, err
	}

	return &parsedDate, nil
}
