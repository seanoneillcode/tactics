package elem

import (
	"image/color"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var StandardFont font.Face

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

	StandardFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
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
		rect := text.BoundString(StandardFont, line)
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
	color color.RGBA
}

func NewText(x int, y int, value string) *Text {
	return &Text{
		value: value,
		x:     x,
		y:     y,
		color: DefaultTextColor,
	}
}

func (t *Text) SetValue(value string) {
	t.value = value
}

func (t *Text) SetColor(color color.RGBA) {
	t.color = color
}

func (t *Text) SetPosition(pos Pos) {
	t.x = pos.X
	t.y = pos.Y
}

func (t *Text) GetValue() string {
	return t.value
}

var DefaultTextColor = color.RGBA{
	R: 223,
	G: 246,
	B: 245,
	A: 255,
}

var GreyTextColor = color.RGBA{
	R: 160,
	G: 147,
	B: 142,
	A: 255,
}

var ActiveColor = color.RGBA{
	R: 200,
	G: 100,
	B: 150,
	A: 255,
}

func (t *Text) Draw(screen *ebiten.Image) {
	text.Draw(screen, t.value, StandardFont, (t.x)*common.Scale, (t.y+8)*common.Scale, t.color)
}
