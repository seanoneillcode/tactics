package equipment

import (
	"fmt"
	"github.com/seanoneillcode/go-tactics/pkg/input"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/explore"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

type List struct {
	// assets
	pos       *elem.Pos
	itemList  []*itemEntry
	bg        *elem.StaticImage
	highlight *elem.Sprite

	// state
	currentIteration int
	index            int
}

func NewList(pos elem.Pos) *List {
	return &List{
		pos:       &elem.Pos{X: pos.X, Y: pos.Y},
		bg:        elem.NewStaticImage("uis/equipment/list-bg.png", float64(pos.X), float64(pos.Y)),
		highlight: elem.NewSprite("uis/equipment/highlight.png", 0, 0),
	}
}

func (r *List) updateList(teamState *explore.TeamState, slot string) {
	r.itemList = r.createList(teamState, slot)
}

func (r *List) createList(teamState *explore.TeamState, slot string) []*itemEntry {
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
	r.index = 0
	return invItems
}

func (r *List) Draw(screen *ebiten.Image) {
	r.bg.Draw(screen)
	r.highlight.Draw(screen)
	for _, item := range r.itemList {
		item.Draw(screen)
	}
}

func (r *List) Update(teamState *explore.TeamState, slot string) {
	if r.currentIteration != teamState.Iteration {
		r.currentIteration = teamState.Iteration
		r.updateList(teamState, slot)
	}
	r.highlight.SetPos(&elem.Pos{
		X: r.pos.X,
		Y: r.pos.Y + (16 * r.index),
	})
}

func (r *List) handleInput() {
	if input.IsUpJustPressed() {
		r.index = r.index - 1
		if r.index < 0 {
			r.index = 0
		}
		return
	}
	if input.IsDownJustPressed() {
		r.index = r.index + 1
		if r.index == len(r.itemList) {
			r.index = r.index - 1
		}
		return
	}
}

func (r *List) currentItem() *itemEntry {
	return r.itemList[r.index]
}

type itemEntry struct {
	itemRef  *explore.Item
	name     *elem.Text
	quantity *elem.Text
}

func (r *itemEntry) Draw(screen *ebiten.Image) {
	r.name.Draw(screen)
	r.quantity.Draw(screen)
}
