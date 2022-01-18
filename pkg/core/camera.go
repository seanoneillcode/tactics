package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Camera struct {
	x            float64
	y            float64
	buffer       *ebiten.Image
	followPlayer bool
}

func NewCamera() *Camera {
	return &Camera{
		x:            0,
		y:            0,
		followPlayer: true,
		buffer:       ebiten.NewImage(common.ScreenWidth*common.Scale, common.ScreenHeight*common.Scale),
	}
}

func (c *Camera) Update(delta int64, state *State) {
	c.buffer.Clear()
	if c.followPlayer {
		c.x = (state.Player.character.x - common.HalfScreenWidth + common.HalfTileSize) * common.Scale
		c.y = (state.Player.character.y - common.HalfScreenHeight + common.HalfTileSize) * common.Scale
	}
}

func (c *Camera) DrawBuffer(screen *ebiten.Image) {
	screen.DrawImage(c.buffer, &ebiten.DrawImageOptions{
		GeoM: c.worldMatrix(),
	})
}

func (c *Camera) worldMatrix() ebiten.GeoM {
	m := ebiten.GeoM{}
	m.Translate(-c.x, -c.y)
	return m
}

func (c *Camera) GetBuffer() *ebiten.Image {
	return c.buffer
}
