package dialog

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/explore"
	"github.com/seanoneillcode/go-tactics/pkg/input"
)

type ui struct {
	IsActive     bool
	isLoaded     bool
	dialogBox    *DialogueBox
	activeDialog *dialogState
}

func NewUi() *ui {
	return &ui{
		dialogBox: NewDialogueBox(),
	}
}

func (r *ui) Draw(screen *ebiten.Image) {
	if !r.IsActive {
		return
	}
	r.dialogBox.Draw(screen)
}

func (r *ui) Update(delta int64, state *explore.State) {
	if !state.DialogHandler.IsActive {
		r.IsActive = false
		r.isLoaded = false
		return
	}
	r.IsActive = true
	if !r.isLoaded {
		r.isLoaded = true
		r.activeDialog = NewDialogState(state.DialogHandler.ActiveDialog)
		return
	}
	r.handleInput(state)
	if r.activeDialog != nil {
		r.activeDialog.Update(delta)
	}
	r.dialogBox.Update(delta, r.activeDialog, state)
}

func (r *ui) handleInput(state *explore.State) {
	if input.IsEnterPressed() {
		if r.activeDialog.IsBuffering() {
			r.activeDialog.SkipBuffer()
		} else {
			if r.activeDialog.HasNextLine() {
				r.activeDialog.NextLine()
			} else {
				r.activeDialog.Reset()
				r.dialogBox.Reset()
				state.DialogHandler.CloseDialog()
			}
		}
	}
}
