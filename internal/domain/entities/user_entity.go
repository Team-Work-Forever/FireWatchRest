package entities

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
)

type User struct {
	IdentityUser
	UserName  string `gorm:"column:user_name"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
}

func NewUser(
	avatar string,
	userName string,
	firstName string,
	lastName string,
	phone vo.Phone,
	address vo.Address,
) *User {
	return &User{
		IdentityUser: IdentityUser{
			Picture:     avatar,
			PhoneNumber: phone,
			Address:     address,
		},
		UserName:  userName,
		FirstName: firstName,
		LastName:  lastName,
	}
}

func (e *User) AsignAuthKey(auth *Auth) {
	e.AuthKeyId = auth.ID
}
