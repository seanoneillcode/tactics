package elem

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

const offsetX = 4
const offsetY = 4

type Button struct {
	pos        *Pos
	image      *ebiten.Image
	label      *Text
	imageColor *color.RGBA
	textColor  *color.RGBA
}

var highlightColor = &color.RGBA{
	R: 223,
	G: 246,
	B: 245,
	A: 255,
}

var greyColor = &color.RGBA{
	R: 160,
	G: 147,
	B: 142,
	A: 255,
}

func NewButton(label string, imageFileName string) *Button {
	b := &Button{
		pos:        &Pos{},
		imageColor: highlightColor,
		textColor:  highlightColor,
	}
	if imageFileName != "" {
		b.image = common.LoadImage(imageFileName)
	}
	if label != "" {
		b.label = NewText(0, 0, label)
	}
	return b
}

func (b *Button) Draw(screen *ebiten.Image) {
	if b.image != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(b.pos.X), float64(b.pos.Y))
		op.GeoM.Scale(common.Scale, common.Scale)
		cr, cg, cb, ca := b.imageColor.RGBA()
		if ca == 0 {
			return
		}
		op.ColorM.Scale(float64(cr)/float64(ca), float64(cg)/float64(ca), float64(cb)/float64(ca), float64(ca)/0xffff)
		screen.DrawImage(b.image, op)
	}
	if b.label != nil {
		b.label.SetPosition(Pos{X: b.pos.X + 2 + offsetX, Y: b.pos.Y + offsetY})
		b.label.SetColor(*b.textColor)
		b.label.Draw(screen)
	}
}

func (b *Button) Update(delta int64, pos *Pos, isHighlighted bool) {
	b.pos.X = pos.X
	b.pos.Y = pos.Y
	if isHighlighted {
		b.textColor = highlightColor
		b.imageColor = highlightColor
	} else {
		b.textColor = greyColor
		b.imageColor = greyColor
	}
}
