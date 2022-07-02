package fight

import (
	"fmt"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/input"
)

type TurnPhase string

const (
	SelectActionTurnPhase      = "select-action"
	SelectMoveTargetTurnPhase  = "select-move-target"
	SelectSkillTargetTurnPhase = "select-skill-target"
)

type Action string

const (
	MoveAction  Action = "move"
	NextAction  Action = "next"
	SkillAction Action = "skill"
)

type PlayerController struct {
	SelectedActor       *Actor
	SelectedSkill       *common.Skill
	SelectedActionIndex int
	CurrentActions      []Action
	CurrentTurnPhase    TurnPhase
	PlayerSelection     *Selection
	PossibleMoves       *PossibleMoves
}

func NewPlayerController() *PlayerController {
	return &PlayerController{
		CurrentActions: []Action{
			MoveAction,
			SkillAction,
			NextAction,
		},
		PlayerSelection: NewSelection(),
		PossibleMoves:   NewPossibleMoves(),
	}
}

func (r *PlayerController) StartTurn(state *State) {
	r.SelectedActor = state.PlayerTeam.Actors[0]
	r.CurrentTurnPhase = SelectActionTurnPhase
	r.SelectedActionIndex = 0
	state.ActiveTeam = state.PlayerTeam
	state.PlayerTeam.StartTurn()
}

func (r *PlayerController) Update(delta int64, state *State) {
	if state.ActiveTeam != state.PlayerTeam {
		return
	}

	switch r.CurrentTurnPhase {
	case SelectActionTurnPhase:
		state.Camera.Target(r.SelectedActor)
		if input.IsNextPressed() {
			r.SelectedActor = state.PlayerTeam.GetNextActor(r.SelectedActor)
		}
		if input.IsUpJustPressed() {
			r.SelectedActionIndex -= 1
			if r.SelectedActionIndex == -1 {
				r.SelectedActionIndex = len(r.CurrentActions) - 1
			}
		}
		if input.IsDownJustPressed() {
			r.SelectedActionIndex += 1
			if r.SelectedActionIndex >= len(r.CurrentActions) {
				r.SelectedActionIndex = 0
			}
		}
		if input.IsEnterPressed() {
			currentAction := r.CurrentActions[r.SelectedActionIndex]
			fmt.Println("doing action: ", currentAction)
			switch currentAction {
			case MoveAction:
				if r.SelectedActor.ActionTokensLeft > 0 {
					r.CurrentTurnPhase = SelectMoveTargetTurnPhase
					r.PlayerSelection.SetPos(r.SelectedActor.Pos.X, r.SelectedActor.Pos.Y)
					state.Camera.Target(r.PlayerSelection)
					r.PossibleMoves.GeneratePossibleMoves(state, 3)
				}
			case NextAction:
				r.SelectedActor = state.PlayerTeam.GetNextActor(r.SelectedActor)
			}
		}
	case SelectMoveTargetTurnPhase:
		// select tile to move actor to
		if input.IsDownJustPressed() {
			if r.PossibleMoves.CanMove(r.PlayerSelection.Pos.X, r.PlayerSelection.Pos.Y+common.TileSize) {
				r.PlayerSelection.SetPos(r.PlayerSelection.Pos.X, r.PlayerSelection.Pos.Y+common.TileSize)
			}
		}
		if input.IsUpJustPressed() {
			if r.PossibleMoves.CanMove(r.PlayerSelection.Pos.X, r.PlayerSelection.Pos.Y-common.TileSize) {
				r.PlayerSelection.SetPos(r.PlayerSelection.Pos.X, r.PlayerSelection.Pos.Y-common.TileSize)
			}
		}
		if input.IsLeftJustPressed() {
			if r.PossibleMoves.CanMove(r.PlayerSelection.Pos.X-common.TileSize, r.PlayerSelection.Pos.Y) {
				r.PlayerSelection.SetPos(r.PlayerSelection.Pos.X-common.TileSize, r.PlayerSelection.Pos.Y)
			}
		}
		if input.IsRightJustPressed() {
			if r.PossibleMoves.CanMove(r.PlayerSelection.Pos.X+common.TileSize, r.PlayerSelection.Pos.Y) {
				r.PlayerSelection.SetPos(r.PlayerSelection.Pos.X+common.TileSize, r.PlayerSelection.Pos.Y)
			}
		}
		if input.IsEnterPressed() {
			if !ContainsActor(state, r.PlayerSelection.Pos.X, r.PlayerSelection.Pos.Y) {
				r.SelectedActor.Pos = common.CopyPosition(r.PlayerSelection.Pos)
				r.SelectedActor.ActionTokensLeft -= 1
				if r.SelectedActor.ActionTokensLeft == 0 {
					r.SelectedActor = state.PlayerTeam.GetNextActor(r.SelectedActor)
				}
				r.CurrentTurnPhase = SelectActionTurnPhase
			}
		}
		if input.IsCancelPressed() {
			r.CurrentTurnPhase = SelectActionTurnPhase
		}
	case SelectSkillTargetTurnPhase:
		// select tile to place skill pattern
	}

	if state.PlayerTeam.RemainingActionTokens() == 0 {
		state.AiController.StartTurn(state)
	}
}

func (r PlayerController) Draw(camera *Camera) {
	switch r.CurrentTurnPhase {
	case SelectActionTurnPhase:
		// draw actions and highlighted action
	case SelectMoveTargetTurnPhase:
		r.PossibleMoves.Draw(camera)
		r.PlayerSelection.Draw(camera)
	}
}

func ContainsActor(state *State, x float64, y float64) bool {
	ix := int(x / common.TileSize)
	iy := int(y / common.TileSize)
	for _, actor := range state.PlayerTeam.Actors {
		ax, ay := common.WorldToTile(actor.Pos)
		if ax == ix && ay == iy {
			return true
		}
	}
	for _, actor := range state.AiTeam.Actors {
		ax, ay := common.WorldToTile(actor.Pos)
		if ax == ix && ay == iy {
			return true
		}
	}
	return false
}
