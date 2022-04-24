package core

import (
	"log"
)

type Item struct {
	Type        string
	Name        string
	Description string
	CanConsume  bool
	CanEquip    bool
	Effects     []StateEffect
	StatEffects []Stats
	EquipSlot   string
	ImagePath   string
}

func NewItem(name string) *Item {
	switch name {
	case MouldyBreadItemName:
		return &Item{
			Type:        MouldyBreadItemName,
			Name:        "Mouldy Bread Roll",
			Description: "A bread roll with suspicious green edges.",
			CanConsume:  true,
			CanEquip:    false,
			Effects: []StateEffect{
				&healthEffect{amount: -3},
			},
		}
	case BreadItemName:
		return &Item{
			Type:        BreadItemName,
			Name:        "Bread Roll",
			Description: "A wholesome bread roll. A handy snack any time, any place.",
			CanConsume:  true,
			CanEquip:    false,
			Effects: []StateEffect{
				&healthEffect{amount: 1},
			},
		}
	case HerbItemName:
		return &Item{
			Type:        HerbItemName,
			Name:        "Herb",
			Description: "A plant that soothes and heals wounds.",
			CanConsume:  true,
			CanEquip:    false,
			Effects: []StateEffect{
				&healthEffect{amount: 4},
			},
		}
	case PotionItemName:
		return &Item{
			Type:        PotionItemName,
			Name:        "Potion",
			Description: "A distilled herb that significantly heals wounds.",
			CanConsume:  true,
			CanEquip:    false,
			Effects: []StateEffect{
				&healthEffect{amount: 15},
			},
			ImagePath: "item/potion.png",
		}
	case EtherItemName:
		return &Item{
			Type:        EtherItemName,
			Name:        "Ether",
			Description: "A distilled crystal that restores magic energy.",
			CanConsume:  true,
			CanEquip:    false,
			Effects: []StateEffect{
				&magicEffect{amount: 15},
			},
			ImagePath: "item/ether.png",
		}
	case PaddedArmorItemName:
		return &Item{
			Type:        PaddedArmorItemName,
			Name:        "Padded Armour",
			Description: "A tunic made from several layer of cotton that reduces damage.",
			CanConsume:  false,
			CanEquip:    true,
			StatEffects: []Stats{
				{Defence: 1},
			},
			ImagePath: "item/padded-armour.png",
		}
	case SteelArmorItemName:
		return &Item{
			Type:        SteelArmorItemName,
			Name:        "Plated Armour",
			Description: "A heavy piece of chest armour made from overlapping steel plates.",
			CanConsume:  false,
			CanEquip:    true,
			StatEffects: []Stats{
				{Defence: 3},
				{Speed: -1},
			},
			ImagePath: "item/steel-armour.png",
		}
	}
	log.Fatalf("unknown item: %s", name)
	return nil
}

const (
	BreadItemName       = "bread"
	MouldyBreadItemName = "mouldy-bread"
	HerbItemName        = "herb"
	PotionItemName      = "potion"
	EtherItemName       = "ether"
	PaddedArmorItemName = "padded-armor"
	SteelArmorItemName  = "steel-armor"
)
