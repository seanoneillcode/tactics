package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/core"
)

type DialogueBox struct {
	IsActive bool
	textBox  *TextBox
}

func NewDialogueBox() *DialogueBox {
	return &DialogueBox{
		textBox: NewTextBox("", 0, 160, 256, 80),
	}
}

func (g *DialogueBox) Draw(screen *ebiten.Image) {
	if g.IsActive {
		g.textBox.Draw(screen)
	}
}

func (g *DialogueBox) Update(delta int64, state *core.State) {
	ad := state.Player.ActiveDialog
	if ad == nil {
		g.IsActive = false
		g.textBox.SetTextValue("")
	} else {
		g.IsActive = true
		_, text := ad.GetCurrentLine()
		g.textBox.SetTextValue(text)
	}
}
