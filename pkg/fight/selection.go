package fight

import "github.com/seanoneillcode/go-tactics/pkg/common"

type Selection struct {
	Pos    *common.Position
	Sprite *common.Sprite
}

func NewSelection() *Selection {
	return &Selection{
		Sprite: common.NewSprite("selection.png"),
		Pos:    common.PositionFromInt(0, 0),
	}
}

func (r Selection) Draw(camera *Camera) {
	r.Sprite.Draw(camera)
}

func (r *Selection) GetPos() *common.Position {
	return r.Pos
}

func (r *Selection) SetPos(x float64, y float64) {
	r.Pos.X = x
	r.Pos.Y = y
	r.Sprite.SetPosition(x, y)
}
