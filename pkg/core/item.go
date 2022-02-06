package core

import (
	"log"
)

type Item struct {
	Name        string
	Description string
	CanConsume  bool
	CanEquip    bool
	Effects     []StateEffect
	StatEffects []Stats
}

func NewItem(name string) *Item {
	switch name {
	case HerbItemName:
		return &Item{
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
			Name:        "Potion",
			Description: "A distilled herb that significantly heals wounds.",
			CanConsume:  true,
			CanEquip:    false,
			Effects: []StateEffect{
				&healthEffect{amount: 15},
			},
		}
	case PaddedArmorItemName:
		return &Item{
			Name:        "Padded Armour",
			Description: "A tunic made from several layer of cotton that reduces damage.",
			CanConsume:  false,
			CanEquip:    true,
			StatEffects: []Stats{
				{Defence: 1},
			},
		}
	case SteelArmorItemName:
		return &Item{
			Name:        "Steel Plated Armour",
			Description: "A heavy piece of chest armour made from overlapping steel plates.",
			CanConsume:  false,
			CanEquip:    true,
			StatEffects: []Stats{
				{Defence: 3},
				{Speed: -1},
			},
		}
	}
	log.Fatalf("unknown item: %s", name)
	return nil
}

const (
	HerbItemName        = "herb"
	PotionItemName      = "potion"
	PaddedArmorItemName = "padded-armor"
	SteelArmorItemName  = "steel-armor"
)
