package vo

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Coordinate struct {
	x float32 `gorm:"-"`
	y float32 `gorm:"-"`
}

func NewCoordinate(x float32, y float32) *Coordinate {
	return &Coordinate{
		x: x,
		y: y,
	}
}

func (c *Coordinate) GetX() float32 {
	return c.x
}

func (c *Coordinate) GetY() float32 {
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
