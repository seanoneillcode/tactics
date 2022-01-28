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
	confirmationInfo       *Text
	informationDescription *Text
	isLoaded               bool
	shopItems              []*listItem
	cursorOffset           float64
	cursorTimer            int64
}

const offsetX = 4
const offsetY = 4

const confirmationX = 274
const confirmationY = 64

const informationX = 136
const informationY = 72

func NewShopUi() *ShopUi {
	s := &ShopUi{
		bgImage:                core.LoadImage("shop-bg.png"),
		shopCursorImage:        core.LoadImage("shop-cursor.png"),
		shopConfirmationImage:  core.LoadImage("shop-confirmation-bg.png"),
		shopInformationImage:   core.LoadImage("shop-information-bg.png"),
		playerName:             NewText(16+offsetX, 32+offsetY),
		playerMoney:            NewText(64+offsetX, 64+offsetY),
		moneyLabel:             NewText(16+offsetX, 64+offsetY),
		shopName:               NewText(128+offsetX, 32+offsetY),
		informationDescription: NewText(informationX+2+offsetX, informationY+offsetY),
		confirmationBuy:        NewText(confirmationX+2+offsetX, confirmationY+offsetY),
		confirmationInfo:       NewText(confirmationX+6+32+offsetX, confirmationY+offsetY),
	}
	s.playerName.SetValue("Player")
	s.moneyLabel.SetValue("Money")
	s.confirmationBuy.SetValue("Buy")
	s.confirmationInfo.SetValue("Info")
	return s
}

func (s *ShopUi) Draw(screen *ebiten.Image) {
	if s.shop != nil && s.shop.IsActive {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(common.Scale, common.Scale)

		screen.DrawImage(s.bgImage, op)

		s.playerName.Draw(screen)
		s.playerMoney.Draw(screen)
		s.shopName.Draw(screen)
		s.moneyLabel.Draw(screen)

		for _, item := range s.shopItems {
			item.Draw(screen)
		}

		switch s.shop.ActiveElement {
		case "list":
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(128-14+s.cursorOffset, 64.0+(float64)(16.0*s.shop.SelectedListIndex))
			op.GeoM.Scale(common.Scale, common.Scale)
			screen.DrawImage(s.shopCursorImage, op)
		case "confirmation":
			cy := confirmationY + (float64)(16.0*s.shop.SelectedListIndex)

			op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(confirmationX, cy)
			op.GeoM.Scale(common.Scale, common.Scale)
			screen.DrawImage(s.shopConfirmationImage, op)

			s.confirmationBuy.Draw(screen)
			s.confirmationInfo.Draw(screen)

			cx := float64(confirmationX - 12)
			if s.shop.SelectedConfirmationIndex == 1 {
				cx = cx + 34
			}
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(cx+s.cursorOffset, cy+2)
			op.GeoM.Scale(common.Scale, common.Scale)
			screen.DrawImage(s.shopCursorImage, op)

		case "information":
			cy := confirmationY + (float64)(16.0*s.shop.SelectedListIndex)

			op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(confirmationX, cy)
			op.GeoM.Scale(common.Scale, common.Scale)
			screen.DrawImage(s.shopConfirmationImage, op)

			s.confirmationBuy.Draw(screen)
			s.confirmationInfo.Draw(screen)

			op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(informationX, informationY)
			op.GeoM.Scale(common.Scale, common.Scale)
			screen.DrawImage(s.shopInformationImage, op)

			s.informationDescription.Draw(screen)

			op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(informationX-12+s.cursorOffset, informationY+32)
			op.GeoM.Scale(common.Scale, common.Scale)
			screen.DrawImage(s.shopCursorImage, op)
		}

	}
}

func createListItems(items []*core.ShopItem, x int, y int) []*listItem {
	var listItems []*listItem
	var offset = 0
	for _, item := range items {
		cost := fmt.Sprintf("%dg", item.Cost)
		costWidth := text.BoundString(standardFont, cost).Size().X / common.ScaleF
		listItems = append(listItems, &listItem{
			item: item,
			name: &Text{
				value: item.Item.Name,
				x:     x,
				y:     y + offset,
			},
			cost: &Text{
				value: cost,
				x:     x + 96 + 32 + offsetX + offsetX - costWidth,
				y:     y + offset,
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
		playerMoneyString := fmt.Sprintf("%dg", state.Player.CharacterState.Money)
		playerMoneyWidth := text.BoundString(standardFont, playerMoneyString).Size().X / common.ScaleF
		s.playerMoney.SetValue(playerMoneyString)
		s.playerMoney.x = 64 + offsetX + 8 + 32 - playerMoneyWidth
		s.shopItems = createListItems(s.shop.Data.Items, 128+offsetX, 64+offsetY)
	}
	if !s.shop.IsActive && s.isLoaded {
		s.isLoaded = false
	}
	if s.shop.IsActive {
		if s.shop.ActiveElement == "information" {
			desc := s.shop.Data.Items[s.shop.SelectedListIndex].Item.Description
			s.informationDescription.SetValue(core.GetFormattedValueMax(desc, 22))
		}
		if s.shop.ActiveElement == "confirmation" {
			s.confirmationBuy.y = confirmationY + offsetY + (s.shop.SelectedListIndex * 16.0)
			s.confirmationInfo.y = confirmationY + offsetY + (s.shop.SelectedListIndex * 16.0)
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

type listItem struct {
	item *core.ShopItem
	name *Text
	cost *Text
}

func (l *listItem) Draw(screen *ebiten.Image) {
	l.name.Draw(screen)
	l.cost.Draw(screen)
}
