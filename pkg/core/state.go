package core

type State struct {
	Player           *Player
	TeamState        *TeamState
	Map              *Map
	DialogHandler    *DialogHandler
	ModeManager      *ModeManager
	Shop             *Shop
	UI               *UI
	Control          *Control
	TotalElapsedTime int64
}
