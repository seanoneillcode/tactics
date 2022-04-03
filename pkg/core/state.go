package core

type State struct {
	Player       *Player
	Map          *Map
	ActiveDialog *Dialog // todo move into UI
	Shop         *Shop
	UI           *UI
	Control      *Control
}
