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

type InventoryUI struct {
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

func NewInventoryUi() *InventoryUI {
	i := &InventoryUI{
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

func (r *InventoryUI) Draw(screen *ebiten.Image) {
	if !r.IsActive {
		return
	}

	r.bg.Draw(screen)
	r.invItemList.Draw(screen)
	if r.activeCtx == actionCtx {
		r.actionBox.Draw(screen)
	}
	if r.activeCtx == characterCtx {
		r.actionBox.Draw(screen)
		r.characterEffect.Draw(screen)
	}
	r.infoBox.Draw(screen)
	r.cursor.Draw(screen)
}

func (r *InventoryUI) Update(delta int64, state *core.State) {
	if !state.UI.IsInventoryActive() {
		r.IsActive = false
		r.justOpened = true
		return
	}
	r.IsActive = true
	if r.justOpened {
		r.justOpened = false
		return
	}
	r.handleInput(delta, state)

	var drawInfoBox bool
	var formattedItemDescription string
	var item *core.Item
	// figure out cursor, actionBox positions
	switch r.activeCtx {
	case listCtx:
		r.cursorPos.X = r.listPos.X - 14
		r.cursorPos.Y = r.listPos.Y + (16.0 * r.selectedListIndex)
	case actionCtx:
		r.actionPos.X = r.listPos.X + 2
		r.actionPos.Y = r.listPos.Y + 11 + (16.0 * r.selectedListIndex)
		r.cursorPos.X = r.actionPos.X - 9
		r.cursorPos.Y = r.actionPos.Y + 5 + (16.0 * r.selectedActionIndex)
		if state.Player.TeamState.HasItems() {
			item = state.Player.TeamState.GetItemWithIndex(r.selectedListIndex)
			drawInfoBox = true
			formattedItemDescription = core.GetFormattedValueMax(item.Description, 22)
		}
	case characterCtx:
		r.cursorPos.X = effectPos.X - 12
		r.cursorPos.Y = effectPos.Y + 16
		if state.Player.TeamState.HasItems() {
			item = state.Player.TeamState.GetItemWithIndex(r.selectedListIndex)
			drawInfoBox = true
			formattedItemDescription = core.GetFormattedValueMax(item.Description, 22)
		}
	}

	var currentItem *core.Item
	if state.Player.TeamState.HasItems() {
		currentItem = state.Player.TeamState.GetItemWithIndex(r.selectedListIndex)
	}

	r.cursor.Update(delta, r.cursorPos)
	r.actionBox.Update(delta, r.actionPos, currentItem)
	r.invItemList.Update(delta, state.Player.TeamState)
	r.infoBox.Update(invInfoPos, drawInfoBox, formattedItemDescription)
	r.characterEffect.Update(effectPos, true, r.selectedCharacter, r.charImages[r.selectedCharacter], item, state.Player.TeamState.Characters[0])
}

func (r *InventoryUI) handleInput(delta int64, state *core.State) {
	teamState := state.Player.TeamState
	if teamState == nil {
		log.Fatal("inventory opened with no team!")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		switch r.activeCtx {
		case listCtx:
			state.UI.Open(core.MenuUI)
		case actionCtx:
			r.activeCtx = listCtx
		case characterCtx:
			r.activeCtx = actionCtx
		}
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.UI.Close()
		state.Player.Activate()
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch r.activeCtx {
		case listCtx:
			if teamState.HasItems() {
				r.activeCtx = actionCtx
				r.selectedActionIndex = 0
			}
		case actionCtx:
			item := teamState.GetItemWithIndex(r.selectedListIndex)
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
			item := teamState.GetItemWithIndex(r.selectedListIndex)
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
			r.activeCtx = listCtx
			if r.selectedListIndex == len(teamState.ItemHolders) {
				r.selectedListIndex = r.selectedListIndex - 1
				if r.selectedListIndex < 0 {
					r.selectedListIndex = 0
				}
			}
		}
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
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
			// todo select up down character
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
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
