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

func GetMaxWidthHeight(lines []string) (int, int) {
	maxHeight := 1
	maxWidth := 1
	for _, line := range lines {
		rect := text.BoundString(standardFont, line)
		height := rect.Max.Y - rect.Min.Y
		if height > maxHeight {
			maxHeight = height
		}
		width := rect.Max.X - rect.Min.X
		if width > maxWidth {
			maxWidth = width
		}
	}
	return maxWidth / 4, maxHeight / 4
}

type Text struct {
	value string
	x     int
	y     int
}

func NewText(x int, y int) *Text {
	return &Text{
		x: x,
		y: y,
	}
}

func (t *Text) SetValue(value string) {
	t.value = value
}

func (t *Text) Draw(screen *ebiten.Image) {
	text.Draw(screen, t.value, standardFont, (t.x)*common.Scale, (t.y+8)*common.Scale, color.White)
}
