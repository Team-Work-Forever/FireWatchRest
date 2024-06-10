package entities

import "github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"

type IdentityUser struct {
	EntityBase
	AuthKeyId   string     `gorm:"column:auth_key_id"`
	PhoneNumber vo.Phone   `gorm:"embedded"`
	Address     vo.Address `gorm:"embedded"`
	Picture     string     `gorm:"column:profile_avatar"`
}
