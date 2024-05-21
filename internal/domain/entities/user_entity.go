package entities

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
)

type User struct {
	EntityBase
	AuthKeyId     string     `gorm:"column:auth_key_id"`
	ProfileAvatar string     `gorm:"column:profile_avatar"`
	UserName      string     `gorm:"column:user_name"`
	FirstName     string     `gorm:"column:first_name"`
	LastName      string     `gorm:"column:last_name"`
	PhoneNumber   vo.Phone   `gorm:"embedded"`
	Address       vo.Address `gorm:"embedded"`
}

func NewUser(avatar string, userName string, firstName string, lastName string, phone vo.Phone, address vo.Address) *User {
	return &User{
		ProfileAvatar: avatar,
		UserName:      userName,
		FirstName:     firstName,
		LastName:      lastName,
		PhoneNumber:   phone,
		Address:       address,
	}
}

func (e *User) AsignAuthKey(auth *Auth) {
	e.AuthKeyId = auth.ID
}
