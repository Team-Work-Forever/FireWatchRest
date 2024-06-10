package entities

import (
	"errors"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
)

type Autarchy struct {
	IdentityUser
	Title       string        `gorm:"column:title"`
	Coordinates vo.Coordinate `gorm:"column:geo_location;type:geometry"`
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
		Coordinates: coordinates,
		IdentityUser: IdentityUser{
			Picture:     picture,
			PhoneNumber: phone,
			Address:     address,
		},
	}, nil
}

func (a *Autarchy) TableName() string {
	return "autarchy"
}

func (e *Autarchy) AsignAuthKey(auth *Auth) {
	e.AuthKeyId = auth.ID
}
