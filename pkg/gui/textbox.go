package gui

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"image/color"
	"strings"
)

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
