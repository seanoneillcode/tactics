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
}

func NewPlayerController() *PlayerController {
	return &PlayerController{
		CurrentActions: []Action{
			MoveAction,
			SkillAction,
			NextAction,
		},
		PlayerSelection: NewSelection(),
	}
}

func (r *PlayerController) StartTurn(state *State) {
	r.SelectedActor = state.PlayerTeam.Actors[0]
	r.CurrentTurnPhase = SelectActionTurnPhase
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
			fmt.Println("action: ", r.CurrentActions[r.SelectedActionIndex])
		}
		if input.IsDownJustPressed() {
			r.SelectedActionIndex += 1
			if r.SelectedActionIndex >= len(r.CurrentActions) {
				r.SelectedActionIndex = 0
			}
			fmt.Println("action: ", r.CurrentActions[r.SelectedActionIndex])
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
				}
			case NextAction:
				r.SelectedActor = state.PlayerTeam.GetNextActor(r.SelectedActor)
			}
		}
	case SelectMoveTargetTurnPhase:
		// select tile to move actor to
		if input.IsDownJustPressed() {
			r.PlayerSelection.SetPos(r.PlayerSelection.Pos.X, r.PlayerSelection.Pos.Y+common.TileSize)
		}
		if input.IsUpJustPressed() {
			r.PlayerSelection.SetPos(r.PlayerSelection.Pos.X, r.PlayerSelection.Pos.Y-common.TileSize)
		}
		if input.IsLeftJustPressed() {
			r.PlayerSelection.SetPos(r.PlayerSelection.Pos.X-common.TileSize, r.PlayerSelection.Pos.Y)
		}
		if input.IsRightJustPressed() {
			r.PlayerSelection.SetPos(r.PlayerSelection.Pos.X+common.TileSize, r.PlayerSelection.Pos.Y)
		}
		if input.IsEnterPressed() {
			r.SelectedActor.Pos = common.CopyPosition(r.PlayerSelection.Pos)
			r.SelectedActor.ActionTokensLeft -= 1
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
		r.PlayerSelection.Draw(camera)
	}
}
