package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
)

const (
	listItemX = 32
	listItemY = 32
	actionX   = 208
	actionY   = 32
)

type InventoryUi struct {
	inventory       *core.Inventory
	bgImage         *ebiten.Image
	shopCursorImage *ebiten.Image
	isDirty         bool
	itemList        []*invListItem
	cursorOffset    float64
	cursorTimer     int64
	oldNumItems     int

	useAction   *Text
	equipAction *Text
	dropAction  *Text
	currentItem *core.Item
}

func NewInventoryUi() *InventoryUi {
	i := &InventoryUi{
		bgImage:         core.LoadImage("inventory-bg.png"),
		shopCursorImage: core.LoadImage("shop-cursor.png"),
		useAction: &Text{
			value: "use",
			x:     actionX + offsetX,
			y:     actionY + offsetY,
			color: defaultTextColor,
		},
		equipAction: &Text{
			value: "equip",
			x:     actionX + offsetX,
			y:     actionY + offsetY,
			color: defaultTextColor,
		},
		dropAction: &Text{
			value: "drop",
			x:     actionX + 64 + offsetX,
			y:     actionY + offsetY,
			color: defaultTextColor,
		},
	}
	return i
}

func (i *InventoryUi) Draw(screen *ebiten.Image) {
	if i.inventory == nil || !i.inventory.IsActive {
		return
	}

	// background
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(i.bgImage, op)

	// list of items
	for _, item := range i.itemList {
		item.Draw(screen)
	}

	// actions for an item
	if i.currentItem != nil {
		activeColor := defaultTextColor
		switch i.inventory.ActiveElement {
		case "list":
			activeColor = greyTextColor
			// cursor
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(listItemX-14+i.cursorOffset, listItemY+(float64)(16.0*i.inventory.SelectedListIndex))
			op.GeoM.Scale(common.Scale, common.Scale)
			screen.DrawImage(i.shopCursorImage, op)
		case "action":
			activeColor = defaultTextColor

			// cursor
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(actionX-14+i.cursorOffset+(float64)(64.0*i.inventory.SelectedActionIndex), actionY)
			op.GeoM.Scale(common.Scale, common.Scale)
			screen.DrawImage(i.shopCursorImage, op)

		}
		i.dropAction.SetColor(activeColor)
		i.dropAction.Draw(screen)
		if i.currentItem.CanConsume {
			i.useAction.SetColor(activeColor)
			i.useAction.Draw(screen)
		} else {
			i.equipAction.SetColor(activeColor)
			i.equipAction.Draw(screen)
		}
	}

}

func (i *InventoryUi) Update(delta int64, state *core.State) {
	if !state.Inventory.IsActive {
		i.isDirty = true
		return
	}
	i.inventory = state.Inventory
	if i.oldNumItems != len(i.inventory.TeamState.Items) {
		i.isDirty = true
		i.oldNumItems = len(i.inventory.TeamState.Items)
	}
	if i.inventory.IsActive && i.isDirty {
		i.isDirty = false
		// create inventory
		i.itemList = i.createItemList(listItemX, listItemY)
	}

	if i.inventory.HasItems() {
		i.currentItem = i.inventory.TeamState.Items[i.inventory.SelectedListIndex]
	} else {
		i.currentItem = nil
	}

	i.cursorTimer = i.cursorTimer + delta
	if i.cursorTimer > 400 {
		i.cursorTimer = i.cursorTimer - 400
		if i.cursorOffset == 0 {
			i.cursorOffset = 2
		} else {
			i.cursorOffset = 0
		}
	}
}

func (i *InventoryUi) createItemList(x int, y int) []*invListItem {
	x = x + offsetX
	y = y + offsetY
	var invItems []*invListItem
	var offset = 0
	for _, item := range i.inventory.TeamState.Items {
		//costWidth := text.BoundString(standardFont, quantity).Size().X / common.ScaleF
		invItems = append(invItems, &invListItem{
			itemRef: item,
			name: &Text{
				value: item.Name,
				x:     x,
				y:     y + offset,
				color: defaultTextColor,
			},
			//quantity: &Text{
			//	value: quantity,
			//	x:     x + 96 + 32 + offsetX + offsetX - costWidth,
			//	y:     y + offset,
			//	color: defaultTextColor,
			//},
		})
		offset = offset + 16

	}
	return invItems
}

type invListItem struct {
	itemRef *core.Item
	name    *Text
	//quantity *Text
}

func (l *invListItem) Draw(screen *ebiten.Image) {
	l.name.Draw(screen)
	//l.quantity.Draw(screen)
}
