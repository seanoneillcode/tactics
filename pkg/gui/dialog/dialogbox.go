package dialog

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

type DialogueBox struct {
	textBox          *elem.TextBox
	nameBox          *elem.TextBox
	currentText      string
	currentName      string
	x                int
	y                int
	width            int
	height           int
	currentDirection *common.Direction
	isLoaded         bool
}

func NewDialogueBox() *DialogueBox {
	return &DialogueBox{
		textBox: elem.NewTextBox(-10, -10, 1, 1),
	}
}

func (r *DialogueBox) Draw(screen *ebiten.Image) {
	r.textBox.Draw(screen)
}

func (r *DialogueBox) Update(delta int64, d *dialogState, state *core.State) {
	name, text := d.GetCurrentLine()
	if text != r.currentText || name != r.currentName {
		r.currentText = text
		if r.currentName != name || name == "" {
			r.width, r.height = elem.GetMaxWidthHeight(d.GetNextLinesForName())
			offset := r.getOffset(name, state)
			r.textBox = elem.NewTextBox(r.x+offset.X, r.y+offset.Y, r.width+1, r.height+1)
		}
		r.currentName = name
		r.textBox.SetTextValue(text)
	}
}

func (r *DialogueBox) Reset() {
	r.textBox.SetTextValue("")
	r.currentName = ""
	r.currentText = ""
}

func (r *DialogueBox) getOffset(name string, state *core.State) *elem.Pos {
	offset := &elem.Pos{
		X: 60,
		Y: 60 - r.height,
	}
	if name == "Player" {
		offset.X = 140
		offset.Y = common.HalfScreenHeight + r.height
	}
	// info box
	if name == "" {
		offset.X = 120
		offset.Y = common.HalfScreenHeight + r.height
	}
	// invert the directions
	if state.Player.Character.Direction.Y == 1 {
		if name == "Player" {
			offset.X = 60
			offset.Y = 60 - r.height
		} else {
			offset.X = 140
			offset.Y = common.HalfScreenHeight + r.height
		}
	}
	return offset
}
