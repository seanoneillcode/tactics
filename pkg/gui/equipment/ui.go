package equipment

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
	"github.com/seanoneillcode/go-tactics/pkg/input"
	"log"
)

type ctx string

const (
	slotCtx          ctx = "slot"
	equipmentListCtx ctx = "equipmentList"
)

var listPos = elem.Pos{
	X: 160,
	Y: 64,
}

var effectPos = elem.Pos{
	X: 160,
	Y: 0,
}

var cardPos = elem.Pos{
	X: 16,
	Y: 16,
}

type ui struct {
	// assets
	bg            *elem.StaticImage
	originalCards []*card
	cards         []*card
	list          *elem.List
	effect        *Effect

	// state
	activeCtx              ctx
	IsActive               bool
	isLoaded               bool
	selectedCharacterIndex int
	currentCardIndex       int
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
		list:   elem.NewList(listPos),
		effect: NewEffect(effectPos),
	}
	return i
}

func createList(teamState *core.TeamState, slot string) []elem.ListItem {
	itemNames := teamState.GetItemList()
	itemMap := teamState.ItemHolders
	var invItems []elem.ListItem
	var offset = 0
	invItems = append(invItems, elem.NewUnEquipItem())
	for _, name := range itemNames {
		teamItem := itemMap[name]
		if teamItem.Item.CanConsume {
			continue
		}
		if teamItem.Item.EquipSlot != slot {
			continue
		}
		entry := elem.NewShopListItem(teamItem.Item, teamItem.Amount)
		invItems = append(invItems, entry)
		offset = offset + 16
	}
	return invItems
}

func (r *ui) Draw(screen *ebiten.Image) {
	if !r.IsActive {
		return
	}
	r.bg.Draw(screen)
	r.cards[r.selectedCharacterIndex].Draw(screen)
	if r.activeCtx == equipmentListCtx {
		r.list.Draw(screen)
		r.effect.Draw(screen)
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
		currentCard(r).selectedSlotIndex = 0 // reset position to start
		r.selectedCharacterIndex = 0
		return
	}
	r.IsActive = true
	r.handleInput(state)

	currentCard(r).Update(cardPos, state.TeamState.Characters[r.selectedCharacterIndex])

	r.list.Update()

	if r.activeCtx == equipmentListCtx {
		item := r.list.CurrentItem()
		r.effect.Update(item, state.TeamState.Characters[r.selectedCharacterIndex])
	}

}

func (r *ui) handleInput(state *core.State) {
	teamState := state.TeamState
	if teamState == nil {
		log.Fatal("equipment opened with no team!")
	}

	if input.IsCancelPressed() {
		switch r.activeCtx {
		case slotCtx:
			state.UI.Open(core.MenuUI)
		case equipmentListCtx:
			r.activeCtx = slotCtx
		}
		return
	}

	if input.IsEnterPressed() {
		switch r.activeCtx {
		case slotCtx:
			newList := createList(teamState, currentSlot(r).SlotType)
			if len(newList) == 0 {
				break
			}
			r.activeCtx = equipmentListCtx
			r.list.SetListItems(newList)
		case equipmentListCtx:
			r.activeCtx = slotCtx
			currentItem := r.list.CurrentItem()
			if currentItem == nil {
				// remove item
				teamState.UnEquipItem(currentSlot(r).SlotType, r.selectedCharacterIndex)
			} else {
				item := teamState.GetItemWithName(currentItem.Name)
				teamState.EquipItem(item.Name, r.selectedCharacterIndex)
			}
		}
		return
	}

	if input.IsRightJustPressed() {
		switch r.activeCtx {
		case slotCtx:
			index := currentCard(r).selectedSlotIndex
			r.selectedCharacterIndex = r.selectedCharacterIndex + 1
			if r.selectedCharacterIndex == 3 {
				r.selectedCharacterIndex = 0
			}
			currentCard(r).selectedSlotIndex = index
		}
		return
	}
	if input.IsLeftJustPressed() {
		switch r.activeCtx {
		case slotCtx:
			index := currentCard(r).selectedSlotIndex
			r.selectedCharacterIndex = r.selectedCharacterIndex - 1
			if r.selectedCharacterIndex == -1 {
				r.selectedCharacterIndex = 2
			}
			currentCard(r).selectedSlotIndex = index
		}
		return
	}

	if input.IsUpJustPressed() {
		switch r.activeCtx {
		case slotCtx:
			currentCard(r).handleInput()
		case equipmentListCtx:
			r.list.HandleInput()
		}

	}
	if input.IsDownJustPressed() {
		switch r.activeCtx {
		case slotCtx:
			currentCard(r).handleInput()
		case equipmentListCtx:
			r.list.HandleInput()
		}
	}
}

func currentSlot(r *ui) *slotEntry {
	return r.cards[r.selectedCharacterIndex].currentSlot()
}

func currentCard(r *ui) *card {
	return r.cards[r.selectedCharacterIndex]
}
