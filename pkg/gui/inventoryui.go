package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/core"
)

type InventoryUi struct {
}

func NewInventoryUi() *InventoryUi {
	return &InventoryUi{}
}

func (i *InventoryUi) Draw(screen *ebiten.Image) {

}

func (i *InventoryUi) Update(delta int64, state *core.State) {

}