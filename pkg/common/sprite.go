package common

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	X      float64
	Y      float64
	Image  *ebiten.Image
	Frame  int
	IsFlip bool
}

func NewSprite(imageFileName string) *Sprite {
	return &Sprite{
		Image: LoadImage(imageFileName),
	}
}

func (s *Sprite) Draw(camera Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.X, s.Y)
	op.GeoM.Scale(Scale, Scale)
	camera.DrawImage(s.Image.SubImage(image.Rect(s.Frame*TileSize, 0, (s.Frame+1)*TileSize, TileSize)).(*ebiten.Image), op)
}

func (s *Sprite) SetPosition(x float64, y float64) {
	s.X = x
	s.Y = y
}

func (s *Sprite) SetFrame(frame int) {
	s.Frame = frame
}
