package repositories

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/adapters"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
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
