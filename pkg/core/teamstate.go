package core

import (
	"log"
	"sort"
)

type TeamState struct {
	Characters  []*CharacterState
	Money       int
	ItemHolders map[string]*ItemHolder
	Iteration   int
	ItemList    []string
}

type ItemHolder struct {
	Item   *Item
	Amount int
}

func NewTeamState() *TeamState {
	ts := &TeamState{
		Characters: []*CharacterState{
			NewCharacterState("alice"),
			NewCharacterState("bob"),
			NewCharacterState("carl"),
		},
		Money:       200,
		ItemHolders: map[string]*ItemHolder{},
		ItemList:    []string{},
	}
	ts.Pickup(&Pickup{itemName: BreadItemName})
	ts.Pickup(&Pickup{itemName: BreadItemName})
	ts.Pickup(&Pickup{itemName: BreadItemName})
	ts.Pickup(&Pickup{itemName: MouldyBreadItemName})
	ts.Pickup(&Pickup{itemName: PotionItemName})
	ts.Pickup(&Pickup{itemName: PotionItemName})
	ts.Pickup(&Pickup{itemName: EtherItemName})
	ts.Pickup(&Pickup{itemName: PaddedArmorItemName})
	ts.Pickup(&Pickup{itemName: PaddedArmorItemName})
	ts.Pickup(&Pickup{itemName: SteelArmorItemName})
	ts.Pickup(&Pickup{itemName: SteelArmorItemName})
	ts.Pickup(&Pickup{itemName: SteelArmorItemName})
	// for testing
	ts.Characters[0].Health = 2
	ts.Characters[1].Magic = 3
	ts.Characters[2].ActiveStatusEffects = []string{"asleep"}

	ts.EquipItem("Padded Armour", 0)

	return ts
}

func (t *TeamState) RestoreHealth() {
	for _, c := range t.Characters {
		c.Health = c.Stats.MaxHealth
	}
}

func (t *TeamState) BuyItem(item *Item, cost int) {
	ti, has := t.ItemHolders[item.Name]
	if !has {
		t.ItemHolders[item.Name] = &ItemHolder{
			Item:   item,
			Amount: 1,
		}
	} else {
		ti.Amount = ti.Amount + 1
	}
	t.Money = t.Money - cost
	t.refreshItemList()
}

func (t *TeamState) Pickup(pickup *Pickup) {
	ni := NewItem(pickup.itemName)
	ti, has := t.ItemHolders[ni.Name]
	if !has {
		t.ItemHolders[ni.Name] = &ItemHolder{
			Item:   ni,
			Amount: 1,
		}
	} else {
		ti.Amount = ti.Amount + 1
	}
	t.refreshItemList()
}

func (t *TeamState) RemoveItem(name string) {
	ti, has := t.ItemHolders[name]
	if !has {
		log.Fatal("tried to remove an item that doesn't exist in the inventory " + name)
	}
	ti.Amount = ti.Amount - 1
	if ti.Amount == 0 {
		delete(t.ItemHolders, name)
	}
	t.refreshItemList()
}

func (t *TeamState) ConsumeItem(name string, index int) {
	selectedCharacter := t.Characters[index]

	ti := t.ItemHolders[name]
	for _, e := range ti.Item.Effects {
		e.Apply(selectedCharacter)
	}
	t.RemoveItem(name)
	t.refreshItemList()
}

func (t *TeamState) EquipItem(name string, index int) {
	selectedCharacter := t.Characters[index]

	ti := t.ItemHolders[name]
	selectedCharacter.EquippedItems[ti.Item.EquipSlot] = ti.Item
	t.refreshItemList()
}

func (t *TeamState) GetItem(name string) *Item {
	ti, has := t.ItemHolders[name]
	if !has {
		log.Fatal("tried to get a non existent item")
		return nil
	}
	return ti.Item
}

func (t *TeamState) GetItemList() []string {
	var list []string
	for k := range t.ItemHolders {
		list = append(list, k)
	}
	sort.Strings(list)
	return list
}

func (t *TeamState) GetItemWithIndex(index int) *Item {
	return t.GetItem(t.ItemList[index])
}

func (t *TeamState) refreshItemList() {
	t.Iteration = t.Iteration + 1
	t.ItemList = t.GetItemList()
}

func (t *TeamState) HasItems() bool {
	return len(t.ItemHolders) > 0
}
