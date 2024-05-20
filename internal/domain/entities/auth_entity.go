package entities

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pwd"
	"gorm.io/gorm"
)

type Auth struct {
	EntityBase
	Email            vo.Email    `gorm:"embedded"`
	Password         vo.Password `gorm:"embedded"`
	Salt             string      `gorm:"column:salt"`
	NIF              vo.NIF      `gorm:"embedded"`
	IsAccountEnabled bool        `gorm:"column:is_account_enabled"`
}

func (e *Auth) TableName() string {
	return "auth_keys"
}

func NewAuth(email vo.Email, password vo.Password, nif vo.NIF) *Auth {
	return &Auth{
		Email:    email,
		Password: password,
		NIF:      nif,
	}
}

func (u *Auth) BeforeCreate(tx *gorm.DB) error {
	u.EntityBase.BeforeCreate(tx)

	salt, err := pwd.GenerateSalt(pwd.BCRYPT_COST)

	if err != nil {
		return err
	}

	password, err := pwd.HashPassword(u.Password.GetValue(), salt)

	if err != nil {
		return err
	}

	u.Password = vo.NewHash(password)
	u.Salt = salt

	return nil
}
