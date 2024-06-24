package entities

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
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
		return nil, exec.TITLE_PROVIDE
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
