package equipment

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
	"log"
)

type ctx string

const (
	slotCtx          ctx = "slot"
	equipmentListCtx ctx = "equipmentList"
)

var listPos = elem.Pos{
	X: 120,
	Y: 0,
}

type ui struct {
	// assets
	bg            *elem.StaticImage
	originalCards []*card
	cards         []*card
	list          *List

	// state
	activeCtx              ctx
	IsActive               bool
	isLoaded               bool
	selectedCharacterIndex int
}

func NewUI() *ui {
	i := &ui{
		bg:        elem.NewStaticImage("uis/equipment/bg.png", 0, 0),
		activeCtx: slotCtx,
		cards: []*card{
			NewCard("alice"),
			NewCard("bob"),
			NewCard("carl"),
		},
		list: NewList(listPos),
	}
	return i
}

func (r *ui) rotateCards(amount int) {
	length := len(r.cards)
	var tmp = make([]*card, length)
	for index, _ := range r.cards {
		newIndex := index + amount
		if newIndex == length {
			newIndex = 0
		}
		if newIndex == -1 {
			newIndex = length - 1
		}
		tmp[newIndex] = r.cards[index]
	}
	r.cards = tmp
}

func (r *ui) Draw(screen *ebiten.Image) {
	if !r.IsActive {
		return
	}

	r.bg.Draw(screen)
	for _, c := range r.cards {
		c.Draw(screen)
	}
	if r.activeCtx == equipmentListCtx {
		r.list.Draw(screen)
	}
}

func (r *ui) Update(delta int64, state *core.State) {
	if !state.UI.IsEquipmentActive() {
		r.IsActive = false
		r.isLoaded = false
		return
	}
	if !r.isLoaded {
		r.isLoaded = true
		return
	}
	r.IsActive = true
	r.handleInput(state)

	for index, c := range r.cards {
		isSelected := index == 0
		pos := elem.Pos{
			X: index * 110,
			Y: 0,
		}
		c.Update(pos, isSelected, state.TeamState.Characters[r.selectedCharacterIndex])
	}

	r.list.Update(state.TeamState, r.currentSlot().SlotType)
}

func (r *ui) handleInput(state *core.State) {
	teamState := state.TeamState
	if teamState == nil {
		log.Fatal("equipment opened with no team!")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		switch r.activeCtx {
		case slotCtx:
			state.UI.Open(core.MenuUI)
		case equipmentListCtx:
			r.activeCtx = slotCtx
		}
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch r.activeCtx {
		case slotCtx:
			r.activeCtx = equipmentListCtx
			r.list.updateList(teamState, r.currentSlot().SlotType)
		case equipmentListCtx:
			r.activeCtx = slotCtx
			item := teamState.GetItemWithName(r.list.currentItem().itemRef.Name)
			teamState.EquipItem(item.Name, r.selectedCharacterIndex)
		}
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		switch r.activeCtx {
		case slotCtx:
			r.selectedCharacterIndex = r.selectedCharacterIndex + 1
			if r.selectedCharacterIndex == 3 {
				r.selectedCharacterIndex = 0
			}
			r.rotateCards(1)
		}
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		switch r.activeCtx {
		case slotCtx:
			r.selectedCharacterIndex = r.selectedCharacterIndex - 1
			if r.selectedCharacterIndex == -1 {
				r.selectedCharacterIndex = 2
			}
			r.rotateCards(-1)
		}
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		switch r.activeCtx {
		case slotCtx:
			r.cards[0].handleInput()
		case equipmentListCtx:
			r.list.handleInput()
		}

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		switch r.activeCtx {
		case slotCtx:
			r.cards[0].handleInput()
		case equipmentListCtx:
			r.list.handleInput()
		}
	}
}

func (r *ui) currentSlot() *slotEntry {
	return r.cards[0].currentSlot()
}
