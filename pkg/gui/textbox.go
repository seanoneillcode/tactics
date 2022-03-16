package gui

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

const marginX = 4
const marginY = 4

type TextBox struct {
	text     *elem.Text
	x        int
	y        int
	width    int
	height   int
	partial1 image.Image
	partial2 image.Image
	partial3 image.Image
	partial4 image.Image
}

func NewTextBox(x int, y int, width int, height int) *TextBox {
	borderImage1 := common.LoadImage("text-border-1.png")
	partial1 := borderImage1.SubImage(image.Rect(0, 0, width+marginX+marginX, height+marginY+marginY))

	borderImage2 := common.LoadImage("text-border-2.png")
	partial2 := borderImage2.SubImage(image.Rect(0, 0, width+marginX+marginX, height))

	borderImage3 := common.LoadImage("text-border-3.png")
	partial3 := borderImage3

	borderImage4 := common.LoadImage("text-border-4.png")
	partial4 := borderImage4.SubImage(image.Rect(0, 0, width, height+marginY))

	return &TextBox{
		text:     elem.NewText(x+marginX, y+marginY, ""),
		x:        x,
		y:        y,
		width:    width,
		height:   height,
		partial1: partial1,
		partial2: partial2,
		partial3: partial3,
		partial4: partial4,
	}
}

func (tb *TextBox) SetTextValue(value string) {
	tb.text.SetValue(value)
}

func (tb *TextBox) Draw(screen *ebiten.Image) {
	if tb.text.GetValue() == "" {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(tb.x), float64(tb.y))
	op.GeoM.Scale(common.Scale, common.Scale)
	temp := ebiten.NewImageFromImage(tb.partial1)

	screen.DrawImage(temp, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(tb.x), float64(tb.y+tb.height+marginY))
	op.GeoM.Scale(common.Scale, common.Scale)
	temp = ebiten.NewImageFromImage(tb.partial2)

	screen.DrawImage(temp, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(tb.x+tb.width+marginX), float64(tb.y+tb.height+marginY))
	op.GeoM.Scale(common.Scale, common.Scale)
	temp = ebiten.NewImageFromImage(tb.partial3)

	screen.DrawImage(temp, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(tb.x+tb.width+marginX), float64(tb.y))
	op.GeoM.Scale(common.Scale, common.Scale)
	temp = ebiten.NewImageFromImage(tb.partial4)

	screen.DrawImage(temp, op)

	tb.text.Draw(screen)
}

func (tb *TextBox) SetPosition(x int, y int) {
	tb.x = x
	tb.y = y
	tb.text.SetPosition(elem.Pos{X: x + marginX, Y: y + marginY})
}
