package repositories

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/daos"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	dbContext *gorm.DB
}

func NewProfileRepository(database *gorm.DB) *ProfileRepository {
	return &ProfileRepository{
		dbContext: database,
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

func (repo *ProfileRepository) GetPublicProfile(email *vo.Email) (*daos.ProfileDao, error) {
	var user *daos.ProfileDao

	err := repo.dbContext.Table("users u").
		Select("u.user_name, u.profile_avatar, ak.email").
		Joins("inner join auth_keys ak on ak.id = u.auth_key_id").
		Where("ak.email = ?", email.Value).
		Scan(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
