package entities

import "github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"

type BurnRequest struct {
	Entity
	AuthId         string           `gorm:"type:uuid;primaryKey;column:auth_key_id"`
	BurnId         string           `gorm:"type:uuid;primaryKey;column:burn_id"`
	InitialPropose string           `gorm:"column:initial_propose"`
	Accepted       bool             `gorm:"column:accepted"`
	State          BurnRequestState `gorm:"foreignKey:AuthId,BurnId;references:AuthId,BurnId"`
	Burn           Burn             `gorm:"foreignKey:BurnId;references:ID"`
}

func NewBurnRequest(authId string, burnId string, initialPropose string) *BurnRequest {
	return &BurnRequest{
		AuthId:         authId,
		BurnId:         burnId,
		InitialPropose: initialPropose,
	}
}

func (br *BurnRequest) SetState(state vo.BurnRequestStates, obs string) *BurnRequestState {
	br.State = *NewBurnRequestState(
		br.AuthId,
		br.BurnId,
		state,
		obs,
	)

	return &br.State
}

func (br *BurnRequest) TableName() string {
	return "burn_requests"
}
