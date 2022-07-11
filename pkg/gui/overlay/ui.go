package overlay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/fight"
)

type Ui struct {
	// assets
	list []*Card

	// state
	IsActive      bool
	isLoaded      bool
	selectedIndex int
}

func NewUI() *Ui {
	i := &Ui{
		list: []*Card{},
	}
	return i
}

func (r *Ui) Draw(screen *ebiten.Image) {
	if !r.IsActive {
		return
	}
	for _, card := range r.list {
		card.Draw(screen)
	}
}

func (r *Ui) Update(delta int64, state *fight.State) {
	if state.PlayerController.CurrentTurnPhase != fight.SelectActionTurnPhase {
		r.IsActive = false
		return
	}

	r.IsActive = true

	for _, card := range r.list {
		card.Update(state)
	}

}
