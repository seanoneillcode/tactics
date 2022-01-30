package core

import "github.com/seanoneillcode/go-tactics/pkg/common"

type ShopData struct {
	Name         string
	MerchantName string
	pos          *common.VectorF
	Items        []*ShopItem
}

func NewShopData(name string, x float64, y float64) *ShopData {
	return &ShopData{
		Name: name,
		pos: &common.VectorF{
			X: x,
			Y: y,
		},
		Items: shopData[name],
	}
}

func (s *ShopData) GetPosition() *common.VectorF {
	return s.pos
}

type ShopItem struct {
	Item *Item
	Cost int
}

var shopData = map[string][]*ShopItem{
	"shop-home": {
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
			Cost: 999999,
		},
	},
}