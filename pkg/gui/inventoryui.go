package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/core"
)

type InventoryUi struct {
	cursor           *Cursor
	actionBox        *ActionBox
	bg               *bg
	invItemList      *InvItemList
	inventory        *core.Inventory
	currentIteration int
	currentItem      *core.Item
	IsActive         bool
	actionPos        *Pos
	listPos          *Pos
	cursorPos        *Pos
}

func NewInventoryUi() *InventoryUi {
	i := &InventoryUi{
		bg:          NewBg("inventory-bg.png"),
		actionBox:   NewActionBox(),
		cursor:      NewCursor(),
		invItemList: NewInvItemList(),
		actionPos:   &Pos{X: 234, Y: 32},
		listPos:     &Pos{X: 32, Y: 32},
		cursorPos:   &Pos{X: 0, Y: 0},
	}
	return i
}

func (i *InventoryUi) Draw(screen *ebiten.Image) {
	if i.inventory == nil || !i.inventory.IsActive {
		return
	}

	i.bg.Draw(screen)
	i.invItemList.Draw(screen)
	if i.inventory.ActiveElement == "action" {
		i.actionBox.Draw(screen)
	}
	i.cursor.Draw(screen)
}

func (i *InventoryUi) Update(delta int64, state *core.State) {
	if !state.Inventory.IsActive {
		return
	}
	i.inventory = state.Inventory

	// set current item
	if i.inventory.HasItems() {
		i.currentItem = i.inventory.TeamState.GetItem(i.inventory.ItemList[i.inventory.SelectedListIndex])
	} else {
		i.currentItem = nil
	}

	// figure out cursor position
	switch i.inventory.ActiveElement {
	case "list":
		i.cursorPos.X = i.listPos.X - 14
		i.cursorPos.Y = i.listPos.Y + (16.0 * i.inventory.SelectedListIndex)
	case "action":
		i.actionPos.X = i.listPos.X + 8
		i.actionPos.Y = i.listPos.Y + 8 + (16.0 * i.inventory.SelectedListIndex)
		i.cursorPos.X = i.actionPos.X - 14
		i.cursorPos.Y = i.actionPos.Y + 4 + (16.0 * i.inventory.SelectedActionIndex)
	}

	i.cursor.Update(delta, i.cursorPos)
	i.actionBox.Update(delta, i.actionPos, i.currentItem)
	i.invItemList.Update(delta, i.inventory)
}
