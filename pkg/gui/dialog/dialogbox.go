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
	line := d.GetCurrentLine()
	if line.Text != r.currentText {
		r.currentText = line.Text
		if r.currentName != line.Name || line.Name == "" {
			r.width, r.height = elem.GetMaxWidthHeight(d.GetNextLinesForName())
			offset := r.getOffset(line, state)
			r.textBox = elem.NewTextBox(r.x+offset.X, r.y+offset.Y, r.width+1, r.height+1)
		}
		r.currentName = line.Name
		r.textBox.SetTextValue(line.FullText())
	}
}

func (r *DialogueBox) Reset() {
	r.textBox.SetTextValue("")
	r.currentName = ""
	r.currentText = ""
}

func (r *DialogueBox) getOffset(line *lineState, state *core.State) *elem.Pos {
	offset := &elem.Pos{
		X: 60,
		Y: 60 - r.height,
	}
	if line.Name == "Player" {
		offset.X = 140
		offset.Y = 140 + r.height
	}
	// info box
	if line.Name == "" {
		offset.X = 120
		offset.Y = 120 + r.height
	}
	// invert the directions
	if state.Player.Character.Direction.Y == 1 {
		if line.Name == "Player" {
			offset.X = 60
			offset.Y = 60 - r.height
		} else {
			offset.X = 140
			offset.Y = 140 + r.height
		}
	}
	return offset
}
