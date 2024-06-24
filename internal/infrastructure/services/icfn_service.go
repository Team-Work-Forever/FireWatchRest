package services

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/api"
)

func CheckICFNIndex(lat, lon float64, hasAidTeam, ignore bool) bool {
	if ignore {
		return true
	}

	index, err := api.GetICNFIndex(lat, lon, hasAidTeam)

	if err != nil {
		return false
	}

	return index.RiscoOperacao != -1 && index.RiscoOperacao <= 50
}
