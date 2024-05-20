package repositories

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/adapters"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"gorm.io/gorm"
)

type AuthRepository struct {
	dbContext *gorm.DB
}

func NewAuthRepository(database adapters.Database) *AuthRepository {
	return &AuthRepository{
		dbContext: database.GetDatabase(),
	}
}

func (repo *AuthRepository) GetAuthByEmail(email *vo.Email) (*entities.Auth, error) {
	var auth *entities.Auth

	if err := repo.dbContext.Where("email = ?", email.GetValue()).First(&auth).Error; err != nil {
		return nil, err
	}

	return auth, nil
}

func (repo *AuthRepository) ExistsUserWithEmail(email *vo.Email) bool {
	if err := repo.dbContext.Where("email = ?", email.GetValue()).First(&entities.Auth{}).Error; err != nil {
		return false
	}

	return true
}

func (repo *AuthRepository) ExistsUserWithNif(nif *vo.NIF) bool {
	if err := repo.dbContext.Where("nif = ?", nif.GetValue()).First(&entities.Auth{}).Error; err != nil {
		return false
	}

	return true
}

func (repo *AuthRepository) CreateAccount(auth *entities.Auth, user *entities.User) error {
	tx := repo.dbContext.Begin()

	if err := tx.Create(&auth).Error; err != nil {
		tx.Rollback()
		return err
	}

	// link auth keys to account
	user.AsignAuthKey(auth)
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (repo *AuthRepository) Update(auth *entities.Auth) error {
	return repo.dbContext.Save(auth).Error
}
