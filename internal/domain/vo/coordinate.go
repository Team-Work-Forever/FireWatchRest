package vo

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Coordinate struct {
	x float64 `gorm:"-"`
	y float64 `gorm:"-"`
}

func NewCoordinate(x float64, y float64) *Coordinate {
	return &Coordinate{
		x: x,
		y: y,
	}
}

func (c *Coordinate) GetX() float64 {
	return c.x
}

func (c *Coordinate) GetY() float64 {
	return c.y
}

func (c Coordinate) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_GeomFromText(?, 4326)",
		Vars: []interface{}{fmt.Sprintf("POINT(%.2f %.2f)", c.GetX(), c.GetY())},
	}
}

func (c *Coordinate) Scan(value interface{}) error {
	c.x = 0
	c.y = 0
	return nil
}
