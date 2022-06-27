package elem

import (
	"fmt"
	"github.com/seanoneillcode/go-tactics/pkg/input"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/explore"
)

const RemoveItem = "Remove Equipment"

type List struct {
	// assets
	pos       *Pos
	bg        *StaticImage
	highlight *Sprite

	// state
	itemList         []ListItem
	currentIteration int
	index            int
}

func NewList(pos Pos) *List {
	return &List{
		pos:       &Pos{X: pos.X, Y: pos.Y},
		bg:        NewStaticImage("uis/equipment/list-bg.png", float64(pos.X), float64(pos.Y)),
		highlight: NewSprite("uis/equipment/highlight.png", 0, 0),
	}
}

func (r *List) SetListItems(itemList []ListItem) {
	r.itemList = itemList
	r.index = 0
}

func (r *List) Draw(screen *ebiten.Image) {
	r.bg.Draw(screen)
	r.highlight.Draw(screen)
	for _, item := range r.itemList {
		item.Draw(screen)
	}
}

func (r *List) Update() {
	var x, y = r.pos.X + 4, r.pos.Y + 4
	for index, item := range r.itemList {
		item.Update(Pos{
			X: x,
			Y: y + (16 * index),
		})
	}
	r.highlight.SetPos(Pos{
		X: r.pos.X,
		Y: r.pos.Y + (16 * r.index),
	})
}

func (r *List) CurrentItem() *explore.Item {
	return r.itemList[r.index].Item()
}

func (r *List) HandleInput() {
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

type ListItem interface {
	Draw(screen *ebiten.Image)
	Update(pos Pos)
	Item() *explore.Item
}

type InventoryListItem struct {
	itemRef *explore.Item
	text    *Text
}

func NewShopListItem(item *explore.Item, amount int) *InventoryListItem {
	value := fmt.Sprintf("%2d - %s", amount, item.Name)
	return &InventoryListItem{
		itemRef: item,
		text:    NewText(0, 0, value),
	}
}

func (r *InventoryListItem) Draw(screen *ebiten.Image) {
	r.text.Draw(screen)
}

func (r *InventoryListItem) Update(pos Pos) {
	r.text.SetPosition(pos)
}

func (r *InventoryListItem) Item() *explore.Item {
	return r.itemRef
}

type UnEquipItem struct {
	text *Text
}

func NewUnEquipItem() *UnEquipItem {
	return &UnEquipItem{
		text: NewText(0, 0, "Remove Equipment"),
	}
}

func (r *UnEquipItem) Draw(screen *ebiten.Image) {
	r.text.Draw(screen)
}

func (r *UnEquipItem) Update(pos Pos) {
	r.text.SetPosition(pos)
}

func (r *UnEquipItem) Item() *explore.Item {
	return nil
}
