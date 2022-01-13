package core

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"io/ioutil"
	"log"
	"strings"
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
	text.Draw(screen, t.value, standardFont, t.x*common.Scale, (t.y+8)*common.Scale, color.White)
}

type TextBox struct {
	text   *Text
	x      int
	y      int
	width  int
	height int
	image  *ebiten.Image
}

func NewTextBox(value string, x int, y int, width int, height int) *TextBox {
	image := ebiten.NewImage(width, height)
	image.Fill(color.White)
	return &TextBox{
		text:   NewText(getFormattedValue(value), x+8, y+8), // introduce a small margin
		x:      x,
		y:      y,
		width:  width,
		height: height,
		image:  image,
	}
}

func (tb *TextBox) SetTextValue(value string) {
	tb.text.value = getFormattedValue(value)
}

func (tb *TextBox) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(tb.x), float64(tb.y))
	op.GeoM.Scale(common.Scale, common.Scale)
	op.ColorM.Scale(0, 0.1, 0.5, 1)
	screen.DrawImage(tb.image, op)
	tb.text.Draw(screen)
}

func getFormattedValue(value string) string {
	// 40 chars per line
	var formatted string
	var line []string
	words := strings.Split(value, " ")

	count := 0
	for _, word := range words {
		// if adding a new word overruns, start a new line
		if count+len(word) > 40 {
			if formatted == "" {
				formatted = strings.Join(line, " ")
			} else {
				formatted = fmt.Sprintf("%s\n%s", formatted, strings.Join(line, " "))
			}
			line = []string{}
			count = 0
		}
		count = count + len(word) + 1
		line = append(line, word)

	}
	if len(line) > 0 {
		formatted = fmt.Sprintf("%s\n%s", formatted, strings.Join(line, " "))
	}
	return formatted
}
