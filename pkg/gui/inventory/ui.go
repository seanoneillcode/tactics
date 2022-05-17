package inventory

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
	"github.com/seanoneillcode/go-tactics/pkg/input"
	"log"
)

const offsetX = 4
const offsetY = 4
const charactersPerInfoLine = 24

var invInfoPos = &elem.Pos{
	X: 170,
	Y: 98,
}

var effectPos = &elem.Pos{
	X: 166,
	Y: 16,
}

var itemImagePos = &elem.Pos{
	X: 210,
	Y: 24,
}

var characterCardsPos = &elem.Pos{
	X: 170,
	Y: 20,
}

var actionPos = &elem.Pos{
	X: 120,
	Y: 40,
}

type ctx string

const (
	listCtx      ctx = "list"
	actionCtx    ctx = "action"
	characterCtx ctx = "character"
)

type ui struct {
	cursor                 *elem.Cursor
	actionBox              *ActionBox
	bg                     *elem.StaticImage
	invItemList            *InvItemList
	teamState              *core.TeamState
	activeCtx              ctx // list, action, character
	selectedListIndex      int
	selectedActionIndex    int
	selectedCharacterIndex int
	currentIteration       int
	//currentItem            *core.Item
	IsActive         bool
	actionPos        *elem.Pos
	listPos          *elem.Pos
	cursorPos        *elem.Pos
	infoBox          *elem.Text
	itemImages       map[string]*elem.StaticImage
	currentItemImage *elem.StaticImage
	itemInfoBg       *elem.StaticImage
	cards            map[string]*elem.EffectCard
	uiDesc           *elem.Text
	isLoaded         bool
}

func NewUi() *ui {
	i := &ui{
		bg:               elem.NewStaticImage("uis/inventory/inventory-bg.png", 0, 0),
		cursor:           elem.NewCursor(),
		infoBox:          elem.NewText(invInfoPos.X+2, invInfoPos.Y+12, ""),
		uiDesc:           elem.NewText(223, 4, "Items"),
		actionBox:        NewActionBox(*actionPos),
		invItemList:      NewInvItemList(),
		actionPos:        actionPos,
		listPos:          &elem.Pos{X: 8, Y: 8},
		cursorPos:        &elem.Pos{X: 0, Y: 0},
		itemImages:       map[string]*elem.StaticImage{},
		cards:            map[string]*elem.EffectCard{},
		activeCtx:        listCtx,
		currentItemImage: nil,
		itemInfoBg:       elem.NewStaticImage("uis/inventory/item-info-bg.png", float64(invInfoPos.X-3), float64(invInfoPos.Y)),
	}
	for _, item := range core.AllItems {
		i.itemImages[item.Type] = elem.NewStaticImage(item.ImagePath, float64(itemImagePos.X), float64(itemImagePos.Y))
	}
	return i
}

func (r *ui) Draw(screen *ebiten.Image) {
	if !r.IsActive {
		return
	}

	r.bg.Draw(screen)
	if r.currentItemImage != nil {
		r.itemInfoBg.Draw(screen)
		r.currentItemImage.Draw(screen)
	}
	r.infoBox.Draw(screen)
	r.invItemList.Draw(screen)
	if r.activeCtx == actionCtx {
		r.actionBox.Draw(screen)
	}
	if r.activeCtx == characterCtx {
		for _, card := range r.cards {
			card.Draw(screen)
		}
	}
	r.cursor.Draw(screen)
	r.uiDesc.Draw(screen)
}

func (r *ui) Update(delta int64, state *core.State) {
	if !state.UI.IsInventoryActive() {
		r.IsActive = false
		r.isLoaded = false
		return
	}
	if !r.isLoaded {
		r.isLoaded = true
		r.rebuild(state.TeamState.Characters)
		r.currentItemImage = nil
	}
	r.IsActive = true
	r.handleInput(state)
	r.invItemList.Update(delta, state.TeamState, r.selectedListIndex)

	var formattedItemDescription string
	var item *core.Item
	switch r.activeCtx {
	case listCtx:
		r.cursorPos.X = r.listPos.X - 8
		r.cursorPos.Y = r.listPos.Y + (16.0 * r.selectedListIndex)
		if state.TeamState.HasItems() {
			item = state.TeamState.GetItemWithName(r.invItemList.CurrentItem().Name)
			formattedItemDescription = core.GetFormattedValueMax(item.Description, charactersPerInfoLine)
			r.currentItemImage = r.itemImages[item.Type]
		} else {
			r.currentItemImage = nil
		}
	case actionCtx:
		r.cursorPos.X = r.actionPos.X - 9
		r.cursorPos.Y = r.actionPos.Y + 5 + (24.0 * r.selectedActionIndex)
		if state.TeamState.HasItems() {
			item = state.TeamState.GetItemWithName(r.invItemList.CurrentItem().Name)
			formattedItemDescription = core.GetFormattedValueMax(item.Description, charactersPerInfoLine)
			r.currentItemImage = r.itemImages[item.Type]
		} else {
			r.currentItemImage = nil
		}

	case characterCtx:
		r.cursorPos.X = effectPos.X - 12
		r.cursorPos.Y = effectPos.Y + 16 + (52.0 * r.selectedCharacterIndex)
		r.currentItemImage = nil
		if state.TeamState.HasItems() {
			item = state.TeamState.GetItemWithName(r.invItemList.CurrentItem().Name)
		}
	}

	r.cursor.Update(delta, r.cursorPos)
	r.actionBox.Update(delta)

	r.infoBox.SetValue(formattedItemDescription)

	for _, character := range state.TeamState.Characters {
		r.cards[character.Name].Update(item, character)
	}
}

