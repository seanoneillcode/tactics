package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"io/ioutil"
	"log"
)

var standardFont font.Face

func init() {
	// credit to pentacom for the font
	// http://www.pentacom.jp/pentacom/bitfontmaker2/gallery/?id=381
	b, err := ioutil.ReadFile("res/HelvetiPixel.ttf")
	if err != nil {
		log.Fatal(err)
	}
	tt, err := opentype.Parse(b)
	if err != nil {
		log.Fatal(err)
	}

	standardFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    64,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Text struct {
	value string
	x     int
	y     int
}

func NewText(value string, x int, y int) *Text {
	return &Text{
		value: value,
		x:     x,
		y:     y,
	}
}

func (t *Text) Draw(screen *ebiten.Image) {
	// add offset to y to account for text height
	text.Draw(screen, t.value, standardFont, (t.x+8)*common.Scale, (t.y)*common.Scale, color.White)
}
