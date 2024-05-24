package repositories

import (
	"fmt"

	"github.com/Team-Work-Forever/FireWatchRest/internal/adapters"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/daos"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pagination"
	"gorm.io/gorm"
)

type AutarchyRepository struct {
	dbContext *gorm.DB
}

func NewAutarchyRepository(database adapters.Database) *AutarchyRepository {
	return &AutarchyRepository{
		dbContext: database.GetDatabase(),
	}
}

func (repo *AutarchyRepository) ExistsAutarchyWithTitle(title string) bool {
	if err := repo.dbContext.Where("title = ?", title).First(&entities.Autarchy{}).Error; err != nil {
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

func (repo *AutarchyRepository) GetAll(params map[string]interface{}, pagination *pagination.Pagination) ([]daos.AutarchyDetailsView, error) {
	var result []daos.AutarchyDetailsView

	expr := repo.dbContext.Where("deleted_at is null")

	if search, ok := params["search"]; ok {
		expr.Where("title like ?", fmt.Sprintf("%%%s%%", search))
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
