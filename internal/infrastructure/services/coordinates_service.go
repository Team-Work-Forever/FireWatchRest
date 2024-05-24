package services

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErroInvalidCoordinate error = errors.New("invalid coordinate")
)

func convertDMStoDD(degrees, minutes, seconds float64) float64 {
	return degrees + (minutes / 60.0) + (seconds / 3600.0)
}

func parseCoordinate(coordinate string) (float32, error) {
	var fields []float64 = make([]float64, 0)

	values := strings.Split(coordinate, ":")

	if len(values) > 3 {
		return 0, ErroInvalidCoordinate
	}

	for index := range values {
		value, err := strconv.ParseFloat(values[index], 64)

		if err != nil {
			return 0, ErroInvalidCoordinate
		}

		fields = append(fields, value)
	}

	return float32(convertDMStoDD(fields[0], fields[1], fields[2])), nil
}

func GetCoordinates(lat, lon string) (float32, float32, error) {
	conLat, err := parseCoordinate(lat)

	if err != nil {
		return 0, 0, err
	}

	conLon, err := parseCoordinate(lon)

	if err != nil {
		return 0, 0, err
	}

	return conLat, conLon, nil
}
