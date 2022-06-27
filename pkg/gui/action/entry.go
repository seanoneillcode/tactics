package action

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/gui/elem"
)

type Entry struct {
	bg            *elem.Sprite
	text          *elem.Text
	isHighlighted bool
}

func NewEntry(msg string) *Entry {
	return &Entry{
		text: elem.NewText(0, 0, msg),
		bg:   elem.NewSprite("uis/fight/entry-bg.png", 0, 0),
	}
}

func (r *Entry) Draw(screen *ebiten.Image) {
	r.bg.Draw(screen)
	r.text.Draw(screen)
}

func (r *Entry) Update(pos elem.Pos) {
	r.bg.SetPos(pos)
	r.text.SetPosition(elem.Pos{
		X: pos.X + 2,
		Y: pos.Y + 2,
	})
}
