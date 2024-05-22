package repositories

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/adapters"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/daos"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"gorm.io/gorm"
)

type BurnRepository struct {
	dbContext *gorm.DB
}

func NewBurnRepository(database adapters.Database) *BurnRepository {
	return &BurnRepository{
		dbContext: database.GetDatabase(),
	}
}

func (repo *BurnRepository) CreateBurn(request daos.CreateBurnDao) (*entities.BurnRequest, error) {
	tx := repo.dbContext.Begin()

	if err := tx.Create(&request.Burn).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	burnRequest := entities.NewBurnRequest(
		request.AuthId,
		request.Burn.ID,
		request.InitialPropose,
	)

	if err := tx.Create(&burnRequest).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	state := burnRequest.SetState(vo.Scheduled, request.InitialPropose)

	if err := tx.Create(&state).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return burnRequest, tx.Commit().Error
}
