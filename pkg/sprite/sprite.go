package sprite

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"image"
	"io/ioutil"
	"log"
)

type Sprite struct {
	x     float64
	y     float64
	image *ebiten.Image
}

func NewSprite() *Sprite {
	b, err := ioutil.ReadFile("res/sprite.png")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	return &Sprite{
		image: ebiten.NewImageFromImage(img),
	}
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.x, s.y)
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(s.image, op)
}
