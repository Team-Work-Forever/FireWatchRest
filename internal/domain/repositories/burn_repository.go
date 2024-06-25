package repositories

import (
	"fmt"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/daos"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pagination"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BurnRepository struct {
	dbContext *gorm.DB
}

func NewBurnRepository(database *gorm.DB) *BurnRepository {
	return &BurnRepository{
		dbContext: database,
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
		request.AutarchyId,
		request.Burn.ID,
		request.InitialPropose,
	)

	if err := tx.Create(&burnRequest).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	state := burnRequest.SetState(request.State, request.InitialPropose)

	if err := tx.Create(&state).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return burnRequest, tx.Commit().Error
}

func (repo *BurnRepository) UserOwnsBurn(authId string, burnId string) bool {
	if err := repo.dbContext.Where("auth_key_id = ?", authId).Where("burn_id = ?", burnId).First(&entities.BurnRequest{}).Error; err != nil {
		return false
	}

	return true
}

func (repo *BurnRepository) GetBurnById(burnId string) (*entities.Burn, error) {
	var result *entities.Burn

	if err := repo.dbContext.Where("id = ?", burnId).Where("deleted_at is null").First(&result).Error; err != nil {
		return nil, err
	}

	coordinates, err := repo.GetCoordinates(burnId)

	if err != nil {
		return nil, err
	}

	result.Coordinates = *coordinates
	return result, nil
}

func (repo *BurnRepository) GetCoordinates(burnId string) (*vo.Coordinate, error) {
	var result *daos.BurnDetailsView

	if err := repo.dbContext.Select([]string{
		"lat",
		"lon",
	}).Where("id = ?", burnId).Where("deleted_at is null").First(&result).Error; err != nil {
		return nil, err
	}

	return vo.NewCoordinate(result.Lat, result.Lon), nil
}

func (repo *BurnRepository) Delete(burn *entities.Burn) error {
	if err := repo.dbContext.Where("burn_id = ?", burn.ID).Delete(&entities.BurnRequest{}).Error; err != nil {
		return err
	}

	if err := repo.dbContext.Where("burn_id = ?", burn.ID).Delete(&entities.BurnRequestState{}).Error; err != nil {
		return err
	}

	return repo.dbContext.Delete(burn).Error
}

func (repo *BurnRepository) GetBurnStatus(authId string, burnId string) (*uint16, error) {
	var burnRequest *entities.BurnRequestState

	if err := repo.dbContext.Where("burn_id = ?", burnId).Where("auth_key_id = ?", authId).First(&burnRequest).Error; err != nil {
		return nil, err
	}

	return &burnRequest.State, nil
}

func (repo *BurnRepository) SetBurnStatus(authId, autarchyId, burnId string, state vo.BurnRequestStates) (*entities.BurnRequestState, error) {
	tx := repo.dbContext.Begin()

	burnRequest := entities.NewBurnRequestState(
		authId,
		autarchyId,
		burnId,
		state,
		"Finish Burn",
	)

	if err := tx.Create(&burnRequest).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return burnRequest, nil
}

func (repo *BurnRepository) GetBurnDetailById(authId string, burnId string) (*daos.BurnDetailsView, error) {
	var result *daos.BurnDetailsView

	if err := repo.dbContext.Where("id = ?", burnId).Where("author = ?", authId).Where("deleted_at is null").First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *BurnRepository) GetAllBurns(authId string, params map[string]interface{}, pagination *pagination.Pagination) ([]daos.BurnDetailsView, error) {
	var result []daos.BurnDetailsView

	expr := repo.dbContext.Where("deleted_at is null")

	if autarchyId, ok := params["autarchyId"]; ok {
		expr.Where("autarchy_id = ?", autarchyId)
	} else {
		expr.Where("author = ?", authId)
	}

	if search, ok := params["search"]; ok {
		expr.Where("lower(title) like LOWER(?)", fmt.Sprintf("%%%s%%", search))
	}

	if state, ok := params["state"]; ok {
		expr.Where("state = ?", state)
	}

	if startDate, ok := params["start_date"]; ok {
		expr.Where("begin_at >= ?", startDate)
	}

	if endDate, ok := params["end_date"]; ok {
		expr.Where("begin_at < ?", endDate)
	}

	if sort, ok := params["sort"]; ok {
		var sortInput = sort == "desc"

		expr.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "begin_at"},
			Desc:   sortInput,
		})
	}

	expr = expr.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())

	if err := expr.Find(&result).Error; err != nil {
		return nil, err
	}

	pagination.SetTotalPages(len(result))
	return result, nil
}

func (repo *BurnRepository) Update(burn *entities.Burn) error {
	return repo.dbContext.Save(burn).Error
}
