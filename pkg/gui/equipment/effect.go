package equipment

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

type Effect struct {
	pos        elem.Pos
	bg         *ebiten.Image
	item       *core.Item
	changeList []*elem.Text
}

func NewEffect(pos elem.Pos) *Effect {
	return &Effect{
		pos: pos,
		bg:  common.LoadImage("uis/equipment/effect.png"),
	}
}

func (r *Effect) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.pos.X), float64(r.pos.Y))
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(r.bg, op)

	for _, ch := range r.changeList {
		ch.Draw(screen)
	}
}

func (r *Effect) Update(item *core.Item, cs *core.CharacterState) {
	if r.item != item {
		r.item = item
		if item != nil && item.Name != elem.RemoveItem {
			r.rebuild(item, cs)
		} else {
			r.changeList = []*elem.Text{}
		}
	}
}

func (r *Effect) rebuild(item *core.Item, cs *core.CharacterState) {
	var changes []*elem.Text
	for _, ef := range item.StatChanges {
		changes = append(changes, elem.NewText(0, 0, ef.Description(cs.EquippedStats)))
	}
	for i, change := range changes {
		change.SetPosition(elem.Pos{
			X: r.pos.X + 8,
			Y: r.pos.Y + (i * 16) + 4,
		})
	}
	r.changeList = changes
}
