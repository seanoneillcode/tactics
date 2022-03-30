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
}

func NewActionBox() *ActionBox {
	return &ActionBox{
		useAction:   elem.NewButton("use", "button-bg.png"),
		equipAction: elem.NewButton("equip", "button-bg.png"),
		dropAction:  elem.NewButton("drop", "button-bg.png"),
	}
}

func (a *ActionBox) Draw(screen *ebiten.Image) {
	a.dropAction.Draw(screen)
	a.useAction.Draw(screen)
	a.equipAction.Draw(screen)
}

func (a *ActionBox) Update(delta int64, pos *elem.Pos, currentItem *core.Item) {
	showUse := currentItem != nil && currentItem.CanConsume
	disableUse := currentItem == nil
	showEquip := !showUse

	a.useAction.Update(delta, &elem.Pos{X: pos.X + offsetX, Y: pos.Y + offsetY}, disableUse, showUse)
	a.equipAction.Update(delta, &elem.Pos{X: pos.X + offsetX, Y: pos.Y + offsetY}, false, showEquip)
	a.dropAction.Update(delta, &elem.Pos{X: pos.X + offsetX, Y: pos.Y + 16 + offsetY}, false, true)
}
