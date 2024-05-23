package contracts

import (
	"time"
)

type (
	CreateBurnRequest struct {
		UserId          string  `swaggerignore:"true"`
		Title           string  `form:"title" binding:"required"`
		Type            string  `form:"type" binding:"required"`
		HasBackUpTeam   bool    `form:"has_backup_team" binding:"required"`
		Reason          string  `form:"reason" binding:"required"`
		InitDate        string  `form:"init_date" binding:"required"`
		Lon             float32 `form:"lon" binding:"required"`
		Lat             float32 `form:"lat" binding:"required"`
		InitialProprose string  `form:"initial_propose" binding:"required"`
	}

	GetBurnRequest struct {
		AuthId string
		BurnId string
	}

	GetAllBurnsRequest struct {
		AuthId    string
		Search    string
		State     string
		StartDate string
		EndDate   string
		PageSize  uint64
		Page      uint64
	}

	BurnResponse struct {
		Id          string    `json:"id"`
		Title       string    `json:"title"`
		HasAidTeam  bool      `json:"has_aid_team"`
		Reason      string    `json:"reason"`
		Type        string    `json:"type"`
		BeginAt     time.Time `json:"begin_at"`
		CompletedAt time.Time `json:"completed_at"`
		Picture     string    `json:"map_picture"`
		State       string    `json:"state"`
	}

	BurnActionResponse struct {
		BurnId string `json:"burn_id"`
		State  string `json:"state"`
	}
)
