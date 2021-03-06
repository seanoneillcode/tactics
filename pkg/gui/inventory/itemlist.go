package inventory

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/explore"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

type InvItemList struct {
	pos              *elem.Pos
	itemList         []*itemEntry
	currentIteration int
	currentIndex     int
}

func NewInvItemList() *InvItemList {
	return &InvItemList{
		pos: &elem.Pos{X: 8, Y: 8},
	}
}

func (i *InvItemList) CurrentItem() *explore.Item {
	if len(i.itemList) > 0 {
		return i.itemList[i.currentIndex].itemRef
	}
	return nil
}

func (i *InvItemList) createItemList(teamState *explore.TeamState) []*itemEntry {
	itemNames := teamState.GetItemList()
	itemMap := teamState.ItemHolders
	x := i.pos.X + offsetX
	y := i.pos.Y + offsetY
	var invItems []*itemEntry
	var offset = 0
	for _, name := range itemNames {
		teamItem := itemMap[name]
		if !teamItem.Item.CanConsume {
			continue
		}
		quantity := fmt.Sprintf("%v", teamItem.Amount)
		costWidth := text.BoundString(elem.StandardFont, quantity).Size().X / common.ScaleF
		entry := &itemEntry{
			itemRef:  teamItem.Item,
			name:     elem.NewText(x, y+offset, teamItem.Item.Name),
			quantity: elem.NewText(x+96+32+offsetX+offsetX-costWidth, y+offset, quantity),
		}
		invItems = append(invItems, entry)
		offset = offset + 16

	}
	return invItems
}

func (i *InvItemList) Draw(screen *ebiten.Image) {
	for _, item := range i.itemList {
		item.Draw(screen)
	}
}

func (i *InvItemList) Update(delta int64, teamState *explore.TeamState, currentIndex int) {
	i.currentIndex = currentIndex
	if i.currentIteration != teamState.Iteration {
		i.currentIteration = teamState.Iteration
		i.itemList = i.createItemList(teamState)
	}
}

type itemEntry struct {
	itemRef  *explore.Item
	name     *elem.Text
	quantity *elem.Text
}

func (l *itemEntry) Draw(screen *ebiten.Image) {
	l.name.Draw(screen)
	l.quantity.Draw(screen)
}
