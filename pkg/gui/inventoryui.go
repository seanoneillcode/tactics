package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
	"log"
)

var invInfoPos = &elem.Pos{
	X: 200,
	Y: 100,
}

var effectPos = &elem.Pos{
	X: 200,
	Y: 32,
}

type InventoryUi struct {
	cursor              *elem.Cursor
	actionBox           *ActionBox
	bg                  *elem.StaticImage
	invItemList         *InvItemList
	teamState           *core.TeamState
	inventory           *core.Inventory
	justOpened          bool
	ActiveElement       string // list, action, character
	SelectedListIndex   int
	SelectedActionIndex int
	SelectedCharacter   string
	currentIteration    int
	currentItem         *core.Item
	IsActive            bool
	actionPos           *elem.Pos
	listPos             *elem.Pos
	cursorPos           *elem.Pos
	infoBox             *elem.InfoBox
	characterEffect     *elem.EffectCard
	charImages          map[string]*ebiten.Image
}

func NewInventoryUi(inventory *core.Inventory) *InventoryUi {
	i := &InventoryUi{
		bg:          elem.NewStaticImage("inventory-bg.png", 0, 0),
		cursor:      elem.NewCursor(),
		infoBox:     elem.NewInfoBox("", "shop-information-bg.png"),
		actionBox:   NewActionBox(),
		invItemList: NewInvItemList(),
		actionPos:   &elem.Pos{X: 234, Y: 32},
		listPos:     &elem.Pos{X: 32, Y: 32},
		cursorPos:   &elem.Pos{X: 0, Y: 0},
		charImages: map[string]*ebiten.Image{
			"default": common.LoadImage("default-avatar.png"),
		},
		characterEffect:   elem.NewEffectCard("effect-card-bg.png"),
		inventory:         inventory,
		ActiveElement:     "list",
		SelectedListIndex: 0,
		SelectedCharacter: "default",
	}
	return i
}

func (i *InventoryUi) Draw(screen *ebiten.Image) {
	if !i.inventory.IsActive {
		return
	}

	i.bg.Draw(screen)
	i.invItemList.Draw(screen)
	if i.ActiveElement == "action" {
		i.actionBox.Draw(screen)
	}
	if i.ActiveElement == "character" {
		i.actionBox.Draw(screen)
		i.characterEffect.Draw(screen)
	}
	i.infoBox.Draw(screen)
	i.cursor.Draw(screen)
}

func (i *InventoryUi) Update(delta int64, state *core.State) {
	if !state.Inventory.IsActive {
		return
	}
	i.inventory = state.Inventory
	i.UpdateState(delta, state)

	var drawInfoBox bool
	var formattedItemDescription string
	var item *core.Item
	// figure out cursor, actionBox positions
	switch i.ActiveElement {
	case "list":
		i.cursorPos.X = i.listPos.X - 14
		i.cursorPos.Y = i.listPos.Y + (16.0 * i.SelectedListIndex)
	case "action":
		i.actionPos.X = i.listPos.X + 2
		i.actionPos.Y = i.listPos.Y + 11 + (16.0 * i.SelectedListIndex)
		i.cursorPos.X = i.actionPos.X - 9
		i.cursorPos.Y = i.actionPos.Y + 5 + (16.0 * i.SelectedActionIndex)
		if i.inventory.HasItems() {
			item = state.Player.TeamState.GetItemWithIndex(i.SelectedListIndex)
			drawInfoBox = true
			formattedItemDescription = core.GetFormattedValueMax(item.Description, 22)
		}
	case "character":
		i.cursorPos.X = effectPos.X - 12
		i.cursorPos.Y = effectPos.Y + 16
		if i.inventory.HasItems() {
			item = state.Player.TeamState.GetItemWithIndex(i.SelectedListIndex)
			drawInfoBox = true
			formattedItemDescription = core.GetFormattedValueMax(item.Description, 22)
		}
	}

	var currentItem *core.Item
	if state.Player.TeamState.HasItems() {
		currentItem = state.Player.TeamState.GetItemWithIndex(i.SelectedListIndex)
	}

	i.cursor.Update(delta, i.cursorPos)
	i.actionBox.Update(delta, i.actionPos, currentItem)
	i.invItemList.Update(delta, state.Player.TeamState)
	i.infoBox.Update(invInfoPos, drawInfoBox, formattedItemDescription)
	i.characterEffect.Update(effectPos, true, i.SelectedCharacter, i.charImages[i.SelectedCharacter], item, state.Player.TeamState.Characters[0])
}

func (i *InventoryUi) UpdateState(delta int64, state *core.State) {
	if !state.Inventory.IsActive {
		i.justOpened = true
		return
	}
	if i.justOpened {
		i.justOpened = false
		return
	}
	teamState := state.Player.TeamState
	if teamState == nil {
		log.Fatal("inventory opened with no team!")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		switch i.ActiveElement {
		case "list":
			state.Inventory.Close()
			state.Player.Activate()
		case "action":
			i.ActiveElement = "list"
		case "character":
			i.ActiveElement = "action"
		}
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch i.ActiveElement {
		case "list":
			if teamState.HasItems() {
				i.ActiveElement = "action"
				i.SelectedActionIndex = 0
			}
		case "action":
			item := teamState.GetItemWithIndex(i.SelectedListIndex)
			if i.SelectedActionIndex == 0 {
				i.ActiveElement = "character"
			} else {
				// drop
				teamState.RemoveItem(item.Name)
				i.ActiveElement = "list"
				if i.SelectedListIndex == len(teamState.ItemHolders) {
					i.SelectedListIndex = i.SelectedListIndex - 1
					if i.SelectedListIndex < 0 {
						i.SelectedListIndex = 0
					}
				}
			}
		case "character":
			item := teamState.GetItemWithIndex(i.SelectedListIndex)
			// use
			log.Printf("selecting use, item: %v", item.Description)
			if item.CanConsume {
				// select character
				teamState.ConsumeItem(item.Name)
			} else {
				if item.CanEquip {
					// select character
					teamState.EquipItem(item.Name)
				}
				// ??
			}
			i.ActiveElement = "list"
			if i.SelectedListIndex == len(teamState.ItemHolders) {
				i.SelectedListIndex = i.SelectedListIndex - 1
				if i.SelectedListIndex < 0 {
					i.SelectedListIndex = 0
				}
			}
		}
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		switch i.ActiveElement {
		case "list":
			i.SelectedListIndex = i.SelectedListIndex - 1
			if i.SelectedListIndex < 0 {
				i.SelectedListIndex = 0
			}
		case "action":
			i.SelectedActionIndex = i.SelectedActionIndex - 1
			if i.SelectedActionIndex < 0 {
				i.SelectedActionIndex = 0
			}
		case "character":
			// todo select up down character
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		switch i.ActiveElement {
		case "list":
			if teamState.HasItems() {
				i.SelectedListIndex = i.SelectedListIndex + 1
				if i.SelectedListIndex == len(teamState.ItemHolders) {
					i.SelectedListIndex = i.SelectedListIndex - 1
					if i.SelectedListIndex < 0 {
						i.SelectedListIndex = 0
					}
				}
			} else {
				i.SelectedListIndex = 0
			}
		case "action":
			i.SelectedActionIndex = i.SelectedActionIndex + 1
			if i.SelectedActionIndex == 2 {
				i.SelectedActionIndex = i.SelectedActionIndex - 1
			}
		case "character":
			// todo select up down character
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		// change to other item list
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		// change to other item list
	}

}
