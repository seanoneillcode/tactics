package explore

type State struct {
	Player           *Player
	TeamState        *TeamState
	Map              *Map
	DialogHandler    *DialogHandler
	ModeManager      *ModeManager
	Shop             *Shop
	UI               *UI
	Control          *Control
	Camera           *Camera
	TotalElapsedTime int64
}

func (r *State) Update(delta int64) {
	r.Map.Update(delta, r)
	r.Player.Update(delta, r)
	r.Shop.Update(delta, r)
	r.ModeManager.Update(delta, r)
	r.Camera.Update(delta, r)
}
