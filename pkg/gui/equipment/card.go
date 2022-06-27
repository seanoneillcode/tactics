package equipment

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/explore"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
	"github.com/seanoneillcode/go-tactics/pkg/input"
)

type card struct {
	portrait          *elem.Sprite
	slots             []*slotEntry
	selectedSlotIndex int
}

func NewCard(name string) *card {
	return &card{
		portrait: elem.NewSprite(fmt.Sprintf("portrait/%s.png", name), 0, 0),
		slots: []*slotEntry{
			NewSlotEntry("weapon"),
			NewSlotEntry("armor"),
			NewSlotEntry("special"),
		},
	}
}

func (r *card) Update(pos elem.Pos, charState *explore.CharacterState) {
	for index, s := range r.slots {
		isHighlighted := index == r.selectedSlotIndex
		s.Update(elem.Pos{
			X: pos.X,
			Y: 72 + (index * 24),
		},
			isHighlighted,
			charState,
		)
	}
	r.portrait.SetPos(elem.Pos{
		X: pos.X + 40,
		Y: pos.Y,
	})
}

func (r *card) Draw(screen *ebiten.Image) {
	for _, s := range r.slots {
		s.Draw(screen)
	}
	r.portrait.Draw(screen)
}

func (r *card) handleInput() {
	if input.IsUpJustPressed() {
		r.selectedSlotIndex = r.selectedSlotIndex - 1
		if r.selectedSlotIndex < 0 {
			r.selectedSlotIndex = 0
		}
		return
	}
	if input.IsDownJustPressed() {
		r.selectedSlotIndex = r.selectedSlotIndex + 1
		if r.selectedSlotIndex == 3 {
			r.selectedSlotIndex = r.selectedSlotIndex - 1
		}
		return
	}
}

func (r *card) currentSlot() *slotEntry {
	return r.slots[r.selectedSlotIndex]
}

type slotEntry struct {
	image       *elem.Sprite
	bg          *elem.Sprite
	bgHighlight *elem.Sprite

	SlotType      string
	label         *elem.Text
	isHighlighted bool
}

func NewSlotEntry(slotType string) *slotEntry {
	return &slotEntry{
		SlotType:    slotType,
		bg:          elem.NewSprite("slots/bg.png", 0, 0),
		bgHighlight: elem.NewSprite("slots/bg-highlight.png", 0, 0),
		image:       elem.NewSprite(fmt.Sprintf("slots/%s.png", slotType), 0, 0),
		label:       elem.NewText(0, 0, "nothing"),
	}
}

func (r *slotEntry) Draw(screen *ebiten.Image) {
	if r.isHighlighted {
		r.bgHighlight.Draw(screen)
	} else {
		r.bg.Draw(screen)
	}
	r.image.Draw(screen)
	r.label.Draw(screen)
}

func (r *slotEntry) Update(pos elem.Pos, isHighlighted bool, charState *explore.CharacterState) {
	r.isHighlighted = isHighlighted
	r.bg.SetPos(pos)
	r.bgHighlight.SetPos(pos)
	r.label.SetPosition(elem.Pos{
		X: pos.X + 16 + 4,
		Y: pos.Y + 6,
	})
	item, ok := charState.EquippedItems[r.SlotType]
	if !ok {
		r.label.SetValue("nothing")
	} else {
		r.label.SetValue(item.Name)
	}
	r.image.SetPos(elem.Pos{
		X: pos.X + 2,
		Y: pos.Y + 2,
	})
}
