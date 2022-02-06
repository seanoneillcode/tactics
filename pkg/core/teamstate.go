package core

type TeamState struct {
	Characters []*CharacterState
	Money      int
	Items      []*Item
}

func NewTeamState() *TeamState {
	return &TeamState{
		Characters: []*CharacterState{},
		Money:      200,
		Items:      []*Item{},
	}
}

func (t *TeamState) RestoreHealth() {
	for _, c := range t.Characters {
		c.Health = c.Stats.MaxHealth
	}
}

func (t *TeamState) BuyItem(item *Item, cost int) {
	t.Items = append(t.Items, item)
	t.Money = t.Money - cost
}

func (t *TeamState) Pickup(pickup *Pickup) {
	t.Items = append(t.Items, NewItem(pickup.itemName))
}

func (t *TeamState) RemoveItem(index int) {
	var j int
	for i, n := range t.Items {
		if i != index {
			t.Items[j] = n
			j++
		}
	}
	t.Items = t.Items[:j]
}

func (t *TeamState) ConsumeItem(index int) {
	// todo select character
	selectedCharacter := t.Characters[0]

	item := t.Items[index]
	for _, e := range item.Effects {
		e.Apply(selectedCharacter)
	}
	t.RemoveItem(index)
}

func (t *TeamState) EquipItem(index int) {
	// todo select character
	selectedCharacter := t.Characters[0]

	item := t.Items[index]
	existingItem, ok := selectedCharacter.EquippedItems[item.EquipSlot]
	if ok {
		// slot already has item, remove it and put it back into the inventory
		t.Items = append(t.Items, existingItem)
	}
	selectedCharacter.EquippedItems[item.EquipSlot] = item
	t.RemoveItem(index)
}
