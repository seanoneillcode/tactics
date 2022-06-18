package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Camera struct {
	pos           *common.Position
	buffer        *ebiten.Image
	followPlayer  bool
	fadeTimer     int64
	isFading      bool
	isFadeOut     bool
	isFadeIn      bool
	fadeTimeTotal int64
}

func NewCamera() *Camera {
	return &Camera{
		pos:          &common.Position{},
		followPlayer: true,
		buffer:       ebiten.NewImage(common.ScreenWidth*common.Scale, common.ScreenHeight*common.Scale),
	}
}

func (c *Camera) Update(delta int64, state *State) {
	c.buffer.Clear()
	if c.followPlayer {
		screenOffset := common.PositionFromInt(common.HalfScreenWidth, common.HalfScreenHeight)
		tileOffset := common.PositionFromInt(common.HalfTileSize, common.HalfTileSize)
		c.pos = (state.Player.Character.pos.Sub(screenOffset).Add(tileOffset)).Mul(common.ScaleF)
	}
	if c.isFadeOut {
		c.fadeTimer = c.fadeTimer - delta
		if c.fadeTimer < 0 {
			c.isFadeOut = false
			c.isFadeIn = true
			c.fadeTimer = state.Map.fadeTotalTime
		}
	}
	if c.isFadeIn {
		c.fadeTimer = c.fadeTimer - delta
		if c.fadeTimer < 0 {
			c.isFadeIn = false
		}
	}
	if c.fadeTimer <= 0 {
		if state.Map.fadeOut {
			c.isFadeOut = true
			c.fadeTimer = state.Map.fadeTotalTime
			c.fadeTimeTotal = state.Map.fadeTotalTime
		}
	}
}

func (c *Camera) DrawBuffer(screen *ebiten.Image) {
	ops := &ebiten.DrawImageOptions{}
	if c.fadeTimer > 0 {
		if c.isFadeOut {
			var level = float64(c.fadeTimer) / float64(c.fadeTimeTotal)
			ops.ColorM.Scale(1, 1, 1, level)
		}
		if c.isFadeIn {
			var level = float64(c.fadeTimeTotal-c.fadeTimer) / float64(c.fadeTimeTotal)
			ops.ColorM.Scale(1, 1, 1, level)
		}
	}
	screen.DrawImage(c.buffer, ops)
}

func (c *Camera) DrawImage(img *ebiten.Image, options *ebiten.DrawImageOptions) {
	options.GeoM.Translate(-c.pos.X, -c.pos.Y)
	c.buffer.DrawImage(img, options)
}
