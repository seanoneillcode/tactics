package elem

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/explore"
)

type EffectCard struct {
	pos        *Pos
	image      *ebiten.Image
	bg         *ebiten.Image
	name       *Text
	item       *explore.Item
	changeList []*Text
}

func NewEffectCard(cs *explore.CharacterState, pos Pos) *EffectCard {
	r := &EffectCard{
		pos:   &Pos{pos.X, pos.Y},
		image: common.LoadImage(fmt.Sprintf("portrait/%s.png", cs.Name)),
		bg:    common.LoadImage("elem/effectcard-bg.png"),
		name:  NewText(0, 0, cs.Name),
	}

	return r
}

func (r *EffectCard) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.pos.X+48), float64(r.pos.Y))
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(r.bg, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.pos.X), float64(r.pos.Y))
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(r.image, op)

	r.name.SetPosition(Pos{X: r.pos.X + offsetX, Y: r.pos.Y + 32 + offsetY})
	r.name.Draw(screen)

	for _, ch := range r.changeList {
		ch.Draw(screen)
	}
}

func (r *EffectCard) Update(item *explore.Item, cs *explore.CharacterState) {
	if r.item != item {
		r.item = item
		if item != nil {
			r.rebuild(item, cs)
		} else {
			r.changeList = []*Text{}
		}
	}
}

func (r *EffectCard) rebuild(item *explore.Item, cs *explore.CharacterState) {
	var changes []*Text
	for _, ef := range item.Effects {
		changes = append(changes, NewText(0, 0, ef.Description(cs)))
	}
	for i, change := range changes {
		change.SetPosition(Pos{
			X: r.pos.X + 48 + 8,
			Y: r.pos.Y + (i * 16) + 4,
		})
	}
	r.changeList = changes
}
