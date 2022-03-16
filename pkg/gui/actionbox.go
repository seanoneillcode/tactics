package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

type ActionBox struct {
	useAction   *elem.Button
	equipAction *elem.Button
	dropAction  *elem.Button
	pos         *elem.Pos
	currentItem *core.Item
}

func NewActionBox() *ActionBox {
	return &ActionBox{
		useAction:   elem.NewButton("use", "button-bg.png"),
		equipAction: elem.NewButton("equip", "button-bg.png"),
		dropAction:  elem.NewButton("drop", "button-bg.png"),
		pos:         &elem.Pos{},
	}
}

func (a *ActionBox) Draw(screen *ebiten.Image) {
	a.dropAction.Draw(screen)
	if a.currentItem != nil {
		if a.currentItem.CanConsume {
			a.useAction.Draw(screen)
		} else {
			a.equipAction.Draw(screen)
		}
	}
}

func (a *ActionBox) Update(delta int64, pos *elem.Pos, currentItem *core.Item) {
	a.currentItem = currentItem
	a.pos.X = pos.X
	a.pos.Y = pos.Y
	a.useAction.Update(delta, &elem.Pos{X: pos.X + offsetX, Y: pos.Y + offsetY}, true)
	a.equipAction.Update(delta, &elem.Pos{X: pos.X + offsetX, Y: pos.Y + offsetY}, true)
	a.dropAction.Update(delta, &elem.Pos{X: pos.X + offsetX, Y: pos.Y + 16 + offsetY}, true)
}
