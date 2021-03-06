package elem

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/explore"
	"strings"
)

type CharacterCard struct {
	pos    *Pos
	image  *ebiten.Image
	bg     *ebiten.Image
	name   *Text
	health *Text
	magic  *Text
	level  *Text
}

func NewCharacterCard(cs *explore.CharacterState, pos Pos) *CharacterCard {
	r := &CharacterCard{
		pos:    &Pos{pos.X, pos.Y},
		image:  common.LoadImage(fmt.Sprintf("portrait/%s.png", cs.Name)),
		bg:     common.LoadImage("elem/card-bg.png"),
		name:   NewText(pos.X+offsetX-2, pos.Y+offsetY+34, strings.Title(cs.Name)),
		health: NewText(pos.X+offsetX+48+8, pos.Y+offsetY+16, fmt.Sprintf("health: %v", cs.Health)),
		magic:  NewText(pos.X+offsetX+48+8, pos.Y+offsetY+32, fmt.Sprintf("magic: %v", cs.Magic)),
		level:  NewText(pos.X+offsetX+144, pos.Y+offsetY, fmt.Sprintf("level: %v", cs.BaseStats.Level)),
	}
	return r
}

func (r *CharacterCard) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.pos.X+48), float64(r.pos.Y))
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(r.bg, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.pos.X), float64(r.pos.Y))
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(r.image, op)

	r.name.Draw(screen)
	r.health.Draw(screen)
	r.magic.Draw(screen)
	r.level.Draw(screen)
}

func (r *CharacterCard) Update(pos *Pos) {
	r.pos.X = pos.X
	r.pos.Y = pos.Y
}
