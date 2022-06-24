package core

import (
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Enemy struct {
	character *Character
	name      string
}

func NewEnemy(name string) *Enemy {
	return &Enemy{
		name:      name,
		character: NewCharacter("enemy/" + name + ".png"),
	}
}

func (n *Enemy) Draw(camera *Camera) {
	n.character.Draw(camera)
}

func (n *Enemy) Update(delta int64, state *State) {
	n.character.Update(delta)
}

func (n *Enemy) SetPosition(x int, y int) {
	n.character.pos = common.PositionFromInt(x, y)
}

func (n *Enemy) GetPosition() *common.Position {
	return n.character.pos
}
