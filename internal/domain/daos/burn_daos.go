package daos

import (
	"time"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
)

type (
	CreateBurnDao struct {
		AuthId         string
		AutarchyId     string
		Burn           *entities.Burn
		InitialPropose string
		State          vo.BurnRequestStates
	}

	BurnDetailsView struct {
		entities.Entity
		Author      string    `gorm:"column:author"`
		Id          string    `gorm:"column:id"`
		Title       string    `gorm:"column:title"`
		MapPicture  string    `gorm:"column:map_picture"`
		HasAidTeam  bool      `gorm:"column:has_aid_team"`
		Lat         float64   `gorm:"column:lat"`
		Lon         float64   `gorm:"column:lon"`
		Street      string    `gorm:"column:address_street"`
		Number      int       `gorm:"column:address_number;type:int"`
		ZipCode     string    `gorm:"column:address_zip_code"`
		City        string    `gorm:"column:address_city"`
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
