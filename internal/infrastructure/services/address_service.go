package services

import (
	"errors"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/api"
)

var (
	ErrNotValidAddress error = errors.New("address is not valid")
)

func GetAddress(lat, lon float64) (*vo.Address, error) {
	location, err := api.GetLocation(lat, lon)

	if err != nil {
		return nil, err
	}

	zipCode, err := vo.NewZipCode(location.CP)

	if err != nil {
		return nil, err
	}

	return vo.NewAddressWithEmptyValues(
		location.Rua,
		12,
		*zipCode,
		location.Concelho,
	)
}

func GetAutarchy(address vo.Address) (string, error) {
	var okStreet bool

	housing, err := api.GetCPHousing(address.ZipCode)

	if err != nil {
		return "", err
	}

	if address.ZipCode != housing.CP {
		return "", ErrNotValidAddress
	}

	if address.City != housing.Concelho {
		return "", ErrNotValidAddress
	}

	for _, value := range housing.Partes {
		if value.Arteria == address.Street {
			okStreet = true
			break
		}
	}

	if !okStreet && len(housing.Partes) != 0 {
		return "", ErrNotValidAddress
	}

	return housing.Municipio, nil
}
