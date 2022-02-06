package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
)

// Inventory manages team state, items money characters etc
type Inventory struct {
	IsActive            bool
	TeamState           *TeamState
	justOpened          bool
	ActiveElement       string // list, action
	SelectedListIndex   int
	SelectedActionIndex int
}

func NewInventory() *Inventory {
	return &Inventory{
		ActiveElement:     "list",
		SelectedListIndex: 0,
	}
}

func (i *Inventory) Open(state *TeamState) {
	log.Println("opening inventory")
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
			log.Println("closing inventory")
		case "action":
			i.ActiveElement = "list"
			log.Println("selecting list")
		}
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch i.ActiveElement {
		case "list":
			if i.HasItems() {
				i.ActiveElement = "action"
			}
		case "action":
			if i.SelectedActionIndex == 0 {
				// use
				item := i.TeamState.Items[i.SelectedListIndex]
				log.Printf("selecting use, item: %v", item.Description)
				if item.CanConsume {
					// select character
					state.Player.TeamState.ConsumeItem(i.SelectedListIndex)
				} else {
					if item.CanEquip {
						// select character
						state.Player.TeamState.EquipItem(i.SelectedListIndex)
					}
					// ??
				}
			} else {
				log.Println("selecting drop")
				// drop
				state.Player.TeamState.RemoveItem(i.SelectedListIndex)
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
		}
		log.Printf("selected index: %v", i.SelectedListIndex)
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
		}
		log.Printf("selected index: %v", i.SelectedListIndex)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		switch i.ActiveElement {
		case "action":
			if i.SelectedActionIndex == 0 {
				i.ActiveElement = "list"
			} else {
				i.SelectedActionIndex = 0
			}
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		switch i.ActiveElement {
		case "list":
			i.ActiveElement = "action"
		case "action":
			if i.SelectedActionIndex == 0 {
				i.SelectedActionIndex = 1
			}
		}
	}

}

func (i *Inventory) HasItems() bool {
	return len(i.TeamState.Items) > 0
}
