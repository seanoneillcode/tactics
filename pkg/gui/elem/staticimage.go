package elem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type StaticImage struct {
	image *ebiten.Image
	x     float64
	y     float64
}

func NewStaticImage(imageFileName string, x float64, y float64) *StaticImage {
	return &StaticImage{
		image: common.LoadImage(imageFileName),
		x:     x,
		y:     y,
	}
}

func (s *StaticImage) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.x, s.y)
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(s.image, op)
}
