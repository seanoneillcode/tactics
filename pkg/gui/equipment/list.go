package equipment

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

type List struct {
	// assets
	pos      *elem.Pos
	itemList []*itemEntry
	bg       *elem.StaticImage

	// state
	currentIteration int
	index            int
}

func NewList(pos elem.Pos) *List {
	return &List{
		pos: &elem.Pos{X: pos.X, Y: pos.Y},
		bg:  elem.NewStaticImage("uis/equipment/list-bg.png", float64(pos.X), float64(pos.Y)),
	}
}

func (r *List) updateList(teamState *core.TeamState, slot string) {
	r.itemList = r.createList(teamState, slot)
}

func (r *List) createList(teamState *core.TeamState, slot string) []*itemEntry {
	itemNames := teamState.GetItemList()
	itemMap := teamState.ItemHolders
	x := r.pos.X + 4
	y := r.pos.Y + 4
	var invItems []*itemEntry
	var offset = 0
	for _, name := range itemNames {
		teamItem := itemMap[name]
		if teamItem.Item.CanConsume {
			continue
		}
		if teamItem.Item.EquipSlot != slot {
			continue
		}
		quantity := fmt.Sprintf("%v", teamItem.Amount)
		costWidth := text.BoundString(elem.StandardFont, quantity).Size().X / common.ScaleF
		entry := &itemEntry{
			itemRef:  teamItem.Item,
			name:     elem.NewText(x, y+offset, teamItem.Item.Name),
			quantity: elem.NewText(x+96+32+4+4-costWidth, y+offset, quantity),
		}
		invItems = append(invItems, entry)
		offset = offset + 16

	}
	return invItems
}

func (r *List) Draw(screen *ebiten.Image) {
	r.bg.Draw(screen)
	for _, item := range r.itemList {
		item.Draw(screen)
	}
}

func (r *List) Update(teamState *core.TeamState, slot string) {
	if r.currentIteration != teamState.Iteration {
		r.currentIteration = teamState.Iteration
		r.updateList(teamState, slot)
	}
}

func (r *List) handleInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		r.index = r.index - 1
		if r.index < 0 {
			r.index = 0
		}
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		r.index = r.index + 1
		if r.index == len(r.itemList) {
			r.index = r.index - 1
		}
		return
	}
}

type itemEntry struct {
	itemRef  *core.Item
	name     *elem.Text
	quantity *elem.Text
}

func (r *itemEntry) Draw(screen *ebiten.Image) {
	r.name.Draw(screen)
	r.quantity.Draw(screen)
}
