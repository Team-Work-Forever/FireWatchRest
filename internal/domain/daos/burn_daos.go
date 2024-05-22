package daos

import "github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"

type (
	CreateBurnDao struct {
		AuthId         string
		Burn           *entities.Burn
		InitialPropose string
	}
)
