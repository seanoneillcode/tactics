package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Camera struct {
	pos          *common.Vector
	buffer       *ebiten.Image
	followPlayer bool
	fadeTimer    int64
	isFading     bool
}

func NewCamera() *Camera {
	return &Camera{
		pos:          &common.Vector{},
		followPlayer: true,
		buffer:       ebiten.NewImage(common.ScreenWidth*common.Scale, common.ScreenHeight*common.Scale),
	}
}

func (c *Camera) Update(delta int64, state *State) {
	c.buffer.Clear()
	if c.followPlayer {
		screenOffset := common.VectorFromInt(common.HalfScreenWidth, common.HalfScreenHeight)
		tileOffset := common.VectorFromInt(common.HalfTileSize, common.HalfTileSize)
		c.pos = (state.Player.Character.pos.Sub(screenOffset).Add(tileOffset)).Mul(common.ScaleF)
	}
	if c.isFading {
		c.fadeTimer = c.fadeTimer - delta
		if c.fadeTimer < 0 {
			c.isFading = false
			state.Player.SetSleep(false)
		}
	} else {
		if state.Player.isSleeping {
			c.isFading = true
			c.fadeTimer = fadeTime
		}
	}
}

const fadeTime = 1000
const fadeTimeHalf = fadeTime / 2

func (c *Camera) DrawBuffer(screen *ebiten.Image) {
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM = c.worldMatrix()
	if c.fadeTimer > 0 {
		if c.fadeTimer >= fadeTimeHalf {
			var level = float64(c.fadeTimer-fadeTimeHalf) / float64(fadeTimeHalf)
			ops.ColorM.Scale(1, 1, 1, level)
		} else {
			ops.ColorM.Scale(1, 1, 1, 0)
		}
	}
	screen.DrawImage(c.buffer, ops)
}

func (c *Camera) worldMatrix() ebiten.GeoM {
	m := ebiten.GeoM{}
	m.Translate(-c.pos.X, -c.pos.Y)
	return m
}

func (c *Camera) GetBuffer() *ebiten.Image {
	return c.buffer
}