func (r *ui) handleInput(state *core.State) {
	teamState := state.TeamState
	if teamState == nil {
		log.Fatal("inventory opened with no team!")
	}

	if input.IsCancelPressed() {
		switch r.activeCtx {
		case listCtx:
			r.reset()
			r.IsActive = false
			state.UI.Open(core.MenuUI)
		case actionCtx:
			r.activeCtx = listCtx
		case characterCtx:
			r.activeCtx = actionCtx
		}
		return
	}
	if input.IsMenuPressed() {
		state.UI.Close()
		state.Player.Activate()
		r.reset()
		r.IsActive = false
		return
	}
	if input.IsEnterPressed() {
		switch r.activeCtx {
		case listCtx:
			if teamState.HasItems() {
				item := teamState.GetItemWithName(r.invItemList.CurrentItem().Name)
				if item.CanConsume {
					r.activeCtx = actionCtx
					r.selectedActionIndex = 0
				}
			}
		case actionCtx:
			item := teamState.GetItemWithName(r.invItemList.CurrentItem().Name)
			if r.selectedActionIndex == 0 {
				r.activeCtx = characterCtx
			} else {
				// drop
				teamState.RemoveItem(item.Name)
				r.activeCtx = listCtx
				if r.selectedListIndex == len(teamState.ItemHolders) {
					r.selectedListIndex = r.selectedListIndex - 1
					if r.selectedListIndex < 0 {
						r.selectedListIndex = 0
					}
				}
			}
		case characterCtx:
			item := teamState.GetItemWithName(r.invItemList.CurrentItem().Name)

			// use
			log.Printf("selecting use, item: %v", item.Description)
			if item.CanConsume {
				teamState.ConsumeItem(item.Name, r.selectedCharacterIndex)
			} else {
				if item.CanEquip {
					teamState.EquipItem(item.Name, r.selectedCharacterIndex)
				}
				// ??
			}
			r.activeCtx = listCtx
			if r.selectedListIndex == len(teamState.ItemHolders) {
				r.selectedListIndex = r.selectedListIndex - 1
				if r.selectedListIndex < 0 {
					r.selectedListIndex = 0
				}
			}
			r.rebuild(state.TeamState.Characters)
		}
		return
	}
	if input.IsUpJustPressed() {
		switch r.activeCtx {
		case listCtx:
			r.selectedListIndex = r.selectedListIndex - 1
			if r.selectedListIndex < 0 {
				r.selectedListIndex = 0
			}
		case actionCtx:
			r.selectedActionIndex = r.selectedActionIndex - 1
			if r.selectedActionIndex < 0 {
				r.selectedActionIndex = 0
			}
		case characterCtx:
			r.selectedCharacterIndex = r.selectedCharacterIndex - 1
			if r.selectedCharacterIndex < 0 {
				r.selectedCharacterIndex = 0
			}
		}
	}
	if input.IsDownJustPressed() {
		switch r.activeCtx {
		case listCtx:
			if teamState.HasItems() {
				r.selectedListIndex = r.selectedListIndex + 1
				if r.selectedListIndex == len(teamState.ItemHolders) {
					r.selectedListIndex = r.selectedListIndex - 1
					if r.selectedListIndex < 0 {
						r.selectedListIndex = 0
					}
				}
			} else {
				r.selectedListIndex = 0
			}
		case actionCtx:
			r.selectedActionIndex = r.selectedActionIndex + 1
			if r.selectedActionIndex == 2 {
				r.selectedActionIndex = r.selectedActionIndex - 1
			}
		case characterCtx:
			r.selectedCharacterIndex = r.selectedCharacterIndex + 1
			if r.selectedCharacterIndex == len(teamState.Characters) {
				r.selectedCharacterIndex = r.selectedCharacterIndex - 1
				if r.selectedCharacterIndex < 0 {
					r.selectedCharacterIndex = 0
				}
			}
		}
	}

}

func (r *ui) reset() {
	r.selectedListIndex = 0
	r.selectedActionIndex = 0
}

func (r *ui) rebuild(characters []*core.CharacterState) {
	cards := map[string]*elem.EffectCard{}
	pos := elem.Pos{
		X: characterCardsPos.X,
		Y: characterCardsPos.Y,
	}
	for _, c := range characters {
		cards[c.Name] = elem.NewEffectCard(c, pos)
		pos.Y = pos.Y + 52
	}
	r.cards = cards
}
