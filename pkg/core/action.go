package core

import "github.com/seanoneillcode/go-tactics/pkg/common"

type Action struct {
	name string
	pos  *common.Position
}

func NewAction(name string, x float64, y float64) *Action {
	return &Action{
		name: name,
		pos: &common.Position{
			X: x,
			Y: y,
		},
	}
}

func (a *Action) GetPosition() *common.Position {
	return a.pos
}
