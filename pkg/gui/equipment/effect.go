package equipment

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/explore"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

type Effect struct {
	pos        elem.Pos
	bg         *ebiten.Image
	item       *explore.Item
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

func (r *Effect) Update(item *explore.Item, cs *explore.CharacterState) {
	if r.item != item {
		r.item = item
		if item != nil && item.Name != elem.RemoveItem {
			r.rebuild(item, cs)
		} else {
			r.changeList = []*elem.Text{}
		}
	}
}

func (r *Effect) rebuild(item *explore.Item, cs *explore.CharacterState) {
	// copy over everything except current item for slot
	minusCurrentItem := map[string]*explore.Item{}
	for _, ei := range cs.EquippedItems {
		minusCurrentItem[ei.EquipSlot] = ei
	}
	minusCurrentItem[item.EquipSlot] = item
	newStats := cs.ApplyItemsToStats(minusCurrentItem, cs.BaseStats)

	var changes []*elem.Text

	if newStats.AttackSkill != cs.EquippedStats.AttackSkill {
		msg := fmt.Sprintf("attack skill %v > %v", cs.EquippedStats.AttackSkill, newStats.AttackSkill)
		changes = append(changes, elem.NewText(0, 0, msg))
	}
	if newStats.AttackStrength != cs.EquippedStats.AttackStrength {
		msg := fmt.Sprintf("attack strength %v > %v", cs.EquippedStats.AttackStrength, newStats.AttackStrength)
		changes = append(changes, elem.NewText(0, 0, msg))
	}
	if newStats.Speed != cs.EquippedStats.Speed {
		msg := fmt.Sprintf("speed %v > %v", cs.EquippedStats.Speed, newStats.Speed)
		changes = append(changes, elem.NewText(0, 0, msg))
	}
	if newStats.Agility != cs.EquippedStats.Agility {
		msg := fmt.Sprintf("agility %v > %v", cs.EquippedStats.Agility, newStats.Agility)
		changes = append(changes, elem.NewText(0, 0, msg))
	}
	if newStats.Defence != cs.EquippedStats.Defence {
		msg := fmt.Sprintf("defense %v > %v", cs.EquippedStats.Defence, newStats.Defence)
		changes = append(changes, elem.NewText(0, 0, msg))
	}
	if newStats.MagicSkill != cs.EquippedStats.MagicSkill {
		msg := fmt.Sprintf("magic skill %v > %v", cs.EquippedStats.MagicSkill, newStats.MagicSkill)
		changes = append(changes, elem.NewText(0, 0, msg))
	}
	if newStats.MagicDef != cs.EquippedStats.MagicDef {
		msg := fmt.Sprintf("magic defense %v > %v", cs.EquippedStats.MagicDef, newStats.MagicDef)
		changes = append(changes, elem.NewText(0, 0, msg))
	}
	if newStats.MaxMagic != cs.EquippedStats.MaxMagic {
		msg := fmt.Sprintf("max magic %v > %v", cs.EquippedStats.MaxMagic, newStats.MaxMagic)
		changes = append(changes, elem.NewText(0, 0, msg))
	}
	if newStats.MaxHealth != cs.EquippedStats.MaxHealth {
		msg := fmt.Sprintf("max health %v > %v", cs.EquippedStats.MaxHealth, newStats.MaxHealth)
		changes = append(changes, elem.NewText(0, 0, msg))
	}

	if len(changes) == 0 {
		changes = append(changes, elem.NewText(0, 0, "no change"))
	}
	for i, change := range changes {
		change.SetPosition(elem.Pos{
			X: r.pos.X + 8,
			Y: r.pos.Y + (i * 16) + 4,
		})
	}
	r.changeList = changes
}
