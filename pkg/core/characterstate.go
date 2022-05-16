package core

type CharacterState struct {
	// state
	Health              int
	Magic               int
	ActiveStatusEffects []string // poison, etc
	EquippedItems       map[string]*Item
	Name                string

	// stats
	BaseStats     *Stats
	EquippedStats *Stats
}

type Stats struct {
	MaxHealth      int
	MaxMagic       int
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

func NewCharacterState(name string) *CharacterState {
	return &CharacterState{
		Name:   name,
		Health: 10,
		BaseStats: &Stats{
			MaxHealth:  10,
			MaxMagic:   10,
			Level:      1,
			Experience: 0,
		},
		EquippedStats: &Stats{
			MaxHealth: 10,
			MaxMagic:  10,
		},
		ActiveStatusEffects: []string{},
		EquippedItems:       map[string]*Item{},
	}
}

func (r *CharacterState) EquipItem(slot string, item *Item) {
	r.EquippedItems[slot] = item
	r.updateStats()
}

func (r *CharacterState) UnEquipItem(slot string) {
	delete(r.EquippedItems, slot)
	r.updateStats()
}

func (r *CharacterState) updateStats() {
	newStats := &Stats{
		MaxHealth:      r.BaseStats.MaxHealth,
		MaxMagic:       r.BaseStats.MaxMagic,
		Level:          r.BaseStats.Level,
		Experience:     r.BaseStats.Experience,
		AttackSkill:    r.BaseStats.AttackSkill,
		AttackStrength: r.BaseStats.AttackStrength,
		Defence:        r.BaseStats.Defence,
		Agility:        r.BaseStats.Agility,
		Speed:          r.BaseStats.Speed,
		MagicSkill:     r.BaseStats.MagicSkill,
		MagicDef:       r.BaseStats.MagicDef,
	}
	for _, ei := range r.EquippedItems {
		for _, sc := range ei.StatChanges {
			sc.Apply(newStats)
		}
	}
	r.EquippedStats = newStats
}
