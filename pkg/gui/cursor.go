package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
)

type Cursor struct {
	pos         Pos
	image       *ebiten.Image
	show        bool
	offset      int
	offsetTimer int64
}

func NewCursor() *Cursor {
	return &Cursor{
		image: core.LoadImage("shop-cursor.png"),
		pos:   Pos{},
	}
}

func (c *Cursor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.pos.X+c.offset), float64(c.pos.Y))
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(c.image, op)
}

func (c *Cursor) Update(delta int64, pos *Pos) {
	c.pos.X = pos.X
	c.pos.Y = pos.Y
	c.offsetTimer = c.offsetTimer + delta
	if c.offsetTimer > 400 {
		c.offsetTimer = c.offsetTimer - 400
		if c.offset == 0 {
			c.offset = 2
		} else {
			c.offset = 0
		}
	}
}
