package dialog

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanoneillcode/go-tactics/pkg/core"
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

func (r *ui) Update(delta int64, state *core.State) {
	if !state.DialogHandler.IsActive {
		r.IsActive = false
		r.isLoaded = false
		return
	}
	r.IsActive = true
	if !r.isLoaded {
		r.isLoaded = true
		r.activeDialog = NewDialog(state.DialogHandler.ActiveDialog)
		return
	}
	r.handleInput(state)
	if r.activeDialog != nil {
		r.activeDialog.Update(delta, state)
	}
	r.dialogBox.Update(delta, r.activeDialog, state)
}

func (r *ui) handleInput(state *core.State) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
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
