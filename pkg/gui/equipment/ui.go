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

type ui struct {
	// assets
	bg            *elem.StaticImage
	originalCards []*card
	cards         []*card

	// state
	activeCtx              ctx
	IsActive               bool
	isLoaded               bool
	selectedCharacterIndex int
}

func NewUI() *ui {
	i := &ui{
		bg:        elem.NewStaticImage("equipment-bg.png", 0, 0),
		activeCtx: slotCtx,
		originalCards: []*card{
			NewCard("alice"),
			NewCard("bob"),
			NewCard("carl"),
		},
		cards: []*card{},
	}
	i.rebuild()
	return i
}

func (r *ui) Draw(screen *ebiten.Image) {
	if !r.IsActive {
		return
	}

	r.bg.Draw(screen)
	for _, c := range r.cards {
		c.Draw(screen)
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
		//r.rebuild(state.TeamState.Characters)
	}
	r.IsActive = true
	r.handleInput(state)

	for index, c := range r.cards {
		isSelected := index == 0
		pos := elem.Pos{
			X: index * 100,
			Y: 0,
		}

		c.Update(pos, isSelected, state.TeamState.Characters[r.selectedCharacterIndex])
	}
}

func (r *ui) handleInput(state *core.State) {
	teamState := state.TeamState
	if teamState == nil {
		log.Fatal("equipment opened with no team!")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		switch r.activeCtx {
		case slotCtx:
			//r.reset()
			r.IsActive = false
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
		case equipmentListCtx:
			r.activeCtx = slotCtx
			// equip item to slot
			//item := teamState.GetItemWithIndex(r.selectedSlotIndex)
			//teamState.EquipItem(item.Name, r.selectedCharacterIndex)
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
			r.rebuild()
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
			r.rebuild()
		}
		return
	}

	//switch r.activeCtx {
	//case slotCtx:
	//	r.cards[r.selectedCharacterIndex].handleInput()
	//}
}

func (r *ui) rebuild() {
	var newCards []*card

	for index := range r.originalCards {
		newCards = append(newCards, r.originalCards[(index+r.selectedCharacterIndex)%3])
	}

	r.cards = newCards
}
