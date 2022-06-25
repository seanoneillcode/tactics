package fight

type TurnPhase string

const (
	SelectActionTurnPhase      = "select-action"
	SelectMoveTargetTurnPhase  = "select-move-target"
	SelectSkillTargetTurnPhase = "select-skill-target"
)

type PlayerController struct {
	SelectedActor    *Actor
	CurrentTurnPhase TurnPhase
}

func NewPlayerController() *PlayerController {
	return &PlayerController{}
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
		// select move, skill or switch to next/prev actor
	case SelectMoveTargetTurnPhase:
		// select tile to move actor to
	case SelectSkillTargetTurnPhase:
		// select tile to place skill pattern
	}

	if state.PlayerTeam.RemainingActionTokens() == 0 {
		state.AiController.StartTurn(state)
	}
}
