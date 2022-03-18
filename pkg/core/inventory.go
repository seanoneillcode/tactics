package core

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Inventory manages team state, items money characters etc
type Inventory struct {
	IsActive            bool
	TeamState           *TeamState
	justOpened          bool
	ActiveElement       string // list, action, character
	SelectedListIndex   int
	SelectedActionIndex int
	SelectedCharacter   string
}

func NewInventory() *Inventory {
	return &Inventory{
		ActiveElement:     "list",
		SelectedListIndex: 0,
		SelectedCharacter: "default",
	}
}

func (i *Inventory) Open(state *TeamState) {
	i.IsActive = true
	i.TeamState = state
}

func (i *Inventory) Update(delta int64, state *State) {
	if !i.IsActive {
		i.justOpened = true
		return
	}
	if i.justOpened {
		i.justOpened = false
		return
	}
	if i.TeamState == nil {
		log.Fatal("inventory opened with no team!")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		switch i.ActiveElement {
		case "list":
			i.IsActive = false // close
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
			if i.HasItems() {
				i.ActiveElement = "action"
				i.SelectedActionIndex = 0
			}
		case "action":
			item := i.TeamState.GetItemWithIndex(i.SelectedListIndex)
			if i.SelectedActionIndex == 0 {
				i.ActiveElement = "character"
			} else {
				// drop
				state.Player.TeamState.RemoveItem(item.Name)
				i.ActiveElement = "list"
				if i.SelectedListIndex == len(i.TeamState.Items) {
					i.SelectedListIndex = i.SelectedListIndex - 1
				}
			}
		case "character":
			item := i.TeamState.GetItemWithIndex(i.SelectedListIndex)
			// use
			log.Printf("selecting use, item: %v", item.Description)
			if item.CanConsume {
				// select character
				state.Player.TeamState.ConsumeItem(item.Name)
			} else {
				if item.CanEquip {
					// select character
					state.Player.TeamState.EquipItem(item.Name)
				}
				// ??
			}
			i.ActiveElement = "list"
			if i.SelectedListIndex == len(i.TeamState.Items) {
				i.SelectedListIndex = i.SelectedListIndex - 1
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
			if i.HasItems() {
				i.SelectedListIndex = i.SelectedListIndex + 1
				if i.SelectedListIndex == len(i.TeamState.Items) {
					i.SelectedListIndex = i.SelectedListIndex - 1
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

func (i *Inventory) HasItems() bool {
	return len(i.TeamState.Items) > 0
}
