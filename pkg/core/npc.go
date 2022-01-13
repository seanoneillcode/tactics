package core

import "github.com/hajimehoshi/ebiten/v2"

type Npc struct {
	character *Character
}

func NewNpc(name string) *Npc {
	return &Npc{
		character: NewCharacter(name + ".png"),
	}
}

func (n *Npc) Draw(screen *ebiten.Image) {
	n.character.Draw(screen)
}

func (n *Npc) Update(delta int64, state *State) {
	n.character.Update(delta)
}

func (n *Npc) SetPosition(x int, y int) {
	n.character.x = float64(x)
	n.character.y = float64(y)
}

func (n *Npc) GetPosition() (float64, float64) {
	return n.character.x, n.character.y
}
