package core

type CharacterState struct {
	// state
	Health              int
	Magic               int
	ActiveStatusEffects []string // poison, etc
	EquippedItems       map[string]*Item

	// stats
	Stats *Stats
}

type Stats struct {
	MaxHealth      int
	Level          int
	Experience     int
	AttackSkill    int // to hit
	AttackStrength int // physical damage
	Defence        int // damage reduction
	Agility        int // dodge hit
	Speed          int // move/turn speed
	MagicSkill     int // magical damage
	MagicDef       int // magical damage reduction
}

func NewCharacterState() *CharacterState {
	return &CharacterState{
		Health: 10,
		Stats: &Stats{
			MaxHealth:  10,
			Level:      1,
			Experience: 0,
		},
		ActiveStatusEffects: []string{},
		EquippedItems:       map[string]*Item{},
	}
}
