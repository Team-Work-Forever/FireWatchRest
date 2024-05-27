package daos

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
)

type (
	AutarchyDetailsView struct {
		entities.Entity
		Email          string     `gorm:"column:email"`
		Id             string     `gorm:"column:id"`
		Title          string     `gorm:"column:title"`
		AutarchyAvatar string     `gorm:"column:autarchy_avatar"`
		Lat            float64    `gorm:"column:lat"`
		Lon            float64    `gorm:"column:lon"`
		PhoneNumber    vo.Phone   `gorm:"embedded"`
		Address        vo.Address `gorm:"embedded"`
	}
)

func (adv *AutarchyDetailsView) TableName() string {
	return "autarchy_details_view"
}
