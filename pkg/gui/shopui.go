package gui

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
)

type ShopUi struct {
	shop                   *core.Shop
	bgImage                *ebiten.Image
	shopCursorImage        *ebiten.Image
	shopConfirmationImage  *ebiten.Image
	shopInformationImage   *ebiten.Image
	playerName             *Text
	playerMoney            *Text
	shopName               *Text
	moneyLabel             *Text
	confirmationBuy        *Text
	informationDescription *Text
	isLoaded               bool
	shopItems              []*listItem
	cursorOffset           float64
	cursorTimer            int64
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
		bgImage:               core.LoadImage("shop-bg.png"),
		shopCursorImage:       core.LoadImage("shop-cursor.png"),
		shopConfirmationImage: core.LoadImage("shop-confirmation-bg.png"),
		shopInformationImage:  core.LoadImage("shop-information-bg.png"),
		playerName:            NewText(240+offsetX, 16+offsetY),
		moneyLabel:            NewText(240+offsetX, 32+offsetY),
		playerMoney:           NewText(240+64+offsetX, 32+offsetY),
		//shopName:               NewText(240+offsetX, 16+offsetY),
		shopName:               NewText(48+offsetX, 32+offsetY),
		informationDescription: NewText(informationX+2+offsetX, informationY+offsetY),
		confirmationBuy:        NewText(confirmationX+2+offsetX, confirmationY+offsetY),
	}
	s.playerName.SetValue("Player")
	s.moneyLabel.SetValue("Money")
	s.confirmationBuy.SetValue("Buy")
	return s
}

func (s *ShopUi) Draw(screen *ebiten.Image) {
	if s.shop == nil || !s.shop.IsActive {
		return
	}

	// background
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(s.bgImage, op)

	// set elements
	s.playerName.Draw(screen)
	s.playerMoney.Draw(screen)
	s.shopName.Draw(screen)
	s.moneyLabel.Draw(screen)

	for _, item := range s.shopItems {
		item.Draw(screen)
	}

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(informationX, informationY)
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(s.shopInformationImage, op)

	s.informationDescription.Draw(screen)
	s.confirmationBuy.Draw(screen)

	switch s.shop.ActiveElement {
	case "list":
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(listX-14+s.cursorOffset, listY+(float64)(16.0*s.shop.SelectedListIndex))
		op.GeoM.Scale(common.Scale, common.Scale)
		screen.DrawImage(s.shopCursorImage, op)
		s.confirmationBuy.SetColor(greyTextColor)

	case "confirmation":
		cy := float64(confirmationY) //+ (float64)(16.0*s.shop.SelectedListIndex)

		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(confirmationX, cy)
		op.GeoM.Scale(common.Scale, common.Scale)
		screen.DrawImage(s.shopConfirmationImage, op)

		s.confirmationBuy.SetColor(defaultTextColor)
		s.confirmationBuy.Draw(screen)

		cx := float64(confirmationX - 12)
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(cx+s.cursorOffset, cy+2)
		op.GeoM.Scale(common.Scale, common.Scale)
		screen.DrawImage(s.shopCursorImage, op)
	}

}

func createListItems(items []*core.ShopItem, x int, y int, playerMoney int) []*listItem {
	var listItems []*listItem
	var offset = 0
	for _, item := range items {
		cost := fmt.Sprintf("%dg", item.Cost)
		costWidth := text.BoundString(standardFont, cost).Size().X / common.ScaleF
		color := defaultTextColor
		if item.Cost > playerMoney {
			color = greyTextColor
		}
		listItems = append(listItems, &listItem{
			item: item,
			name: &Text{
				value: item.Item.Name,
				x:     x,
				y:     y + offset,
				color: color,
			},
			cost: &Text{
				value: cost,
				x:     x + 96 + 32 + offsetX + offsetX - costWidth,
				y:     y + offset,
				color: color,
			},
		})
		offset = offset + 16
	}
	return listItems
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
	s.cursorTimer = s.cursorTimer + delta
	if s.cursorTimer > 400 {
		s.cursorTimer = s.cursorTimer - 400
		if s.cursorOffset == 0 {
			s.cursorOffset = 2
		} else {
			s.cursorOffset = 0
		}
	}
}

func (s *ShopUi) updatePlayerMoney(money int) {
	playerMoneyString := fmt.Sprintf("%dg", money)
	playerMoneyWidth := text.BoundString(standardFont, playerMoneyString).Size().X / common.ScaleF
	s.playerMoney.SetValue(playerMoneyString)
	s.playerMoney.x = 240 + 64 + offsetX + 8 + 32 - playerMoneyWidth
	s.oldPlayerMoney = money
}

type listItem struct {
	item *core.ShopItem
	name *Text
	cost *Text
}

func (l *listItem) Draw(screen *ebiten.Image) {
	l.name.Draw(screen)
	l.cost.Draw(screen)
}
