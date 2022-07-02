package action

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/fight"
)

type Ui struct {
	// assets
	list *List

	// state
	IsActive      bool
	isLoaded      bool
	selectedIndex int
}

func NewUI() *Ui {
	i := &Ui{
		list: NewList([]string{
			"move",
			"skill",
			"next",
		}),
	}
	return i
}

func (r *Ui) Draw(screen *ebiten.Image) {
	if !r.IsActive {
		return
	}
	r.list.Draw(screen)
}

func (r *Ui) Update(delta int64, state *fight.State) {
	if state.PlayerController.CurrentTurnPhase != fight.SelectActionTurnPhase {
		r.IsActive = false
		return
	}

	r.IsActive = true

	r.list.Update(state)

}
