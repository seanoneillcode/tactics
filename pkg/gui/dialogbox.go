package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/core"
)

type DialogueBox struct {
	isActive bool
	textBox  *TextBox
	nameBox  *TextBox
	lastLine string
}

func NewDialogueBox() *DialogueBox {
	return &DialogueBox{
		textBox: NewTextBox("", 0, 160, 256, 80),
		//nameBox: NewTextBox("", 0, 136, 64, 23),
	}
}

func (b *DialogueBox) Draw(screen *ebiten.Image) {
	if b.isActive {
		//b.nameBox.Draw(screen)
		b.textBox.Draw(screen)
	}
}

func (b *DialogueBox) Update(delta int64, state *core.State) {
	ad := state.Player.ActiveDialog

	if ad == nil {
		b.isActive = false
		b.lastLine = ""
		b.textBox.SetTextValue("")
		//b.nameBox.SetTextValue("")
	} else {
		b.isActive = true
		_, text := ad.GetCurrentLine()
		if text != b.lastLine {
			b.lastLine = text
			b.textBox.SetTextValue(text)
			//b.nameBox.SetTextValue(name)
			//if name != "" {
			//	b.nameBox.SetActive(true)
			//} else {
			//	b.nameBox.SetActive(false)
			//}
		}

	}
}
