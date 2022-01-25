package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Camera struct {
	pos          *common.VectorF
	buffer       *ebiten.Image
	followPlayer bool
}

func NewCamera() *Camera {
	return &Camera{
		pos:          &common.VectorF{},
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
}

func (c *Camera) DrawBuffer(screen *ebiten.Image) {
	screen.DrawImage(c.buffer, &ebiten.DrawImageOptions{
		GeoM: c.worldMatrix(),
	})
}

func (c *Camera) worldMatrix() ebiten.GeoM {
	m := ebiten.GeoM{}
	m.Translate(-c.pos.X, -c.pos.Y)
	return m
}

func (c *Camera) GetBuffer() *ebiten.Image {
	return c.buffer
}
