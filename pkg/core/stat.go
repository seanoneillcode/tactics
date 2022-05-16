package core

import "fmt"

type StatChange interface {
	Apply(stats *Stats)
	Description(stats *Stats) string
}

type defenseChange struct {
	amount int
}

func (h *defenseChange) Apply(stats *Stats) {
	stats.Defence = stats.Defence + h.amount
}
func (h *defenseChange) Description(stats *Stats) string {
	return fmt.Sprintf("defense %v > %v", stats.Defence, stats.Defence+h.amount)
}

type speedChange struct {
	amount int
}

func (h *speedChange) Apply(stats *Stats) {
	stats.Speed = stats.Speed + h.amount
}
func (h *speedChange) Description(stats *Stats) string {
	return fmt.Sprintf("speed %v > %v", stats.Speed, stats.Speed+h.amount)
}

type attackChange struct {
	amount int
}

func (h *attackChange) Apply(stats *Stats) {
	stats.AttackSkill = stats.AttackSkill + h.amount
}
func (h *attackChange) Description(stats *Stats) string {
	return fmt.Sprintf("attack %v > %v", stats.AttackSkill, stats.AttackSkill+h.amount)
}

type magicAttChange struct {
	amount int
}

func (h *magicAttChange) Apply(stats *Stats) {
	stats.MagicSkill = stats.MagicSkill + h.amount
}
func (h *magicAttChange) Description(stats *Stats) string {
	return fmt.Sprintf("magic %v > %v", stats.MagicSkill, stats.MagicSkill+h.amount)
}

type magicDefChange struct {
	amount int
}

func (h *magicDefChange) Apply(stats *Stats) {
	stats.MagicDef = stats.MagicDef + h.amount
}
func (h *magicDefChange) Description(stats *Stats) string {
	return fmt.Sprintf("magic def %v > %v", stats.MagicDef, stats.MagicDef+h.amount)
}

type magicMaxChange struct {
	amount int
}

func (h *magicMaxChange) Apply(stats *Stats) {
	stats.MaxMagic = stats.MaxMagic + h.amount
}
func (h *magicMaxChange) Description(stats *Stats) string {
	return fmt.Sprintf("max magic %v > %v", stats.MaxMagic, stats.MaxMagic+h.amount)
}
