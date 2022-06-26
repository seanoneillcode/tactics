package fight

import "github.com/seanoneillcode/go-tactics/pkg/common"

type Actor struct {
	Name             string
	ActionTokensLeft int
	Sprite           *common.Sprite
	Pos              *common.Position
}

func (a Actor) Draw(camera *Camera) {
	a.Sprite.Draw(camera)
}

func (a Actor) Update(delta int64, state *State) {
	a.Sprite.SetPosition(a.Pos.X, a.Pos.Y)
}

func NewActor(name string) *Actor {
	return &Actor{
		Name:   name,
		Sprite: common.NewSprite("actors/" + name + ".png"),
		Pos:    &common.Position{},
	}
}

func (a Actor) GetPos() *common.Position {
	return a.Pos
}
