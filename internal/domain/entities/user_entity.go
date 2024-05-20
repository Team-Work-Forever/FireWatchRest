package entities

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
)

type User struct {
	EntityBase
	AuthKeyId     string     `gorm:"column:auth_key_id"`
	ProfileAvatar string     `gorm:"column:profile_avatar"`
	FirstName     string     `gorm:"column:first_name"`
	LastName      string     `gorm:"column:last_name"`
	PhoneNumber   vo.Phone   `gorm:"embedded"`
	Address       vo.Address `gorm:"embedded"`
	UserType      int        `gorm:"column:user_type"`
}

func NewUser(avatar string, firstName string, lastName string, phone vo.Phone, address vo.Address) *User {
	return &User{
		ProfileAvatar: avatar,
		FirstName:     firstName,
		LastName:      lastName,
		PhoneNumber:   phone,
		Address:       address,
		UserType:      0,
	}
}

func (e *User) AsignAuthKey(auth *Auth) {
	e.AuthKeyId = auth.ID
}
