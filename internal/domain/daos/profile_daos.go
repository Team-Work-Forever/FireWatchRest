package daos

import "github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"

type (
	ProfileDao struct {
		Email         string   `gorm:"column:email"`
		NIF           string   `gorm:"column:nif"`
		UserName      string   `gorm:"column:user_name"`
		ProfileAvatar string   `gorm:"column:profile_avatar"`
		Phone         vo.Phone `gorm:"embedded"`
	}
)
