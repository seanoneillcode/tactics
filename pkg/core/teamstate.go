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
		Characters:  []*CharacterState{},
		Money:       200,
		ItemHolders: map[string]*ItemHolder{},
		ItemList:    []string{},
	}
	ts.Pickup(&Pickup{itemName: BreadItemName})
	ts.Pickup(&Pickup{itemName: BreadItemName})
	ts.Pickup(&Pickup{itemName: BreadItemName})
	ts.Pickup(&Pickup{itemName: PotionItemName})
	ts.Pickup(&Pickup{itemName: PaddedArmorItemName})
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

func (t *TeamState) ConsumeItem(name string) {
	// todo select character
	selectedCharacter := t.Characters[0]

	ti := t.ItemHolders[name]
	for _, e := range ti.Item.Effects {
		e.Apply(selectedCharacter)
	}
	t.RemoveItem(name)
	t.refreshItemList()
}

func (t *TeamState) EquipItem(name string) {
	// todo select character
	selectedCharacter := t.Characters[0]

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
