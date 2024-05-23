package entities

import (
	"errors"
	"time"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
)

type Burn struct {
	EntityBase
	Title       string        `gorm:"column:title"`
	HasAidTeam  bool          `gorm:"column:has_aid_team"`
	Reason      uint16        `gorm:"column:reason"`
	Type        uint16        `gorm:"column:type"`
	Coordinates vo.Coordinate `gorm:"column:geo_location;type:geometry"`
	BeginAt     time.Time     `gorm:"column:begin_at"`
	CompletedAt time.Time     `gorm:"column:completed_at"`
	Picture     string        `gorm:"column:map_picture"`
}

func NewBurn(
	title string,
	reason uint16,
	ttype uint16,
	coordinates vo.Coordinate,
	beginAt time.Time,
) (*Burn, error) {
	if title == "" {
		return nil, errors.New("title is not provided")
	}

	if beginAt.Before(time.Now()) {
		return nil, errors.New("provide an valid start date")
	}

	return &Burn{
		Title:       title,
		Reason:      reason,
		Type:        ttype,
		Coordinates: coordinates,
		BeginAt:     beginAt,
	}, nil
}

func (b *Burn) TableName() string {
	return "burn"
}
