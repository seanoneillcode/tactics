package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/core"
)

type UI interface {
	Draw(screen *ebiten.Image)
	Update(delta int64, state *core.State)
}
