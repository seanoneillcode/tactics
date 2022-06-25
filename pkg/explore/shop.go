package explore

import (
	"github.com/seanoneillcode/go-tactics/pkg/input"
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
	if input.IsCancelPressed() {
		switch s.ActiveElement {
		case "list":
			s.IsActive = false // close shop
			state.Player.Activate()
		case "confirmation":
			s.ActiveElement = "list"
		}
		return
	}
	if input.IsEnterPressed() {
		switch s.ActiveElement {
		case "list":
			shopItem := s.Data.Items[s.SelectedListIndex]
			if shopItem.Cost > state.TeamState.Money {
				return
			}
			s.ActiveElement = "confirmation"
		case "confirmation":
			shopItem := s.Data.Items[s.SelectedListIndex]
			if shopItem.Cost > state.TeamState.Money {
				return
			}
			state.TeamState.BuyItem(shopItem.Item, shopItem.Cost)
			s.ActiveElement = "list"
		}
		return
	}

	if input.IsLeftJustPressed() {
		switch s.ActiveElement {
		case "confirmation":
			s.ActiveElement = "list"
		}
	}
	if input.IsRightJustPressed() {
		switch s.ActiveElement {
		case "list":
			s.ActiveElement = "confirmation"
		}
	}
	if input.IsUpJustPressed() {
		switch s.ActiveElement {
		case "list":
			s.SelectedListIndex = s.SelectedListIndex - 1
			if s.SelectedListIndex < 0 {
				s.SelectedListIndex = 0
			}
		}
	}
	if input.IsDownJustPressed() {
		switch s.ActiveElement {
		case "list":
			s.SelectedListIndex = s.SelectedListIndex + 1
			if s.SelectedListIndex == len(s.Data.Items) {
				s.SelectedListIndex = s.SelectedListIndex - 1
			}
		}
	}
}
