package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
)

type Shop struct {
	Data                      *ShopData
	IsActive                  bool
	ActiveElement             string
	SelectedListIndex         int
	SelectedConfirmationIndex int
	justOpened                bool
}

func NewShop() *Shop {
	return &Shop{
		ActiveElement: "list",
	}
}

func (s *Shop) Open(data *ShopData) {
	s.IsActive = true
	s.Data = data
}

func (s *Shop) Update(delta int64, state *State) {
	if !s.IsActive {
		s.justOpened = true
		return
	}
	if s.justOpened {
		s.justOpened = false
		return
	}
	if s.Data == nil {
		log.Fatalf("opened a shop with no Data")
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		switch s.ActiveElement {
		case "list":
			s.IsActive = false // close shop
			state.Player.Activate()
		case "confirmation":
			s.ActiveElement = "list"
		case "information":
			s.ActiveElement = "confirmation"
		}
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		log.Printf("shop is open, pressed accept")
		switch s.ActiveElement {
		case "list":
			s.ActiveElement = "confirmation"
		case "confirmation":
			if s.SelectedConfirmationIndex == 0 {
				s.ActiveElement = "list"
				// buy the item
			} else {
				s.ActiveElement = "information"
			}
		case "information":
			s.ActiveElement = "confirmation"
		}
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		switch s.ActiveElement {
		case "confirmation":
			switch s.SelectedConfirmationIndex {
			case 0:
				s.ActiveElement = "list"
			case 1:
				s.SelectedConfirmationIndex = 0
			}
		case "information":
			s.ActiveElement = "confirmation"
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		switch s.ActiveElement {
		case "list":
			s.ActiveElement = "confirmation"
		case "confirmation":
			switch s.SelectedConfirmationIndex {
			case 0:
				s.SelectedConfirmationIndex = 1
			case 1:
				s.ActiveElement = "information"
			}
		case "information":
			s.ActiveElement = "confirmation"
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		switch s.ActiveElement {
		case "list":
			s.SelectedListIndex = s.SelectedListIndex - 1
			if s.SelectedListIndex < 0 {
				s.SelectedListIndex = 0
			}
		case "confirmation":
			s.SelectedListIndex = s.SelectedListIndex - 1
			if s.SelectedListIndex < 0 {
				s.SelectedListIndex = 0
			}
		}
		log.Printf("selected index: %v", s.SelectedListIndex)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		switch s.ActiveElement {
		case "list":
			s.SelectedListIndex = s.SelectedListIndex + 1
			if s.SelectedListIndex == len(s.Data.Items) {
				s.SelectedListIndex = s.SelectedListIndex - 1
			}
		case "confirmation":
			s.SelectedListIndex = s.SelectedListIndex + 1
			if s.SelectedListIndex == len(s.Data.Items) {
				s.SelectedListIndex = s.SelectedListIndex - 1
			}
		}
		log.Printf("selected index: %v", s.SelectedListIndex)
	}
}
