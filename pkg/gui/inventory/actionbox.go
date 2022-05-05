package inventory

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

type ActionBox struct {
	useAction  *elem.Button
	dropAction *elem.Button
	bg         *elem.StaticImage
	pos        elem.Pos
}

func NewActionBox(pos elem.Pos) *ActionBox {
	return &ActionBox{
		useAction:  elem.NewButton("Use", "uis/inventory/button-bg.png"),
		dropAction: elem.NewButton("Drop", "uis/inventory/button-bg.png"),
		bg:         elem.NewStaticImage("uis/inventory/action-bg.png", float64(pos.X), float64(pos.Y)),
		pos: elem.Pos{
			X: pos.X,
			Y: pos.Y,
		},
	}
}

func (a *ActionBox) Draw(screen *ebiten.Image) {
	a.bg.Draw(screen)
	a.dropAction.Draw(screen)
	a.useAction.Draw(screen)
}

func (a *ActionBox) Update(delta int64) {
	a.useAction.Update(delta, &elem.Pos{X: a.pos.X + offsetX, Y: a.pos.Y + offsetY}, false, true)
	a.dropAction.Update(delta, &elem.Pos{X: a.pos.X + offsetX, Y: a.pos.Y + 20 + offsetY + offsetY}, false, true)
}
