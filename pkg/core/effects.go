package core

import "fmt"

type StateEffect interface {
	Apply(cs *CharacterState)
	Description() string
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
func (h *healthEffect) Description() string {
	return fmt.Sprintf("hp: %v", h.amount)
}
