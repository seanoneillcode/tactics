package elem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Sprite struct {
	image *ebiten.Image
	x     float64
	y     float64
}

func NewSprite(imageFileName string, x float64, y float64) *Sprite {
	return &Sprite{
		image: common.LoadImage(imageFileName),
		x:     x,
		y:     y,
	}
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.x, s.y)
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(s.image, op)
}

func (s *Sprite) SetPos(pos Pos) {
	s.x = float64(pos.X)
	s.y = float64(pos.Y)
}
