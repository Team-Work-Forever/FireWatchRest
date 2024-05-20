package entities

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pwd"
	"gorm.io/gorm"
)

type Auth struct {
	EntityBase
	Email            string `gorm:"column:email"`
	Password         string `gorm:"column:password"`
	Salt             string `gorm:"column:salt"`
	NIF              string `gorm:"column:nif"`
	IsAccountEnabled bool   `gorm:"column:is_account_enabled"`
}

func (e *Auth) TableName() string {
	return "auth_keys"
}

func NewAuth(email string, password string, nif string) *Auth {
	return &Auth{
		Email:    email,
		Password: password,
		Salt:     "assdads",
		NIF:      nif,
	}
}

func (u *Auth) BeforeCreate(tx *gorm.DB) error {
	u.EntityBase.BeforeCreate(tx)

	salt, err := pwd.GenerateSalt(pwd.BCRYPT_COST)

	if err != nil {
		return err
	}

	password, err := pwd.HashPassword(u.Password, salt)

	if err != nil {
		return err
	}

	u.Password = password
	u.Salt = salt

	return nil
}
