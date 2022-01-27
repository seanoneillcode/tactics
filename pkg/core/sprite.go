package core

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"image"
	"io/ioutil"
	"log"
)

type Sprite struct {
	x      float64
	y      float64
	image  *ebiten.Image
	frame  int
	isFlip bool
}

func NewSprite(imageFileName string) *Sprite {
	return &Sprite{
		image: LoadImage(imageFileName),
	}
}

func LoadImage(imageFileName string) *ebiten.Image {
	b, err := ioutil.ReadFile("res/" + imageFileName)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.x, s.y)
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(s.image.SubImage(image.Rect(s.frame*common.TileSize, 0, (s.frame+1)*common.TileSize, common.TileSize)).(*ebiten.Image), op)
}

func (s *Sprite) SetPosition(x float64, y float64) {
	s.x = x
	s.y = y
}

func (s *Sprite) SetFrame(frame int) {
	s.frame = frame
}
