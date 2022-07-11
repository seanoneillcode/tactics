package fight

import "github.com/seanoneillcode/go-tactics/pkg/common"

type Effect interface {
	DoEffect(state *State, pos common.Tile) bool
}

type DamageActorEffect struct {
	Amount int
}

func (r *DamageActorEffect) DoEffect(state *State, pos common.Tile) bool {
	actor := GetActorAtPos(state, pos)
	if actor == nil {
		return false
	}
	actor.TakeDamage(r.Amount)
	return true
}

type HealActorEffect struct {
	Amount int
}

func (r *HealActorEffect) DoEffect(state *State, pos common.Tile) bool {
	actor := GetActorAtPos(state, pos)
	if actor == nil {
		return false
	}
	actor.TakeHealing(r.Amount)
	return true
}

func GetActorAtPos(state *State, pos common.Tile) *Actor {
	for _, actor := range state.PlayerTeam.Actors {
		apos := common.WorldToTile(actor.Pos)
		if apos.X == pos.X && apos.Y == pos.Y {
			return actor
		}
	}
	for _, actor := range state.AiTeam.Actors {
		apos := common.WorldToTile(actor.Pos)
		if apos.X == pos.X && apos.Y == pos.Y {
			return actor
		}
	}
	return nil
}
