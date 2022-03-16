package gui

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

type ShopUi struct {
	shop                   *core.Shop
	bg                     *elem.StaticImage
	cursor                 *elem.Cursor
	confirmation           *elem.Button
	shopInformationImage   *ebiten.Image
	playerName             *elem.Text
	playerMoney            *elem.Text
	shopName               *elem.Text
	moneyLabel             *elem.Text
	informationDescription *elem.Text
	isLoaded               bool
	shopItems              []*listItem
	oldPlayerMoney         int
}

const offsetX = 4
const offsetY = 4

const listX = 48
const listY = 64

const confirmationX = 208
const confirmationY = 64
const informationX = 208
const informationY = 80

func NewShopUi() *ShopUi {
	s := &ShopUi{
		bg:                     elem.NewStaticImage("shop-bg.png", 0, 0),
		cursor:                 elem.NewCursor(),
		confirmation:           elem.NewButton("Buy", "shop-confirmation-bg.png"),
		shopInformationImage:   common.LoadImage("shop-information-bg.png"),
		playerName:             elem.NewText(240+offsetX, 16+offsetY, "Player"),
		moneyLabel:             elem.NewText(240+offsetX, 32+offsetY, "Money"),
		playerMoney:            elem.NewText(240+64+offsetX, 32+offsetY, ""),
		shopName:               elem.NewText(48+offsetX, 32+offsetY, ""),
		informationDescription: elem.NewText(informationX+2+offsetX, informationY+offsetY, ""),
	}
	return s
}

func (s *ShopUi) Draw(screen *ebiten.Image) {
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

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(informationX, informationY)
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(s.shopInformationImage, op)

	s.informationDescription.Draw(screen)
	s.confirmation.Draw(screen)
	s.cursor.Draw(screen)
}

func (s *ShopUi) Update(delta int64, state *core.State) {
	s.shop = state.Shop
	if s.shop.IsActive && !s.isLoaded {
		s.isLoaded = true
		s.shopName.SetValue(s.shop.Data.MerchantName)
		s.updatePlayerMoney(state.Player.TeamState.Money)
		s.shopItems = createListItems(s.shop.Data.Items, listX+offsetX, listY+offsetY, state.Player.TeamState.Money)
	}
	if !s.shop.IsActive && s.isLoaded {
		s.isLoaded = false
	}
	if s.shop.IsActive {
		desc := s.shop.Data.Items[s.shop.SelectedListIndex].Item.Description
		s.informationDescription.SetValue(core.GetFormattedValueMax(desc, 22))
		if s.oldPlayerMoney != state.Player.TeamState.Money {
			s.updatePlayerMoney(state.Player.TeamState.Money)
			s.shopItems = createListItems(s.shop.Data.Items, listX+offsetX, listY+offsetY, state.Player.TeamState.Money)
		}
	}
	var cursorPos *elem.Pos
	switch s.shop.ActiveElement {
	case "list":
		cursorPos = &elem.Pos{
			X: listX - 14,
			Y: listY + (16.0 * s.shop.SelectedListIndex),
		}
	case "confirmation":
		cursorPos = &elem.Pos{
			X: confirmationX - 12,
			Y: confirmationY + 2,
		}
	}
	s.cursor.Update(delta, cursorPos)
	s.confirmation.Update(delta, &elem.Pos{X: confirmationX, Y: confirmationY}, s.shop.ActiveElement == "confirmation")
}

func (s *ShopUi) updatePlayerMoney(money int) {
	playerMoneyString := fmt.Sprintf("%dg", money)
	playerMoneyWidth := text.BoundString(elem.StandardFont, playerMoneyString).Size().X / common.ScaleF
	s.playerMoney.SetValue(playerMoneyString)
	pos := elem.Pos{X: 240 + 64 + offsetX + 8 + 32 - playerMoneyWidth, Y: 32 + offsetY}
	s.playerMoney.SetPosition(pos)
	s.oldPlayerMoney = money
}

type listItem struct {
	item *core.ShopItem
	name *elem.Text
	cost *elem.Text
}

func (l *listItem) Draw(screen *ebiten.Image) {
	l.name.Draw(screen)
	l.cost.Draw(screen)
}

func createListItems(items []*core.ShopItem, x int, y int, playerMoney int) []*listItem {
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
