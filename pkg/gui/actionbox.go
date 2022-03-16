package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/core"
)

type ActionBox struct {
	useAction   *button
	equipAction *button
	dropAction  *button
	pos         *Pos
	currentItem *core.Item
}

func NewActionBox() *ActionBox {
	return &ActionBox{
		useAction:   NewButton("use", "button-bg.png", activeColor),
		equipAction: NewButton("equip", "button-bg.png", activeColor),
		dropAction:  NewButton("drop", "button-bg.png", activeColor),
		pos:         &Pos{},
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

func (a *ActionBox) Update(delta int64, pos *Pos, currentItem *core.Item) {
	a.currentItem = currentItem
	a.pos.X = pos.X
	a.pos.Y = pos.Y
	a.useAction.Update(delta, &Pos{X: pos.X + offsetX, Y: pos.Y + offsetY})
	a.equipAction.Update(delta, &Pos{X: pos.X + offsetX, Y: pos.Y + offsetY})
	a.dropAction.Update(delta, &Pos{X: pos.X + offsetX, Y: pos.Y + 16 + offsetY})
}
