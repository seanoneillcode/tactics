package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type DialogueBox struct {
	textBoxes []*TextBox
}

func NewDialogueBox() *DialogueBox {
	return &DialogueBox{}
}

func (g *DialogueBox) AddTextBox(value string) {
	tb := NewTextBox(value, 0, 120, 256, 120)
	g.textBoxes = append(g.textBoxes, tb)
}

func (g *DialogueBox) Draw(screen *ebiten.Image) {
	for _, tb := range g.textBoxes {
		tb.Draw(screen)
	}
}
