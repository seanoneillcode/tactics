package core

type TeamState struct {
	Characters []*CharacterState
	Money      int
	Items      []*Item
}

func NewTeamState() *TeamState {
	return &TeamState{
		Characters: []*CharacterState{},
		Money:      10,
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
	item := t.Items[index]
	for _, e := range item.Effects {
		e.Apply(t.Characters[0]) // todo select character
	}
	t.RemoveItem(index)
}
