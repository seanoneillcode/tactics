package action

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

type List struct {
	// assets
	pos elem.Pos

	// state
	items []*Entry
	index int
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
	for _, item := range r.items {
		item.Draw(screen)
	}
}

func (r *List) Update() {
	r.pos = elem.Pos{
		X: common.HalfScreenWidth - common.HalfTileSize,
		Y: common.HalfScreenHeight + common.TileSize,
	}
	pos := elem.Pos{
		X: r.pos.X,
		Y: r.pos.Y,
	}
	for _, item := range r.items {
		item.Update(pos)
		pos.Y = pos.Y + 16
	}
}
