package daos

import (
	"time"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
)

type (
	CreateBurnDao struct {
		AuthId         string
		Burn           *entities.Burn
		InitialPropose string
	}

	BurnDetailsView struct {
		Author      string    `gorm:"column:author"`
		Id          string    `gorm:"column:id"`
		Title       string    `gorm:"column:title"`
		MapPicture  string    `gorm:"column:map_picture"`
		HasAidTeam  bool      `gorm:"column:has_aid_team"`
		Lat         float32   `gorm:"column:lat"`
		Lon         float32   `gorm:"column:lon"`
		Reason      int       `gorm:"column:reason"`
		Type        int       `gorm:"column:type"`
		BeginAt     time.Time `gorm:"column:begin_at"`
		CompletedAt time.Time `gorm:"column:completed_at"`
		State       uint16    `gorm:"column:state"`
	}
)

func (bd *BurnDetailsView) TableName() string {
	return "burn_details_view"
}