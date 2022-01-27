package core

import "github.com/seanoneillcode/go-tactics/pkg/common"

type ShopData struct {
	name  string
	pos   *common.VectorF
	items []*shopItem
}

func NewShopData(name string, x float64, y float64) *ShopData {
	return &ShopData{
		name: name,
		pos: &common.VectorF{
			X: x,
			Y: y,
		},
		items: shopData[name],
	}
}

func (s *ShopData) GetPosition() *common.VectorF {
	return s.pos
}

type shopItem struct {
	name string
	cost int
}

var shopData = map[string][]*shopItem{
	"shop-home": {
		&shopItem{
			name: HerbItemName,
			cost: 1,
		},
		&shopItem{
			name: PotionItemName,
			cost: 5,
		},
		&shopItem{
			name: PaddedArmorItemName,
			cost: 25,
		},
	},
}
