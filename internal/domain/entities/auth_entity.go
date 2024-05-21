package entities

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pwd"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
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

func (a *Auth) ChangePassword(password *vo.Password) error {
	if pwd.CheckPasswordHash(password.GetValue(), a.Salt, a.Password.GetValue()) {
		return exec.CANNOT_CHANGE_PASSWORD_AGAIN
	}

	salt, err := pwd.GenerateSalt(pwd.BCRYPT_COST)

	if err != nil {
		return err
	}

	hash, err := pwd.HashPassword(password.GetValue(), salt)

	if err != nil {
		return err
	}

	a.Password = vo.NewHash(hash)
	a.Salt = salt

	return nil
}

func (a *Auth) BeforeCreate(tx *gorm.DB) error {
	a.EntityBase.BeforeCreate(tx)

	salt, err := pwd.GenerateSalt(pwd.BCRYPT_COST)

	if err != nil {
		return err
	}

	password, err := pwd.HashPassword(a.Password.GetValue(), salt)

	if err != nil {
		return err
	}

	a.Password = vo.NewHash(password)
	a.Salt = salt

	return nil
}
