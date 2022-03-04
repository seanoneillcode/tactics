package core

import "github.com/seanoneillcode/go-tactics/pkg/common"

type ShopData struct {
	Name         string
	MerchantName string
	pos          *common.Vector
	Items        []*ShopItem
}

func NewShopData(name string, x float64, y float64) *ShopData {
	return &ShopData{
		Name: name,
		pos: &common.Vector{
			X: x,
			Y: y,
		},
		Items: shopData[name],
	}
}

func (s *ShopData) GetPosition() *common.Vector {
	return s.pos
}

type ShopItem struct {
	Item *Item
	Cost int
}

var shopData = map[string][]*ShopItem{
	"shop-home": {
		&ShopItem{
			Item: NewItem(BreadItemName),
			Cost: 1,
		},
		&ShopItem{
			Item: NewItem(HerbItemName),
			Cost: 2,
		},
		&ShopItem{
			Item: NewItem(PotionItemName),
			Cost: 5,
		},
		&ShopItem{
			Item: NewItem(PaddedArmorItemName),
			Cost: 20,
		},
		&ShopItem{
			Item: NewItem(SteelArmorItemName),
			Cost: 100,
		},
	},
}
