package elem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
)

type EffectCard struct {
	pos        *Pos
	image      *ebiten.Image
	bg         *ebiten.Image
	name       *Text
	isDraw     bool
	item       *core.Item
	changeList []*Text
}

func NewEffectCard(imageFileName string) *EffectCard {
	r := &EffectCard{
		pos: &Pos{},
	}
	r.bg = common.LoadImage(imageFileName)
	r.name = NewText(0, 0, "")
	return r
}

func (r *EffectCard) Draw(screen *ebiten.Image) {
	if !r.isDraw {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.pos.X), float64(r.pos.Y))
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(r.bg, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.pos.X+2), float64(r.pos.Y+2))
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(r.image, op)

	r.name.SetPosition(Pos{X: r.pos.X + offsetX, Y: r.pos.Y + 32 + offsetY})
	r.name.Draw(screen)

	for _, ch := range r.changeList {
		ch.Draw(screen)
	}
}

func (r *EffectCard) Update(pos *Pos, isDraw bool, name string, charImage *ebiten.Image, item *core.Item, cs *core.CharacterState) {
	r.pos.X = pos.X
	r.pos.Y = pos.Y
	r.isDraw = isDraw
	r.name.SetValue(name)
	r.image = charImage
	if r.item != item {
		r.item = item
		if item != nil {
			r.rebuild(item, cs, pos)
		} else {
			r.changeList = []*Text{}
		}
	}
}

func (r *EffectCard) rebuild(item *core.Item, cs *core.CharacterState, pos *Pos) {
	var changes []*Text
	for _, ef := range item.Effects {
		changes = append(changes, NewText(0, 0, ef.Description(cs)))
	}
	for i, change := range changes {
		change.SetPosition(Pos{
			X: pos.X + 48,
			Y: pos.Y + (i * 16) + 4,
		})
	}
	r.changeList = changes
}
