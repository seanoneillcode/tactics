package equipment

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

type card struct {
	portrait          *elem.Sprite
	slots             []*slotEntry
	selectedSlotIndex int
	isSelected        bool
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

func (r *card) Update(pos elem.Pos, isSelected bool, charState *core.CharacterState) {
	r.isSelected = isSelected
	for index, s := range r.slots {
		isHighlighted := index == r.selectedSlotIndex
		s.Update(elem.Pos{
			X: pos.X + 4,
			Y: 72 + (pos.Y+24)*index,
		},
			isHighlighted,
			charState,
		)
	}
	r.portrait.SetPos(&elem.Pos{
		X: pos.X + 24 + 1,
		Y: pos.Y + 8 + 1,
	})
}

func (r *card) Draw(screen *ebiten.Image) {
	if r.isSelected {
		for _, s := range r.slots {
			s.Draw(screen)
		}
	}
	r.portrait.Draw(screen)
}

func (r *card) handleInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		r.selectedSlotIndex = r.selectedSlotIndex - 1
		if r.selectedSlotIndex < 0 {
			r.selectedSlotIndex = 0
		}
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		r.selectedSlotIndex = r.selectedSlotIndex + 1
		if r.selectedSlotIndex == 3 {
			r.selectedSlotIndex = r.selectedSlotIndex - 1
		}
		return
	}
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

func (r *slotEntry) Update(pos elem.Pos, isHighlighted bool, charState *core.CharacterState) {
	r.isHighlighted = isHighlighted
	r.bg.SetPos(&pos)
	r.bgHighlight.SetPos(&pos)
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
	r.image.SetPos(&elem.Pos{
		X: pos.X + 2,
		Y: pos.Y + 2,
	})
}
