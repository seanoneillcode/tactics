package gui

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

type InvItemList struct {
	pos              *elem.Pos
	itemList         []*itemEntry
	currentIteration int
}

func NewInvItemList() *InvItemList {
	return &InvItemList{
		pos: &elem.Pos{X: 32, Y: 32},
	}
}

func (i *InvItemList) createItemList(inventory *core.Inventory) []*itemEntry {
	itemNames := inventory.TeamState.GetItemList()
	itemMap := inventory.TeamState.ItemHolders
	x := i.pos.X + offsetX
	y := i.pos.Y + offsetY
	var invItems []*itemEntry
	var offset = 0
	for _, name := range itemNames {
		teamItem := itemMap[name]
		quantity := fmt.Sprintf("%v", teamItem.Amount)
		costWidth := text.BoundString(elem.StandardFont, quantity).Size().X / common.ScaleF
		invItems = append(invItems, &itemEntry{
			itemRef:  teamItem.Item,
			name:     elem.NewText(x, y+offset, teamItem.Item.Name),
			quantity: elem.NewText(x+96+32+offsetX+offsetX-costWidth, y+offset, quantity),
		})
		offset = offset + 16

	}
	return invItems
}

func (i *InvItemList) Draw(screen *ebiten.Image) {
	for _, item := range i.itemList {
		item.Draw(screen)
	}
}

func (i *InvItemList) Update(delta int64, inventory *core.Inventory) {
	if i.currentIteration != inventory.TeamState.Iteration {
		i.currentIteration = inventory.TeamState.Iteration
		i.itemList = i.createItemList(inventory)
	}
}

type itemEntry struct {
	itemRef  *core.Item
	name     *elem.Text
	quantity *elem.Text
}

func (l *itemEntry) Draw(screen *ebiten.Image) {
	l.name.Draw(screen)
	l.quantity.Draw(screen)
}
