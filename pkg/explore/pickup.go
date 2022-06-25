package explore

import (
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Pickup struct {
	sprite     *common.Sprite
	usedSprite *common.Sprite
	isUsed     bool
	itemName   string
}

func NewPickup(name string, itemName string, usedImageName string) *Pickup {
	p := &Pickup{
		sprite:   common.NewSprite(name + ".png"),
		itemName: itemName,
	}
	if usedImageName != "" {
		p.usedSprite = common.NewSprite(usedImageName + ".png")
	}
	return p
}

func (n *Pickup) SetPosition(x int, y int) {
	n.sprite.SetPosition(float64(x), float64(y))
	if n.usedSprite != nil {
		n.usedSprite.SetPosition(float64(x), float64(y))
	}
}

func (n *Pickup) GetPosition() *common.Position {
	return &common.Position{
		X: n.sprite.X,
		Y: n.sprite.Y,
	}
}

func (n *Pickup) Draw(camera *Camera) {
	if !n.isUsed {
		n.sprite.Draw(camera)
	} else {
		if n.usedSprite != nil {
			n.usedSprite.Draw(camera)
		}
	}
}

func (n *Pickup) Update(delta int64, state *State) {
	// no op ? animate ?
}
