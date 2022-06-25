package fight

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Camera struct {
	pos    *common.Position
	buffer *ebiten.Image
}

func NewCamera() *Camera {
	return &Camera{
		pos:    &common.Position{},
		buffer: ebiten.NewImage(common.ScreenWidth*common.Scale, common.ScreenHeight*common.Scale),
	}
}

func (c *Camera) Update(delta int64, state *State) {
	c.buffer.Clear()
	//if c.followPlayer {
	//	screenOffset := common.PositionFromInt(common.HalfScreenWidth, common.HalfScreenHeight)
	//	tileOffset := common.PositionFromInt(common.HalfTileSize, common.HalfTileSize)
	//	c.pos = (state.Player.Character.pos.Sub(screenOffset).Add(tileOffset)).Mul(common.ScaleF)
	//}
}

func (c *Camera) DrawBuffer(screen *ebiten.Image) {
	ops := &ebiten.DrawImageOptions{}
	screen.DrawImage(c.buffer, ops)
}

func (c *Camera) DrawImage(img *ebiten.Image, options *ebiten.DrawImageOptions) {
	options.GeoM.Translate(-c.pos.X, -c.pos.Y)
	c.buffer.DrawImage(img, options)
}
