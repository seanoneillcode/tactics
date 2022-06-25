package explore

import (
	"fmt"
)

type StateEffect interface {
	Apply(cs *CharacterState)
	Description(cs *CharacterState) string
}

type healthEffect struct {
	amount int
}

func (h *healthEffect) Apply(cs *CharacterState) {
	cs.Health = cs.Health + h.amount
	if cs.Health > cs.EquippedStats.MaxHealth {
		cs.Health = cs.EquippedStats.MaxHealth
	}
}
func (h *healthEffect) Description(cs *CharacterState) string {
	newValue := cs.Health + h.amount
	if newValue > cs.EquippedStats.MaxHealth {
		newValue = cs.EquippedStats.MaxHealth
	}
	return fmt.Sprintf("health\n%v > %v", cs.Health, newValue)
}

type magicEffect struct {
	amount int
}

func (h *magicEffect) Apply(cs *CharacterState) {
	cs.Magic = cs.Magic + h.amount
	if cs.Magic > cs.EquippedStats.MaxMagic {
		cs.Magic = cs.EquippedStats.MaxMagic
	}
}
func (h *magicEffect) Description(cs *CharacterState) string {
	newValue := cs.Magic + h.amount
	if newValue > cs.EquippedStats.MaxMagic {
		newValue = cs.EquippedStats.MaxMagic
	}
	return fmt.Sprintf("magic\n%v > %v", cs.Magic, newValue)
}
