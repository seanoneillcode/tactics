package explore

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
	r.EquippedStats = r.ApplyItemsToStats(r.EquippedItems, r.BaseStats)
}

func (r *CharacterState) UnEquipItem(slot string) {
	delete(r.EquippedItems, slot)
	r.EquippedStats = r.ApplyItemsToStats(r.EquippedItems, r.BaseStats)
}

func (r *CharacterState) ApplyItemsToStats(items map[string]*Item, stats *Stats) *Stats {
	newStats := &Stats{
		MaxHealth:      stats.MaxHealth,
		MaxMagic:       stats.MaxMagic,
		Level:          stats.Level,
		Experience:     stats.Experience,
		AttackSkill:    stats.AttackSkill,
		AttackStrength: stats.AttackStrength,
		Defence:        stats.Defence,
		Agility:        stats.Agility,
		Speed:          stats.Speed,
		MagicSkill:     stats.MagicSkill,
		MagicDef:       stats.MagicDef,
	}
	for _, ei := range items {
		for _, sc := range ei.StatChanges {
			sc.Apply(newStats)
		}
	}
	return newStats
}

func (r *CharacterState) GetSkills() []string {
	var skillNames []string
	for _, item := range r.EquippedItems {
		for _, skill := range item.Skills {
			skillNames = append(skillNames, skill)
		}
	}
	return skillNames
}
