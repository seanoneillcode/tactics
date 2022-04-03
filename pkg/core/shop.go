package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
)

type Shop struct {
	Data              *ShopData
	IsActive          bool
	ActiveElement     string // todo move UI state into gui package
	SelectedListIndex int    // todo move UI state into gui package
	justOpened        bool   // todo move UI state into gui package
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
		}
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch s.ActiveElement {
		case "list":
			shopItem := s.Data.Items[s.SelectedListIndex]
			if shopItem.Cost > state.Player.TeamState.Money {
				return
			}
			s.ActiveElement = "confirmation"
		case "confirmation":
			shopItem := s.Data.Items[s.SelectedListIndex]
			if shopItem.Cost > state.Player.TeamState.Money {
				return
			}
			state.Player.TeamState.BuyItem(shopItem.Item, shopItem.Cost)
			s.ActiveElement = "list"
		}
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		switch s.ActiveElement {
		case "confirmation":
			s.ActiveElement = "list"
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		switch s.ActiveElement {
		case "list":
			s.ActiveElement = "confirmation"
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		switch s.ActiveElement {
		case "list":
			s.SelectedListIndex = s.SelectedListIndex - 1
			if s.SelectedListIndex < 0 {
				s.SelectedListIndex = 0
			}
		}
		log.Printf("selected index: %v", s.SelectedListIndex)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		switch s.ActiveElement {
		case "list":
			s.SelectedListIndex = s.SelectedListIndex + 1
			if s.SelectedListIndex == len(s.Data.Items) {
				s.SelectedListIndex = s.SelectedListIndex - 1
			}
		}
		log.Printf("selected index: %v", s.SelectedListIndex)
	}
}
