package daos

type (
	ProfileDao struct {
		Email         string `gorm:"column:email"`
		UserName      string `gorm:"column:user_name"`
		ProfileAvatar string `gorm:"column:profile_avatar"`
	}
)
