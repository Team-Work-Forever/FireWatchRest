package repositories

import (
	"database/sql"

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

func (repo *BurnRepository) GetBurnById(authId string, burnId string) (*daos.GetBurnDao, error) {
	var result *daos.GetBurnDao

	sqlQuery := `
		SELECT
			b.id,
			b.title,
			b.map_picture,
			ST_X(geo_location)::float AS lat,
			ST_Y(geo_location)::float AS lon,
			b.has_aid_team,
			b.reason,
			b."type",
			b.begin_at,
			b.completed_at,
			brs.state
		FROM
				burn b
		inner join burn_requests_states brs
			on brs.burn_id = b.id and brs.auth_key_id = @authId
		inner join burn_requests br 
			on br.burn_id = b.id and br.auth_key_id = @authId
	`

	if err := repo.dbContext.Raw(sqlQuery, sql.Named("authId", authId)).Where("id = ?", burnId).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
