package repositories

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
	"gorm.io/gorm"
)

type AuthRepository struct {
	dbContext *gorm.DB
}

func NewAuthRepository(database *gorm.DB) *AuthRepository {
	return &AuthRepository{
		dbContext: database,
	}
}

func (repo *AuthRepository) GetAuthById(authId string) (*entities.Auth, error) {
	var auth *entities.Auth

	if err := repo.dbContext.Where("id = ?", authId).First(&auth).Error; err != nil {
		return nil, err
	}

	return auth, nil
}

func (repo *AuthRepository) GetAuthByEmail(email *vo.Email) (*entities.Auth, error) {
	var auth *entities.Auth

	if err := repo.dbContext.Where("email = ?", email.GetValue()).First(&auth).Error; err != nil {
		return nil, err
	}

	return auth, nil
}

func (repo *AuthRepository) ExistsAuthById(id string) bool {
	if err := repo.dbContext.Where("id = ?", id).First(&entities.Auth{}).Error; err != nil {
		return false
	}

	return true
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

func (repo *AuthRepository) CreateAccount(auth *entities.Auth, user interface{}) error {
	tx := repo.dbContext.Begin()

	if err := tx.Create(&auth).Error; err != nil {
		tx.Rollback()
		return err
	}

	switch userType := user.(type) {
	case *entities.User:
		userType.AsignAuthKey(auth)
		if err := tx.Create(&userType).Error; err != nil {
			tx.Rollback()
			return err
		}
	case *entities.Autarchy:
		userType.AsignAuthKey(auth)
		if err := tx.Create(&userType).Error; err != nil {
			tx.Rollback()
			return err
		}
	default:
		return exec.ACCOUNT_USER_TYPE_UNDEFINED
	}

	return tx.Commit().Error
}

func (repo *AuthRepository) Update(auth *entities.Auth) error {
	return repo.dbContext.Save(auth).Error
}
