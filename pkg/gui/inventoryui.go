package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

var invInfoPos = &elem.Pos{
	X: 208,
	Y: 80,
}

type InventoryUi struct {
	cursor           *elem.Cursor
	actionBox        *ActionBox
	bg               *elem.StaticImage
	invItemList      *InvItemList
	inventory        *core.Inventory
	currentIteration int
	currentItem      *core.Item
	IsActive         bool
	actionPos        *elem.Pos
	listPos          *elem.Pos
	cursorPos        *elem.Pos
	infoBox          *elem.InfoBox
}

func NewInventoryUi() *InventoryUi {
	i := &InventoryUi{
		bg:          elem.NewStaticImage("inventory-bg.png", 0, 0),
		cursor:      elem.NewCursor(),
		infoBox:     elem.NewInfoBox("", "shop-information-bg.png"),
		actionBox:   NewActionBox(),
		invItemList: NewInvItemList(),
		actionPos:   &elem.Pos{X: 234, Y: 32},
		listPos:     &elem.Pos{X: 32, Y: 32},
		cursorPos:   &elem.Pos{X: 0, Y: 0},
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
	i.infoBox.Draw(screen)
	i.cursor.Draw(screen)
}

func (i *InventoryUi) Update(delta int64, state *core.State) {
	if !state.Inventory.IsActive {
		return
	}
	i.inventory = state.Inventory

	var drawInfoBox bool
	var formattedItemDescription string
	// figure out cursor, actionBox positions
	switch i.inventory.ActiveElement {
	case "list":
		i.cursorPos.X = i.listPos.X - 14
		i.cursorPos.Y = i.listPos.Y + (16.0 * i.inventory.SelectedListIndex)
	case "action":
		i.actionPos.X = i.listPos.X + 2
		i.actionPos.Y = i.listPos.Y + 11 + (16.0 * i.inventory.SelectedListIndex)
		i.cursorPos.X = i.actionPos.X - 9
		i.cursorPos.Y = i.actionPos.Y + 5 + (16.0 * i.inventory.SelectedActionIndex)
		if i.inventory.HasItems() {
			item := i.inventory.TeamState.GetItemWithIndex(i.inventory.SelectedListIndex)
			drawInfoBox = true
			formattedItemDescription = core.GetFormattedValueMax(item.Description, 22)
		}
	}

	i.cursor.Update(delta, i.cursorPos)
	i.actionBox.Update(delta, i.actionPos, i.inventory)
	i.invItemList.Update(delta, i.inventory)
	i.infoBox.Update(invInfoPos, drawInfoBox, formattedItemDescription)
}
