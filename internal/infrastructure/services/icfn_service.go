package services

import (
	"errors"

	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/api"
)

var (
	ErrCannotDoBurn error = errors.New("is not possible to do burn is dangerous")
)

func CheckICFNIndex(lat, lon float64, hasAidTeam bool) bool {
	index, err := api.GetICNFIndex(lat, lon, hasAidTeam)

	if err != nil {
		return false
	}

	return index.RiscoOperacao != -1 && index.RiscoOperacao <= 50
}
