package repositories

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/adapters"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	dbContext *gorm.DB
}

func NewProfileRepository(database adapters.Database) *ProfileRepository {
	return &ProfileRepository{
		dbContext: database.GetDatabase(),
	}
}

func (repo *ProfileRepository) GetUserByAuthId(authId string) (*entities.User, error) {
	var user *entities.User

	if err := repo.dbContext.Where("auth_key_id = ?", authId).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *ProfileRepository) Update(profile *entities.User) error {
	return repo.dbContext.Save(profile).Error
}
