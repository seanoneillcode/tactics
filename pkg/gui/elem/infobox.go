package elem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type InfoBox struct {
	pos    *Pos
	image  *ebiten.Image
	label  *Text
	isDraw bool
}

func NewInfoBox(label string, imageFileName string) *InfoBox {
	r := &InfoBox{
		pos: &Pos{},
	}
	if imageFileName != "" {
		r.image = common.LoadImage(imageFileName)
	}
	r.label = NewText(0, 0, label)
	return r
}

func (r *InfoBox) Draw(screen *ebiten.Image) {
	if !r.isDraw {
		return
	}
	if r.image != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(r.pos.X), float64(r.pos.Y))
		op.GeoM.Scale(common.Scale, common.Scale)
		screen.DrawImage(r.image, op)
	}
	if r.label != nil {
		r.label.SetPosition(Pos{X: r.pos.X + 2 + offsetX, Y: r.pos.Y + offsetY})
		r.label.Draw(screen)
	}
}

func (r *InfoBox) Update(pos *Pos, isDraw bool, label string) {
	r.pos.X = pos.X
	r.pos.Y = pos.Y
	r.isDraw = isDraw
	r.label.SetValue(label)
}
