package gui

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/explore"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

type ShopUI struct {
	shop           *explore.Shop
	bg             *elem.StaticImage
	cursor         *elem.Cursor
	confirmation   *elem.Button
	playerName     *elem.Text
	playerMoney    *elem.Text
	shopName       *elem.Text
	moneyLabel     *elem.Text
	isLoaded       bool
	shopItems      []*listItem
	oldPlayerMoney int
	infoBox        *elem.InfoBox
}

const offsetX = 4
const offsetY = 4

var listPos = &elem.Pos{
	X: 16,
	Y: 48,
}
var infoPos = &elem.Pos{
	X: 168,
	Y: 66,
}
var confirmationPos = &elem.Pos{
	X: 168,
	Y: 48,
}

func NewShopUi() *ShopUI {
	s := &ShopUI{
		bg:           elem.NewStaticImage("uis/shop/shop-bg.png", 0, 0),
		cursor:       elem.NewCursor(),
		confirmation: elem.NewButton("Buy", "uis/shop/shop-confirmation-bg.png"),
		playerName:   elem.NewText(168+offsetX, 8+offsetY, "Player"),
		moneyLabel:   elem.NewText(168+offsetX, 24+offsetY, "Money"),
		playerMoney:  elem.NewText(168+32+offsetX, 24+offsetY, ""),
		shopName:     elem.NewText(16+offsetX, 8+offsetY, "Shop"),
		infoBox:      elem.NewInfoBox("", "uis/shop/shop-information-bg.png"),
	}
	return s
}

func (s *ShopUI) Draw(screen *ebiten.Image) {
	if s.shop == nil || !s.shop.IsActive {
		return
	}

	s.bg.Draw(screen)
	s.playerName.Draw(screen)
	s.playerMoney.Draw(screen)
	s.shopName.Draw(screen)
	s.moneyLabel.Draw(screen)

	for _, item := range s.shopItems {
		item.Draw(screen)
	}

	s.infoBox.Draw(screen)
	s.confirmation.Draw(screen)
	s.cursor.Draw(screen)
}

func (s *ShopUI) Update(delta int64, state *explore.State) {
	s.shop = state.Shop
	if !s.shop.IsActive {
		s.isLoaded = false
		return
	}
	if !s.isLoaded {
		s.isLoaded = true
		s.shopName.SetValue(s.shop.Data.MerchantName)
		s.updatePlayerMoney(state.TeamState.Money)
		s.shopItems = createListItems(s.shop.Data.Items, listPos.X+offsetX, listPos.Y+offsetY, state.TeamState.Money)
	}
	desc := s.shop.Data.Items[s.shop.SelectedListIndex].Item.Description
	s.infoBox.Update(infoPos, true, explore.GetFormattedValueMax(desc, 22))
	if s.oldPlayerMoney != state.TeamState.Money {
		s.updatePlayerMoney(state.TeamState.Money)
		s.shopItems = createListItems(s.shop.Data.Items, listPos.X+offsetX, listPos.Y+offsetY, state.TeamState.Money)
	}
	var cursorPos *elem.Pos
	switch s.shop.ActiveElement {
	case "list":
		cursorPos = &elem.Pos{
			X: listPos.X - 14,
			Y: listPos.Y + (16.0 * s.shop.SelectedListIndex),
		}
	case "confirmation":
		cursorPos = &elem.Pos{
			X: confirmationPos.X - 12,
			Y: confirmationPos.Y + 2,
		}
	}
	s.cursor.Update(delta, cursorPos)
	confirmationDisable := s.shop.ActiveElement != "confirmation"
	s.confirmation.Update(delta, confirmationPos, confirmationDisable, true)
}

func (s *ShopUI) updatePlayerMoney(money int) {
	playerMoneyString := fmt.Sprintf("%dg", money)
	playerMoneyWidth := text.BoundString(elem.StandardFont, playerMoneyString).Size().X / common.ScaleF
	s.playerMoney.SetValue(playerMoneyString)
	pos := elem.Pos{X: 168 + 64 + offsetX + 8 + 32 - playerMoneyWidth, Y: 24 + offsetY}
	s.playerMoney.SetPosition(pos)
	s.oldPlayerMoney = money
}

type listItem struct {
	item *explore.ShopItem
	name *elem.Text
	cost *elem.Text
}

func (l *listItem) Draw(screen *ebiten.Image) {
	l.name.Draw(screen)
	l.cost.Draw(screen)
}

func createListItems(items []*explore.ShopItem, x int, y int, playerMoney int) []*listItem {
	var listItems []*listItem
	var offset = 0
	for _, item := range items {
		cost := fmt.Sprintf("%dg", item.Cost)
		costWidth := text.BoundString(elem.StandardFont, cost).Size().X / common.ScaleF

		li := &listItem{
			item: item,
			name: elem.NewText(x, y+offset, item.Item.Name),
			cost: elem.NewText(x+96+32+offsetX+offsetX-costWidth, y+offset, cost),
		}
		if item.Cost > playerMoney {
			li.name.SetColor(elem.GreyTextColor)
			li.cost.SetColor(elem.GreyTextColor)
		}
		listItems = append(listItems, li)
		offset = offset + 16
	}
	return listItems
}
