package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
)

type DialogueBox struct {
	isActive         bool
	textBox          *TextBox
	nameBox          *TextBox
	currentText      string
	currentName      string
	x                int
	y                int
	width            int
	height           int
	currentDirection *common.Direction
}

func NewDialogueBox() *DialogueBox {
	return &DialogueBox{
		textBox: NewTextBox(-10, -10, 1, 1),
	}
}

func (b *DialogueBox) Draw(screen *ebiten.Image) {
	if b.isActive {
		b.textBox.Draw(screen)
	}
}

func (b *DialogueBox) Update(delta int64, state *core.State) {
	ad := state.ActiveDialog

	if ad != nil {
		if !b.isActive {
			// switch
			b.currentDirection = state.Player.Character.Direction
		}

		b.isActive = true

		line := ad.GetCurrentLine()
		if line.Text != b.currentText {
			b.currentText = line.Text
			if b.currentName != line.Name || line.Name == "" {
				b.width, b.height = GetMaxWidthHeight(ad.GetNextLinesForName())
				if b.currentName != "" {
					b.currentDirection = invertDirection(b.currentDirection)
				}
			}
			b.currentName = line.Name
			offset := b.getOffset()
			b.textBox = NewTextBox(b.x+offset.X, b.y+offset.Y, b.width+1, b.height+1)
			b.textBox.SetTextValue(line.FullText())
		}
	}

	if ad == nil && b.isActive {
		b.isActive = false
		b.textBox.SetTextValue("")
		b.currentName = ""
		b.currentText = ""
	}

}

func invertDirection(direction *common.Direction) *common.Direction {
	newDir := &common.Direction{
		X: direction.X * -1,
		Y: direction.Y * -1,
	}
	if newDir.X == 0 {
		newDir.X = -1
	}
	if newDir.Y == 0 {
		newDir.Y = -1
	}
	return newDir
}

func (b *DialogueBox) getOffset() *Point {
	offset := &Point{
		X: 8,
		Y: 80 - b.height,
	}
	if b.currentDirection.Y == 1 {
		offset.Y = 150
	}
	if b.currentDirection.X == 1 {
		offset.X = 240 - 8 - b.width
	}
	return offset
}
