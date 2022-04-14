package core

type State struct {
	Player           *Player
	TeamState        *TeamState
	Map              *Map
	ActiveDialog     *Dialog
	Shop             *Shop
	UI               *UI
	Control          *Control
	TotalElapsedTime int64
}
