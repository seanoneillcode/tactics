package fight

import (
	"fmt"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Actor struct {
	Name             string
	ActionTokensLeft int
	HasMove          bool
	Sprite           *common.Sprite
	Pos              *common.Position
	Skills           []*Skill
	Health           int
}

func (a *Actor) Draw(camera *Camera) {
	a.Sprite.Draw(camera)
}

func (a *Actor) Update(delta int64, state *State) {
	a.Sprite.SetPosition(a.Pos.X, a.Pos.Y)
	// check health
}

func NewActor(name string, skills []*Skill) *Actor {
	return &Actor{
		Name:   name,
		Sprite: common.NewSprite("actors/" + name + ".png"),
		Pos:    &common.Position{},
		Skills: skills,
		Health: 2,
	}
}

func (a *Actor) GetPos() *common.Position {
	return a.Pos
}

func (a *Actor) SetPos(pos *common.Position) {
	a.Pos = &common.Position{
		X: pos.X,
		Y: pos.Y,
	}
}

func (a *Actor) TakeDamage(amount int) {
	a.Health -= amount // check death ?
	fmt.Printf("reduced actor %s health by %v to %v\n", a.Name, amount, a.Health)
}

func (a *Actor) TakeHealing(amount int) {
	a.Health += amount // check max ?
	fmt.Printf("healed actor %s health by %v to %v\n", a.Name, amount, a.Health)
}
