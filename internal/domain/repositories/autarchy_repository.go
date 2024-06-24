package repositories

import (
	"fmt"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/daos"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pagination"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
	"gorm.io/gorm"
)

type AutarchyRepository struct {
	dbContext *gorm.DB
}

func NewAutarchyRepository(database *gorm.DB) *AutarchyRepository {
	return &AutarchyRepository{
		dbContext: database,
	}
}

func (repo *AutarchyRepository) ExistsAutarchyWithTitle(title string) bool {
	if err := repo.dbContext.Where("lower(title) = lower(?)", title).First(&entities.Autarchy{}).Error; err != nil {
		return false
	}

	return true
}

func (repo *AutarchyRepository) GetAutarchtDetailById(autarchyId string) (*daos.AutarchyDetailsView, error) {
	var result *daos.AutarchyDetailsView

	if err := repo.dbContext.Where("id = ?", autarchyId).Where("deleted_at is null").First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *AutarchyRepository) GetAutarchyById(autarchyId string) (*entities.Autarchy, error) {
	var result *entities.Autarchy

	if err := repo.dbContext.Where("id = ?", autarchyId).Where("deleted_at is null").First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *AutarchyRepository) GetAutarchyByCity(city string) (*entities.Autarchy, error) {
	var result *entities.Autarchy

	if err := repo.dbContext.Where("address_city = ?", city).Where("deleted_at is null").First(&result).Error; err != nil {
		return nil, exec.AUTARCHY_NOT_FOUND
	}

	return result, nil
}

func (repo *AutarchyRepository) GetAutarchyBurnCount(autarchyId string) (int, error) {
	var burnRequests []entities.BurnRequest

	expr := repo.dbContext.Where("autarchy_id = ?", autarchyId)

	if err := expr.Find(&burnRequests).Error; err != nil {
		return 0, err
	}

	return len(burnRequests), nil
}

func (repo *AutarchyRepository) GetAll(params map[string]interface{}, pagination *pagination.Pagination) ([]daos.AutarchyDetailsView, error) {
	var result []daos.AutarchyDetailsView

	expr := repo.dbContext.Where("deleted_at is null")

	if search, ok := params["search"]; ok {
		expr.Where("lower(title) like lower(?)", fmt.Sprintf("%%%s%%", search))
	}

	expr = expr.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())

	if err := expr.Find(&result).Error; err != nil {
		return nil, err
	}

	pagination.SetTotalPages(len(result))
	return result, nil
}

func (repo *AutarchyRepository) Update(burn *entities.Autarchy) error {
	return repo.dbContext.Save(burn).Error
}

func (repo *AutarchyRepository) Delete(autarchy *entities.Autarchy) error {
	return repo.dbContext.Delete(autarchy).Error
}
