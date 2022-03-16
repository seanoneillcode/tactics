package gui

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
)

type InvItemList struct {
	pos              *Pos
	itemList         []*itemEntry
	currentIteration int
}

func NewInvItemList() *InvItemList {
	return &InvItemList{
		pos: &Pos{32, 32},
	}
}

func (i *InvItemList) createItemList(inventory *core.Inventory) []*itemEntry {

	itemNames := inventory.TeamState.GetItemList()
	itemMap := inventory.TeamState.Items

	x := i.pos.X + offsetX
	y := i.pos.Y + offsetY
	var invItems []*itemEntry
	var offset = 0
	for _, name := range itemNames {
		teamItem := itemMap[name]
		quantity := fmt.Sprintf("%v", teamItem.Amount)
		costWidth := text.BoundString(standardFont, quantity).Size().X / common.ScaleF
		invItems = append(invItems, &itemEntry{
			itemRef: teamItem.Item,
			name: &Text{
				value: teamItem.Item.Name,
				x:     x,
				y:     y + offset,
				color: defaultTextColor,
			},
			quantity: &Text{
				value: quantity,
				x:     x + 96 + 32 + offsetX + offsetX - costWidth,
				y:     y + offset,
				color: defaultTextColor,
			},
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
	name     *Text
	quantity *Text
}

func (l *itemEntry) Draw(screen *ebiten.Image) {
	l.name.Draw(screen)
	l.quantity.Draw(screen)
}
