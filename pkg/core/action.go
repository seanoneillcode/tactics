package core

import "github.com/seanoneillcode/go-tactics/pkg/common"

type Action struct {
	name string
	pos  *common.Vector
}

func NewAction(name string, x float64, y float64) *Action {
	return &Action{
		name: name,
		pos: &common.Vector{
			X: x,
			Y: y,
		},
	}
}

func (a *Action) GetPosition() *common.Vector {
	return a.pos
}
