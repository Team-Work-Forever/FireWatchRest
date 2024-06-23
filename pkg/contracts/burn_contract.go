package contracts

import (
	"time"

	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pagination"
)

type (
	CreateBurnRequest struct {
		UserId          string  `swaggerignore:"true"`
		Title           string  `form:"title" binding:"required"`
		Type            string  `form:"type" binding:"required"`
		HasBackUpTeam   bool    `form:"has_backup_team" binding:"required"`
		Reason          string  `form:"reason" binding:"required"`
		InitDate        string  `form:"init_date" binding:"required"`
		Lon             float64 `form:"lon" binding:"required"`
		Lat             float64 `form:"lat" binding:"required"`
		InitialProprose string  `form:"initial_propose" binding:"required"`
		Ignore          bool    `form:"ignore"`
	}

	UpdateBurnRequest struct {
		UserId        string `swaggerignore:"true"`
		BurnId        string `swaggerignore:"true"`
		Title         string `form:"title" binding:"required"`
		Type          string `form:"type" binding:"required"`
		HasBackUpTeam string `form:"has_backup_team" binding:"required"`
		Reason        string `form:"reason" binding:"required"`
		InitDate      string `form:"init_date" binding:"required"`
		Lat           string `form:"lat" binding:"required"`
		Lon           string `form:"lon" binding:"required"`
	}

	TerminateBurnRequest struct {
		UserId string
		BurnId string
	}

	StartBurnRequest struct {
		UserId string
		BurnId string
	}

	DeleteBurnRequest struct {
		UserId string
		BurnId string
	}

	GetBurnRequest struct {
		AuthId string
		BurnId string
	}

	GetAllBurnsRequest struct {
		AutarchyId string
		AuthId     string
		Search     string
		Sort       string
		State      string
		StartDate  string
		EndDate    string
		Pagination *pagination.Pagination
	}

	BurnResponse struct {
		Id          string                `json:"id"`
		Title       string                `json:"title"`
		Author      PublicProfileResponse `json:"author"`
		HasAidTeam  bool                  `json:"has_aid_team"`
		Reason      string                `json:"reason"`
		Type        string                `json:"type"`
		Address     AddressResponse       `json:"address"`
		BeginAt     time.Time             `json:"begin_at"`
		CompletedAt *time.Time            `json:"completed_at"`
		Picture     string                `json:"map_picture"`
		State       string                `json:"state"`
	}

	BurnWithAuthorResponse struct {
		BurnResponse
		Author PublicProfileResponse `json:"author"`
	}

	BurnActionResponse struct {
		BurnId string `json:"burn_id"`
		State  string `json:"state"`
	}
)
