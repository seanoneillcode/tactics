package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
)

type bg struct {
	image *ebiten.Image
}

func NewBg(imageFileName string) *bg {
	return &bg{image: core.LoadImage(imageFileName)}
}

func (b *bg) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(b.image, op)
}
