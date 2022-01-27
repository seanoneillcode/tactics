package core

import (
	"log"
)

type Item struct {
	Name        string
	Description string
	CanConsume  bool
	CanEquip    bool
	Effects     []*Effect
}

func NewItem(name string) *Item {
	switch name {
	case HerbItemName:
		return &Item{
			Name:        "Herb",
			Description: "A plant that soothes and heals wounds.",
			CanConsume:  true,
			CanEquip:    false,
			Effects: []*Effect{
				{
					Property: "health",
					Value:    4,
				},
			},
		}
	case PotionItemName:
		return &Item{
			Name:        "Potion",
			Description: "A distilled herb that significantly heals wounds.",
			CanConsume:  true,
			CanEquip:    false,
			Effects: []*Effect{
				{
					Property: "health",
					Value:    15,
				},
			},
		}
	case PaddedArmorItemName:
		return &Item{
			Name:        "Padded Armour",
			Description: "A tunic made from several layer of cotton that reduces damage.",
			CanConsume:  false,
			CanEquip:    true,
			Effects: []*Effect{
				{
					Property: "defense",
					Value:    2,
				},
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
)
