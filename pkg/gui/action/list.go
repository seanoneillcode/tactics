package action

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/fight"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

type List struct {
	// assets
	pos elem.Pos

	// state
	items    []*Entry
	index    int
	isActive bool
}

func NewList(actions []string) *List {
	l := &List{
		items: []*Entry{},
	}
	for _, action := range actions {
		l.items = append(l.items, NewEntry(action))
	}
	return l
}

func (r *List) Draw(screen *ebiten.Image) {
	if !r.isActive {
		return
	}
	for _, item := range r.items {
		item.Draw(screen)
	}
}

func (r *List) Update(state *fight.State) {
	r.isActive = state.ActiveTeam == state.PlayerTeam
	r.pos = elem.Pos{
		X: common.HalfScreenWidth - common.HalfTileSize,
		Y: common.HalfScreenHeight + common.TileSize,
	}
	pos := elem.Pos{
		X: r.pos.X,
		Y: r.pos.Y,
	}
	highlightedIndex := state.PlayerController.SelectedActionIndex
	for index, item := range r.items {
		item.Update(pos, index == highlightedIndex)
		pos.Y = pos.Y + 16
	}
}
