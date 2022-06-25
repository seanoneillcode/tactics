package fight

type AiController struct {
	SelectedActor *Actor
}

func (r *AiController) StartTurn(state *State) {
	r.SelectedActor = state.AiTeam.Actors[0]
	state.ActiveTeam = state.AiTeam
	state.AiTeam.StartTurn()
}

func (r *AiController) Update(delta int64, state *State) {
	if state.ActiveTeam != state.AiTeam {
		return
	}
}
