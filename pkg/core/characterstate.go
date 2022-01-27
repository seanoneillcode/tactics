package core

type CharacterState struct {
	Health     int
	MaxHealth  int
	Level      int
	Experience int
	Copper     int
	Items      []*Item
	// active effects
	// stats like str, att ...etc
}

func NewCharacterState() *CharacterState {
	return &CharacterState{
		Health:     10,
		MaxHealth:  10,
		Level:      1,
		Experience: 0,
		Copper:     0,
		Items:      []*Item{},
	}
}
