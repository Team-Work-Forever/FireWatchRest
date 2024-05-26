package entities

import (
	"errors"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
)

type Autarchy struct {
	EntityBase
	AuthKeyId   string        `gorm:"column:auth_key_id"`
	Title       string        `gorm:"column:title"`
	Coordinates vo.Coordinate `gorm:"column:geo_location;type:geometry"`
	PhoneNumber vo.Phone      `gorm:"embedded"`
	Address     vo.Address    `gorm:"embedded"`
	Picture     string        `gorm:"column:autarchy_avatar"`
}

func NewAutarchy(
	title string,
	picture string,
	coordinates vo.Coordinate,
	phone vo.Phone,
	address vo.Address,
) (*Autarchy, error) {
	if title == "" {
		return nil, errors.New("title is not provided")
	}

	return &Autarchy{
		Title:       title,
		Picture:     picture,
		Coordinates: coordinates,
		PhoneNumber: phone,
		Address:     address,
	}, nil
}

func (a *Autarchy) TableName() string {
	return "autarchy"
}

func (e *Autarchy) AsignAuthKey(auth *Auth) {
	e.AuthKeyId = auth.ID
}
