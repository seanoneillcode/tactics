package explore

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Camera struct {
	pos          *common.Position
	buffer       *ebiten.Image
	followPlayer bool
	fadeLevel    float64
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
	c.fadeLevel = state.Map.fader.fadeLevel
}

func (c *Camera) DrawBuffer(screen *ebiten.Image) {
	ops := &ebiten.DrawImageOptions{}

	ops.ColorM.Scale(1, 1, 1, c.fadeLevel)
	screen.DrawImage(c.buffer, ops)
}

func (c *Camera) DrawImage(img *ebiten.Image, options *ebiten.DrawImageOptions) {
	options.GeoM.Translate(-c.pos.X, -c.pos.Y)
	c.buffer.DrawImage(img, options)
}
