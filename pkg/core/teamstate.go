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
		c.Health = c.MaxHealth
	}
}

func (t *TeamState) BuyItem(item *Item, cost int) {
	t.Items = append(t.Items, item)
	t.Money = t.Money - cost
}

func (t *TeamState) Pickup(pickup *Pickup) {
	t.Items = append(t.Items, NewItem(pickup.itemName))
}
