package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/core"
)

type DialogueBox struct {
	isActive    bool
	textBox     *TextBox
	nameBox     *TextBox
	currentText string
	x           int
	y           int
	width       int
	height      int
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
	ad := state.Player.ActiveDialog

	if ad != nil {
		if !b.isActive {
			// switch
			b.width, b.height = GetMaxWidthHeight(ad.GetAllFormattedLines())
		}

		b.isActive = true
		text2 := ad.GetCurrentText()
		if text2 != b.currentText {
			offsetX, offsetY := 240-8-b.width, 100-b.height
			if ad.GetNameOrder() == 0 {
				offsetX = 8
				offsetY = 160
				if b.width < 120 {
					offsetX = offsetX + 120 - b.width
				}
			}
			b.textBox = NewTextBox(b.x+offsetX, b.y+offsetY, b.width+1, b.height+1)
			b.textBox.SetTextValue(text2)
		}
	}

	if ad == nil && b.isActive {
		b.isActive = false
		b.textBox.SetTextValue("")
	}

}
