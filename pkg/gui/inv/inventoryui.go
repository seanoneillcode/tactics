package inv

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
	"log"
)

const offsetX = 4
const offsetY = 4

var invInfoPos = &elem.Pos{
	X: 200,
	Y: 100,
}

var effectPos = &elem.Pos{
	X: 200,
	Y: 32,
}

type ctx string

const (
	listCtx      ctx = "list"
	actionCtx    ctx = "action"
	characterCtx ctx = "character"
)

type InventoryUi struct {
	cursor              *elem.Cursor
	actionBox           *ActionBox
	bg                  *elem.StaticImage
	invItemList         *InvItemList
	teamState           *core.TeamState
	justOpened          bool
	activeCtx           ctx // list, action, character
	selectedListIndex   int
	selectedActionIndex int
	selectedCharacter   string
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

func NewInventoryUi() *InventoryUi {
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
		activeCtx:         listCtx,
		selectedCharacter: "default",
	}
	return i
}

func (i *InventoryUi) Draw(screen *ebiten.Image) {
	if !i.IsActive {
		return
	}

	i.bg.Draw(screen)
	i.invItemList.Draw(screen)
	if i.activeCtx == actionCtx {
		i.actionBox.Draw(screen)
	}
	if i.activeCtx == characterCtx {
		i.actionBox.Draw(screen)
		i.characterEffect.Draw(screen)
	}
	i.infoBox.Draw(screen)
	i.cursor.Draw(screen)
}

func (i *InventoryUi) Update(delta int64, state *core.State) {
	if !state.UI.IsInventoryActive() {
		i.IsActive = false
		i.justOpened = true
		return
	}
	i.IsActive = true
	if i.justOpened {
		i.justOpened = false
		return
	}
	i.handleInput(delta, state)

	var drawInfoBox bool
	var formattedItemDescription string
	var item *core.Item
	// figure out cursor, actionBox positions
	switch i.activeCtx {
	case listCtx:
		i.cursorPos.X = i.listPos.X - 14
		i.cursorPos.Y = i.listPos.Y + (16.0 * i.selectedListIndex)
	case actionCtx:
		i.actionPos.X = i.listPos.X + 2
		i.actionPos.Y = i.listPos.Y + 11 + (16.0 * i.selectedListIndex)
		i.cursorPos.X = i.actionPos.X - 9
		i.cursorPos.Y = i.actionPos.Y + 5 + (16.0 * i.selectedActionIndex)
		if state.Player.TeamState.HasItems() {
			item = state.Player.TeamState.GetItemWithIndex(i.selectedListIndex)
			drawInfoBox = true
			formattedItemDescription = core.GetFormattedValueMax(item.Description, 22)
		}
	case characterCtx:
		i.cursorPos.X = effectPos.X - 12
		i.cursorPos.Y = effectPos.Y + 16
		if state.Player.TeamState.HasItems() {
			item = state.Player.TeamState.GetItemWithIndex(i.selectedListIndex)
			drawInfoBox = true
			formattedItemDescription = core.GetFormattedValueMax(item.Description, 22)
		}
	}

	var currentItem *core.Item
	if state.Player.TeamState.HasItems() {
		currentItem = state.Player.TeamState.GetItemWithIndex(i.selectedListIndex)
	}

	i.cursor.Update(delta, i.cursorPos)
	i.actionBox.Update(delta, i.actionPos, currentItem)
	i.invItemList.Update(delta, state.Player.TeamState)
	i.infoBox.Update(invInfoPos, drawInfoBox, formattedItemDescription)
	i.characterEffect.Update(effectPos, true, i.selectedCharacter, i.charImages[i.selectedCharacter], item, state.Player.TeamState.Characters[0])
}

func (i *InventoryUi) handleInput(delta int64, state *core.State) {
	teamState := state.Player.TeamState
	if teamState == nil {
		log.Fatal("inventory opened with no team!")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		switch i.activeCtx {
		case listCtx:
			state.UI.Close()
			state.Player.Activate()
		case actionCtx:
			i.activeCtx = listCtx
		case characterCtx:
			i.activeCtx = actionCtx
		}
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.UI.Close()
		state.Player.Activate()
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch i.activeCtx {
		case listCtx:
			if teamState.HasItems() {
				i.activeCtx = actionCtx
				i.selectedActionIndex = 0
			}
		case actionCtx:
			item := teamState.GetItemWithIndex(i.selectedListIndex)
			if i.selectedActionIndex == 0 {
				i.activeCtx = characterCtx
			} else {
				// drop
				teamState.RemoveItem(item.Name)
				i.activeCtx = listCtx
				if i.selectedListIndex == len(teamState.ItemHolders) {
					i.selectedListIndex = i.selectedListIndex - 1
					if i.selectedListIndex < 0 {
						i.selectedListIndex = 0
					}
				}
			}
		case characterCtx:
			item := teamState.GetItemWithIndex(i.selectedListIndex)
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
			i.activeCtx = listCtx
			if i.selectedListIndex == len(teamState.ItemHolders) {
				i.selectedListIndex = i.selectedListIndex - 1
				if i.selectedListIndex < 0 {
					i.selectedListIndex = 0
				}
			}
		}
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		switch i.activeCtx {
		case listCtx:
			i.selectedListIndex = i.selectedListIndex - 1
			if i.selectedListIndex < 0 {
				i.selectedListIndex = 0
			}
		case actionCtx:
			i.selectedActionIndex = i.selectedActionIndex - 1
			if i.selectedActionIndex < 0 {
				i.selectedActionIndex = 0
			}
		case characterCtx:
			// todo select up down character
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		switch i.activeCtx {
		case listCtx:
			if teamState.HasItems() {
				i.selectedListIndex = i.selectedListIndex + 1
				if i.selectedListIndex == len(teamState.ItemHolders) {
					i.selectedListIndex = i.selectedListIndex - 1
					if i.selectedListIndex < 0 {
						i.selectedListIndex = 0
					}
				}
			} else {
				i.selectedListIndex = 0
			}
		case actionCtx:
			i.selectedActionIndex = i.selectedActionIndex + 1
			if i.selectedActionIndex == 2 {
				i.selectedActionIndex = i.selectedActionIndex - 1
			}
		case characterCtx:
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
