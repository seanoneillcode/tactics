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
	return fmt.Sprintf("health %v -> %v", cs.Health, newValue)
}
