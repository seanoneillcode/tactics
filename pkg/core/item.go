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
			EquipSlot: "armor",
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
				{Defence: 5},
				{Speed: -2},
			},
			EquipSlot: "armor",
			ImagePath: "item/steel-armour.png",
		}
	case LeatherArmorItemName:
		return &Item{
			Type:        LeatherArmorItemName,
			Name:        "Leather Armour",
			Description: "A tough chest piece made from thick leather. It smells like burnt sweat.",
			CanConsume:  false,
			CanEquip:    true,
			StatEffects: []Stats{
				{Defence: 2},
				{Speed: -1},
			},
			EquipSlot: "armor",
			ImagePath: "item/leather-armour.png",
		}
	case ChainArmorItemName:
		return &Item{
			Type:        ChainArmorItemName,
			Name:        "Chain Mail Armor",
			Description: "A light and thin shirt of chain mail.",
			CanConsume:  false,
			CanEquip:    true,
			StatEffects: []Stats{
				{Defence: 3},
				{Speed: -1},
			},
			EquipSlot: "armor",
			ImagePath: "item/chain-mail-armour.png",
		}
	case BatteredSwordItemName:
		return &Item{
			Type:        BatteredSwordItemName,
			Name:        "Battered Sword",
			Description: "Dents, scratches and signs of repair. It still cuts if swung hard enough.",
			CanConsume:  false,
			CanEquip:    true,
			StatEffects: []Stats{
				{Defence: 1},
				{AttackSkill: 2},
			},
			EquipSlot: "weapon",
			ImagePath: "item/short-sword.png",
		}
	case RapierItemName:
		return &Item{
			Type:        RapierItemName,
			Name:        "Rapier",
			Description: "Long and gentle with a sharp tip. It's too light to parry with, but that's okay if you're fast enough.",
			CanConsume:  false,
			CanEquip:    true,
			StatEffects: []Stats{
				{AttackSkill: 4},
				{Speed: 1},
			},
			EquipSlot: "weapon",
			ImagePath: "item/short-sword.png",
		}
	case SaberItemName:
		return &Item{
			Type:        SaberItemName,
			Name:        "Saber",
			Description: "Curved and sharp with a little weight. It's meant for horse back, but found everywhere.",
			CanConsume:  false,
			CanEquip:    true,
			StatEffects: []Stats{
				{AttackSkill: 4},
				{Defence: 1},
			},
			EquipSlot: "weapon",
			ImagePath: "item/short-sword.png",
		}
	case KnifeItemName:
		return &Item{
			Type:        KnifeItemName,
			Name:        "Knife",
			Description: "A hands length of sharp steel. Used in the kitchen or the pub.",
			CanConsume:  false,
			CanEquip:    true,
			StatEffects: []Stats{
				{AttackSkill: 1},
			},
			EquipSlot: "weapon",
			ImagePath: "item/short-sword.png",
		}
	case PineStaffItemName:
		return &Item{
			Type:        PineStaffItemName,
			Name:        "Green Pine Staff",
			Description: "A newly cut length of pine, fashioned into a staff. Springy and soft.",
			CanConsume:  false,
			CanEquip:    true,
			StatEffects: []Stats{
				{AttackSkill: 2},
			},
			EquipSlot: "weapon",
			ImagePath: "item/short-sword.png",
		}
	case HuntingBowItemName:
		return &Item{
			Type:        HuntingBowItemName,
			Name:        "Hunting Bow",
			Description: "A common hunting bow, found in every village.",
			CanConsume:  false,
			CanEquip:    true,
			StatEffects: []Stats{
				{AttackSkill: 3},
			},
			EquipSlot: "weapon",
			ImagePath: "item/short-sword.png",
		}
	case MagicSockItemName:
		return &Item{
			Type:        MagicSockItemName,
			Name:        "Magic Sock",
			Description: "An old sock with holes. After returning from the lost sock void, it is filled with eldritch knowledge.",
			CanConsume:  false,
			CanEquip:    true,
			StatEffects: []Stats{
				{MagicDef: 3},
			},
			EquipSlot: "accessory",
			ImagePath: "item/ring.png",
		}
	case MagicRingItemName:
		return &Item{
			Type:        MagicRingItemName,
			Name:        "Magic Ring",
			Description: "A ring of unknown materials, it glows lightly. It shifts and warps when it's in your peripheral vision.",
			CanConsume:  false,
			CanEquip:    true,
			StatEffects: []Stats{
				{MagicSkill: 3},
			},
			EquipSlot: "accessory",
			ImagePath: "item/ring.png",
		}
	case MagicHatItemName:
		return &Item{
			Type:        MagicHatItemName,
			Name:        "Magic Hat",
			Description: "A unfashionable hat, dusty and stained. It's feels like wearing an old couch on your head.",
			CanConsume:  false,
			CanEquip:    true,
			StatEffects: []Stats{
				{MaxMagic: 3},
			},
			EquipSlot: "accessory",
			ImagePath: "item/ring.png",
		}
	}
	log.Fatalf("unknown item: %s", name)
	return nil
}

const (
	// consumables

	BreadItemName       = "bread"
	MouldyBreadItemName = "mouldy-bread"
	HerbItemName        = "herb"
	PotionItemName      = "potion"
	EtherItemName       = "ether"

	// armor

	PaddedArmorItemName  = "padded-armor"
	SteelArmorItemName   = "steel-armor"
	LeatherArmorItemName = "leather-armor"
	ChainArmorItemName   = "chain-armor"

	// weapons

	WoodenPracticeSwordItemName = "wooden-sword"
	BatteredSwordItemName       = "battered-sword"
	RapierItemName              = "rapier-sword"
	ShortSwordItemName          = "short-sword"
	LongSwordItemName           = "long-sword"
	SaberItemName               = "saber-sword"
	PineStaffItemName           = "pine-staff"
	OakStaffItemName            = "oak-staff"
	WalnutStaffItemName         = "walnut-staff"
	ShortBowItemName            = "short-bow"
	LongBowItemName             = "long-bow"
	HuntingBowItemName          = "hunting-bow"
	KnifeItemName               = "knife"
	ClubItemName                = "club"
	PikeItemName                = "pike"
	SpearItemName               = "speak"

	// accessory

	MagicRingItemName = "magic-ring"
	MagicHatItemName  = "magic-hat"
	MagicSockItemName = "magic-sock"
)
