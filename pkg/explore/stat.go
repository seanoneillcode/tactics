package explore

type StatChange interface {
	Apply(stats *Stats)
}

type defenseChange struct {
	amount int
}

func (h *defenseChange) Apply(stats *Stats) {
	stats.Defence = stats.Defence + h.amount
}

type speedChange struct {
	amount int
}

func (h *speedChange) Apply(stats *Stats) {
	stats.Speed = stats.Speed + h.amount
}

type attackChange struct {
	amount int
}

func (h *attackChange) Apply(stats *Stats) {
	stats.AttackSkill = stats.AttackSkill + h.amount
}

type magicAttChange struct {
	amount int
}

func (h *magicAttChange) Apply(stats *Stats) {
	stats.MagicSkill = stats.MagicSkill + h.amount
}

type magicDefChange struct {
	amount int
}

func (h *magicDefChange) Apply(stats *Stats) {
	stats.MagicDef = stats.MagicDef + h.amount
}

type magicMaxChange struct {
	amount int
}

func (h *magicMaxChange) Apply(stats *Stats) {
	stats.MaxMagic = stats.MaxMagic + h.amount
}
