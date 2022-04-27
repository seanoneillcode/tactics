package core

type State struct {
	Player           *Player
	TeamState        *TeamState
	Map              *Map
	DialogHandler    *DialogHandler
	Shop             *Shop
	UI               *UI
	Control          *Control
	TotalElapsedTime int64
}
