package core

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
	if cs.Health > cs.Stats.MaxHealth {
		cs.Health = cs.Stats.MaxHealth
	}
}
func (h *healthEffect) Description(cs *CharacterState) string {
	newValue := cs.Health + h.amount
	if newValue > cs.Stats.MaxHealth {
		newValue = cs.Stats.MaxHealth
	}
	return fmt.Sprintf("healt\n%v > %v", cs.Health, newValue)
}

type magicEffect struct {
	amount int
}

func (h *magicEffect) Apply(cs *CharacterState) {
	cs.Magic = cs.Magic + h.amount
	if cs.Magic > cs.Stats.MaxMagic {
		cs.Magic = cs.Stats.MaxMagic
	}
}
func (h *magicEffect) Description(cs *CharacterState) string {
	newValue := cs.Magic + h.amount
	if newValue > cs.Stats.MaxMagic {
		newValue = cs.Stats.MaxMagic
	}
	return fmt.Sprintf("magic\n%v > %v", cs.Magic, newValue)
}
