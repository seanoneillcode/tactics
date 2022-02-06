package core

type State struct {
	Player       *Player
	Map          *Map
	ActiveDialog *Dialog
	Shop         *Shop
	Inventory    *Inventory
}
