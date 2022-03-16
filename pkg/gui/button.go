package gui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/core"
)

type button struct {
	pos            *Pos
	image          *ebiten.Image
	label          *Text
	isHighlighted  bool
	highlightColor color.RGBA
}

func NewButton(label string, imageFileName string, highlightColor color.RGBA) *button {
	b := &button{
		highlightColor: highlightColor,
		pos:            &Pos{},
	}
	if imageFileName != "" {
		b.image = core.LoadImage(imageFileName)
	}
	if label != "" {
		b.label = &Text{
			value: label,
			color: defaultTextColor,
		}
	}
	return b
}

func (b *button) Draw(screen *ebiten.Image) {
	if b.image != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(b.pos.X), float64(b.pos.Y))
		op.GeoM.Scale(common.Scale, common.Scale)
		if b.isHighlighted {
			cr, cg, cb, ca := b.highlightColor.RGBA()
			if ca == 0 {
				return
			}
			op.ColorM.Scale(float64(cr)/float64(ca), float64(cg)/float64(ca), float64(cb)/float64(ca), float64(ca)/0xffff)
		}
		screen.DrawImage(b.image, op)
	}
	if b.label != nil {
		b.label.SetPosition(*b.pos)
		b.label.Draw(screen)
	}
}

func (b *button) Update(delta int64, pos *Pos) {
	b.pos.X = pos.X
	b.pos.Y = pos.Y
}
