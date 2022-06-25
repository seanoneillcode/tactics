package fight

type Team struct {
	Actors []*Actor
}

func NewTeam(actors []*Actor) *Team {
	return &Team{
		Actors: actors,
	}
}

func (r *Team) StartTurn() {
	for _, a := range r.Actors {
		a.ActionTokensLeft = 2
	}
}

func (r *Team) RemainingActionTokens() int {
	total := 0
	for _, a := range r.Actors {
		total += a.ActionTokensLeft
	}
	return total
}
