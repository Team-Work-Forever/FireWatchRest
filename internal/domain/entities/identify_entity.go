package entities

import "github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"

type UpdateIdentity interface {
	SetPhone(phone *vo.Phone)
	SetZipCode(zipCode *vo.ZipCode)
	SetStreet(street string)
	SetStreetNumber(number int)
	SetCity(city string)
	SetPicture(url string)
	GetAddress() vo.Address
}

type IdentityUser struct {
	EntityBase
	AuthKeyId   string     `gorm:"column:auth_key_id"`
	PhoneNumber vo.Phone   `gorm:"embedded"`
	Address     vo.Address `gorm:"embedded"`
	Picture     string     `gorm:"column:profile_avatar"`
}

func (i *IdentityUser) SetPhone(phone *vo.Phone) {
	i.PhoneNumber = *phone
}

func (i *IdentityUser) SetZipCode(zipCode *vo.ZipCode) {
	i.Address.ZipCode = zipCode.GetValue()
}

func (i *IdentityUser) SetStreet(street string) {
	i.Address.Street = street
}

func (i *IdentityUser) SetStreetNumber(number int) {
	i.Address.SetStreetNumber(number)
}

func (i *IdentityUser) SetCity(city string) {
	i.Address.City = city
}

func (i *IdentityUser) SetPicture(url string) {
	i.Picture = url
}

func (i *IdentityUser) GetAddress() vo.Address {
	return i.Address
}
