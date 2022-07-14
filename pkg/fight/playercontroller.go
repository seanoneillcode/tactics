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
	DoneAction  Action = "done"
)

type PlayerController struct {
	SelectedActor        *Actor
	SelectedActionIndex  int
	CurrentActions       []Action
	CurrentTurnPhase     TurnPhase
	PlayerSelection      *Selection
	PossibleMoves        *PossibleMoves
	PossibleTargets      *PossibleTargets
	EffectedTiles        []common.Tile
	EffectedSelection    *common.Sprite
	EffectSelectionTimer int64
}

func NewPlayerController() *PlayerController {
	return &PlayerController{
		CurrentActions: []Action{
			MoveAction,
			SkillAction,
			DoneAction,
		},
		EffectedTiles:     []common.Tile{},
		PlayerSelection:   NewSelection(),
		PossibleMoves:     NewPossibleMoves(),
		PossibleTargets:   NewPossibleTargets(),
		EffectedSelection: common.NewSprite("effect-tile-selection.png"),
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
		if input.IsRightJustPressed() {
			r.SelectedActor = state.PlayerTeam.GetNextActor(r.SelectedActor)
		}
		if input.IsLeftJustPressed() {
			r.SelectedActor = state.PlayerTeam.GetNextActor(r.SelectedActor)
		}
		if input.IsEnterPressed() {
			currentAction := r.CurrentActions[r.SelectedActionIndex]
			fmt.Println("doing action: ", currentAction)
			switch currentAction {
			case MoveAction:
				if r.SelectedActor.HasMove {
					r.CurrentTurnPhase = SelectMoveTargetTurnPhase
					r.PlayerSelection.SetPos(r.SelectedActor.Pos.X, r.SelectedActor.Pos.Y)
					state.Camera.Target(r.PlayerSelection)
					r.PossibleMoves.GeneratePossibleMoves(state, 3)
				}
			case SkillAction:
				if r.SelectedActor.ActionTokensLeft > 0 {
					r.CurrentTurnPhase = SelectSkillTargetTurnPhase
					r.PlayerSelection.SetPos(r.SelectedActor.Pos.X, r.SelectedActor.Pos.Y)
					state.Camera.Target(r.PlayerSelection)
					r.PossibleTargets.GeneratePossibleTargets(state)
					effectSourcePos := common.WorldToTile(r.PlayerSelection.Pos)
					r.EffectedTiles = r.SelectedActor.Skills[0].EffectPattern.GetPattern(effectSourcePos, state)
				}
			case DoneAction:
				r.SelectedActor.ActionTokensLeft = 0
				r.SelectedActor = state.PlayerTeam.GetNextActor(r.SelectedActor)
				r.SelectedActionIndex = 0
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
			actor := GetActorAtPos(state, common.WorldToTile(r.PlayerSelection.Pos))
			if actor == nil {
				r.SelectedActor.Pos = common.CopyPosition(r.PlayerSelection.Pos)
				r.SelectedActor.HasMove = false
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
		if input.IsDownJustPressed() {
			next := findNextTile(state, r.PlayerSelection.Pos, 0, 1)
			if next == nil {
				next = findNextTile(state, r.SelectedActor.Pos, 0, 1)
			}
			if next != nil {
				r.PlayerSelection.SetPos(float64(next.X*common.TileSize), float64(next.Y*common.TileSize))
				effectSourcePos := common.WorldToTile(r.PlayerSelection.Pos)
				r.EffectedTiles = r.SelectedActor.Skills[0].EffectPattern.GetPattern(effectSourcePos, state)
			}
		}
		if input.IsUpJustPressed() {
			next := findNextTile(state, r.PlayerSelection.Pos, 0, -1)
			if next == nil {
				next = findNextTile(state, r.SelectedActor.Pos, 0, -1)
			}
			if next != nil {
				r.PlayerSelection.SetPos(float64(next.X*common.TileSize), float64(next.Y*common.TileSize))
				effectSourcePos := common.WorldToTile(r.PlayerSelection.Pos)
				r.EffectedTiles = r.SelectedActor.Skills[0].EffectPattern.GetPattern(effectSourcePos, state)
			}
		}
		if input.IsLeftJustPressed() {
			next := findNextTile(state, r.PlayerSelection.Pos, -1, 0)
			if next == nil {
				next = findNextTile(state, r.SelectedActor.Pos, -1, 0)
			}
			if next != nil {
				r.PlayerSelection.SetPos(float64(next.X*common.TileSize), float64(next.Y*common.TileSize))
				effectSourcePos := common.WorldToTile(r.PlayerSelection.Pos)
				r.EffectedTiles = r.SelectedActor.Skills[0].EffectPattern.GetPattern(effectSourcePos, state)
			}
		}
		if input.IsRightJustPressed() {
			next := findNextTile(state, r.PlayerSelection.Pos, 1, 0)
			if next == nil {
				next = findNextTile(state, r.SelectedActor.Pos, 1, 0)
			}
			if next != nil {
				r.PlayerSelection.SetPos(float64(next.X*common.TileSize), float64(next.Y*common.TileSize))
				effectSourcePos := common.WorldToTile(r.PlayerSelection.Pos)
				r.EffectedTiles = r.SelectedActor.Skills[0].EffectPattern.GetPattern(effectSourcePos, state)
			}
		}
		if input.IsEnterPressed() {
			if r.PossibleTargets.CanTarget(r.PlayerSelection.Pos) {
				// do the skill on the selected position
				for _, effect := range r.SelectedActor.Skills[0].Effects {
					for _, target := range r.EffectedTiles {
						effect.DoEffect(state, target)
					}
				}
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

	}

	if state.PlayerTeam.RemainingActionTokens() == 0 {
		state.AiController.StartTurn(state)
	}
}

func findNextTile(state *State, pos *common.Position, dx int, dy int) *common.Tile {

	var posmap = map[int]map[int]bool{}
	for _, t := range state.PlayerController.PossibleTargets.targets {
		v, ok := posmap[t.X]
		if !ok {
			posmap[t.X] = map[int]bool{
				t.Y: true,
			}
		} else {
			v[t.Y] = true
		}
	}

	count := 0
	start := common.WorldToTile(pos)
	if dx != 0 {
		index := start.X
		for count < 100 {
			index = index + dx
			v, ok := posmap[index]
			if ok {
				_, oky := v[start.Y]
				if oky {
					return &common.Tile{
						X: index,
						Y: start.Y,
					}
				}
			}
			count += 1
		}

	}
	if dy != 0 {
		index := start.Y
		for count < 100 {
			index = index + dy
			v, ok := posmap[start.X]
			if ok {
				_, oky := v[index]
				if oky {
					return &common.Tile{
						X: start.X,
						Y: index,
					}
				}
			}
			count += 1
		}
	}

	return nil
}

func (r PlayerController) Draw(camera *Camera) {
	switch r.CurrentTurnPhase {
	case SelectActionTurnPhase:
		// draw actions and highlighted action
	case SelectMoveTargetTurnPhase:
		r.PossibleMoves.Draw(camera)
		r.PlayerSelection.Draw(camera)
	case SelectSkillTargetTurnPhase:
		r.PossibleTargets.Draw(camera)
		for _, t := range r.EffectedTiles {
			r.EffectedSelection.SetPosition(float64(t.X*common.TileSize), float64(t.Y*common.TileSize))
			r.EffectedSelection.Draw(camera)
		}
		r.PlayerSelection.Draw(camera)
	}
}
