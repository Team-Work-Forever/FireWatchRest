package services

import (
	"errors"
	"strconv"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/api"
)

var (
	ErrNotValidAddress error = errors.New("address is not valid")
)

func GetAutarchy(address vo.Address) (string, error) {
	var okStreet, okNumber bool

	housing, err := api.GetCPHousing(address.ZipCode)

	if err != nil {
		return "", err
	}

	if address.ZipCode != housing.CP {
		return "", ErrNotValidAddress
	}

	if address.City != housing.Localidade {
		return "", ErrNotValidAddress
	}

	for _, value := range housing.Partes {
		if value.Arteria == address.Street {
			okStreet = true
			break
		}
	}

	if !okStreet {
		return "", ErrNotValidAddress
	}

	for _, value := range housing.Pontos {
		if value.Casa == strconv.Itoa(address.Number) {
			okNumber = true
			break
		}
	}

	if !okNumber {
		return "", ErrNotValidAddress
	}

	return housing.Municipio, nil
}
