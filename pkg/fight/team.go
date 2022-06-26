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

func (r *Team) Draw(camera *Camera) {
	for _, actor := range r.Actors {
		actor.Draw(camera)
	}
}

func (r *Team) Update(delta int64, state *State) {
	for _, actor := range r.Actors {
		actor.Update(delta, state)
	}
}

func (r *Team) GetNextActor(current *Actor) *Actor {
	for index, a := range r.Actors {
		if a == current {
			nextPlayerIndex := index + 1
			if nextPlayerIndex >= len(r.Actors) {
				nextPlayerIndex = 0
			}
			return r.Actors[nextPlayerIndex]
		}
	}
	return nil
}
