package repositories

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/daos"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
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

func (repo *ProfileRepository) GetAutarchyByAuthId(authId string) (*entities.Autarchy, error) {
	var autarchy *entities.Autarchy

	if err := repo.dbContext.Where("auth_key_id = ?", authId).First(&autarchy).Error; err != nil {
		return nil, err
	}

	return autarchy, nil
}

func (repo *ProfileRepository) Update(profile interface{}) error {
	return repo.dbContext.Save(profile).Error
}

func (repo *ProfileRepository) GetPublicProfile(email *vo.Email) (*daos.ProfileDao, error) {
	var user *daos.ProfileDao

	err := repo.dbContext.Table("users u").
		Select("u.user_name, u.profile_avatar, ak.email, ak.nif, u.phone_code, u.phone_number").
		Joins("inner join auth_keys ak on ak.id = u.auth_key_id").
		Where("ak.email = ?", email.Value).
		Scan(&user).Error

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, exec.USER_NOT_FOUND
	}

	return user, nil
}
