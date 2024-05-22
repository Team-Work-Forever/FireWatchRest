package entities

import "github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"

type BurnRequestState struct {
	Entity
	AuthId      string `gorm:"type:uuid;primaryKey;column:auth_key_id"`
	BurnId      string `gorm:"type:uuid;primaryKey;column:burn_id"`
	State       uint16 `gorm:"column:state"`
	Observation string `gorm:"column:observation"`
}

func NewBurnRequestState(authId string, burnId string, state vo.BurnRequestStates, observation string) *BurnRequestState {
	return &BurnRequestState{
		AuthId:      authId,
		BurnId:      burnId,
		State:       uint16(state),
		Observation: observation,
	}
}

func (brs *BurnRequestState) TableName() string {
	return "burn_requests_states"
}

func (brs *BurnRequestState) GetState() string {
	state, _ := vo.GetBurnRequestState(brs.State)
	return state
}
