package explore

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
	StatChanges []StatChange
	EquipSlot   string
	ImagePath   string
}

func NewItem(name string) *Item {
	i, ok := AllItems[name]
	if !ok {
		log.Fatalf("unknown item: %s", name)
	}
	newItem := &Item{
		Type:        i.Type,
		Name:        i.Name,
		Description: i.Description,
		CanConsume:  i.CanConsume,
		CanEquip:    i.CanEquip,
		Effects:     i.Effects,
		ImagePath:   i.ImagePath,
		StatChanges: i.StatChanges,
		EquipSlot:   i.EquipSlot,
	}
	return newItem
}

var AllItems = map[string]*Item{
	MouldyBreadItemName: {
		Type:        MouldyBreadItemName,
		Name:        "Mouldy Bread Roll",
		Description: "A bread roll with suspicious green edges.",
		CanConsume:  true,
		CanEquip:    false,
		Effects: []StateEffect{
			&healthEffect{amount: -3},
		},
		ImagePath: "item/mouldy-bread.png",
	},
	BreadItemName: {
		Type:        BreadItemName,
		Name:        "Bread Roll",
		Description: "A wholesome bread roll. A handy snack any time, any place.",
		CanConsume:  true,
		CanEquip:    false,
		Effects: []StateEffect{
			&healthEffect{amount: 1},
		},
		ImagePath: "item/bread.png",
	},
	HerbItemName: {
		Type:        HerbItemName,
		Name:        "Herb",
		Description: "A plant that soothes and heals wounds.",
		CanConsume:  true,
		CanEquip:    false,
		Effects: []StateEffect{
			&healthEffect{amount: 4},
		},
		ImagePath: "item/herb.png",
	},
	PotionItemName: {
		Type:        PotionItemName,
		Name:        "Potion",
		Description: "A distilled herb that significantly heals wounds.",
		CanConsume:  true,
		CanEquip:    false,
		Effects: []StateEffect{
			&healthEffect{amount: 15},
		},
		ImagePath: "item/potion.png",
	},
	EtherItemName: {
		Type:        EtherItemName,
		Name:        "Ether",
		Description: "A distilled crystal that restores magic energy.",
		CanConsume:  true,
		CanEquip:    false,
		Effects: []StateEffect{
			&magicEffect{amount: 15},
		},
		ImagePath: "item/ether.png",
	},
	PaddedArmorItemName: {
		Type:        PaddedArmorItemName,
		Name:        "Padded Armour",
		Description: "A tunic made from several layer of cotton that reduces damage.",
		CanConsume:  false,
		CanEquip:    true,
		StatChanges: []StatChange{
			&defenseChange{amount: 1},
		},
		EquipSlot: "armor",
		ImagePath: "item/padded-armor.png",
	},
	SteelArmorItemName: {
		Type:        SteelArmorItemName,
		Name:        "Plated Armour",
		Description: "A heavy piece of chest armour made from overlapping steel plates.",
		CanConsume:  false,
		CanEquip:    true,
		StatChanges: []StatChange{
			&defenseChange{amount: 5},
			&speedChange{amount: -2},
		},
		EquipSlot: "armor",
		ImagePath: "item/steel-armor.png",
	},
	LeatherArmorItemName: {
		Type:        LeatherArmorItemName,
		Name:        "Leather Armour",
		Description: "A tough chest piece made from thick leather. It smells like burnt sweat.",
		CanConsume:  false,
		CanEquip:    true,
		StatChanges: []StatChange{
			&defenseChange{amount: 2},
			&speedChange{amount: -1},
		},
		EquipSlot: "armor",
		ImagePath: "item/leather-armor.png",
	},
	ChainArmorItemName: {
		Type:        ChainArmorItemName,
		Name:        "Chain Mail Armor",
		Description: "A light and thin shirt of chain mail.",
		CanConsume:  false,
		CanEquip:    true,
		StatChanges: []StatChange{
			&defenseChange{amount: 3},
			&speedChange{amount: -1},
		},
		EquipSlot: "armor",
		ImagePath: "item/chain-mail-armor.png",
	},
	BatteredSwordItemName: {
		Type:        BatteredSwordItemName,
		Name:        "Battered Sword",
		Description: "Dents, scratches and signs of repair. It still cuts if swung hard enough.",
		CanConsume:  false,
		CanEquip:    true,
		StatChanges: []StatChange{
			&defenseChange{amount: 1},
			&attackChange{amount: 2},
		},
		EquipSlot: "weapon",
		ImagePath: "item/short-sword.png",
	},
	RapierItemName: {
		Type:        RapierItemName,
		Name:        "Rapier",
		Description: "Long and gentle with a sharp tip. It's too light to parry with, but that's okay if you're fast enough.",
		CanConsume:  false,
		CanEquip:    true,
		StatChanges: []StatChange{
			&attackChange{amount: 4},
			&speedChange{amount: 1},
		},
		EquipSlot: "weapon",
		ImagePath: "item/short-sword.png",
	},
	SaberItemName: {
		Type:        SaberItemName,
		Name:        "Saber",
		Description: "Curved and sharp with a little weight. It's meant for horse back, but found everywhere.",
		CanConsume:  false,
		CanEquip:    true,
		StatChanges: []StatChange{
			&attackChange{amount: 4},
			&defenseChange{amount: 1},
		},
		EquipSlot: "weapon",
		ImagePath: "item/short-sword.png",
	},
	KnifeItemName: {
		Type:        KnifeItemName,
		Name:        "Knife",
		Description: "A hands length of sharp steel. Used in the kitchen or the pub.",
		CanConsume:  false,
		CanEquip:    true,
		StatChanges: []StatChange{
			&attackChange{amount: 1},
		},
		EquipSlot: "weapon",
		ImagePath: "item/short-sword.png",
	},
	PineStaffItemName: {
		Type:        PineStaffItemName,
		Name:        "Green Pine Staff",
		Description: "A newly cut length of pine, fashioned into a staff. Springy and soft.",
		CanConsume:  false,
		CanEquip:    true,
		StatChanges: []StatChange{
			&attackChange{amount: 2},
			&defenseChange{amount: 2},
		},
		EquipSlot: "weapon",
		ImagePath: "item/short-sword.png",
	},
	HuntingBowItemName: {
		Type:        HuntingBowItemName,
		Name:        "Hunting Bow",
		Description: "A common hunting bow, found in every village.",
		CanConsume:  false,
		CanEquip:    true,
		StatChanges: []StatChange{
			&attackChange{amount: 3},
		},
		EquipSlot: "weapon",
		ImagePath: "item/short-sword.png",
	},
	MagicSockItemName: {
		Type:        MagicSockItemName,
		Name:        "Magic Sock",
		Description: "An old sock with holes. After returning from the lost sock void, it is filled with eldritch knowledge.",
		CanConsume:  false,
		CanEquip:    true,
		StatChanges: []StatChange{
			&magicDefChange{amount: 3},
		},
		EquipSlot: "special",
		ImagePath: "item/ring.png",
	},
	MagicRingItemName: {
		Type:        MagicRingItemName,
		Name:        "Magic Ring",
		Description: "A ring of unknown materials, it glows lightly. It shifts and warps when it's in your peripheral vision.",
		CanConsume:  false,
		CanEquip:    true,
		StatChanges: []StatChange{
			&magicAttChange{amount: 3},
		},
		EquipSlot: "special",
		ImagePath: "item/ring.png",
	},
	MagicHatItemName: {
		Type:        MagicHatItemName,
		Name:        "Magic Hat",
		Description: "A unfashionable hat, dusty and stained. It's feels like wearing an old couch on your head.",
		CanConsume:  false,
		CanEquip:    true,
		StatChanges: []StatChange{
			&magicMaxChange{amount: 3},
		},
		EquipSlot: "special",
		ImagePath: "item/ring.png",
	},
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
