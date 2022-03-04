package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Pickup struct {
	sprite     *Sprite
	usedSprite *Sprite
	isUsed     bool
	itemName   string
}

func NewPickup(name string, itemName string, usedImageName string) *Pickup {
	p := &Pickup{
		sprite:   NewSprite(name + ".png"),
		itemName: itemName,
	}
	if usedImageName != "" {
		p.usedSprite = NewSprite(usedImageName + ".png")
	}
	return p
}

func (n *Pickup) SetPosition(x int, y int) {
	n.sprite.SetPosition(float64(x), float64(y))
	if n.usedSprite != nil {
		n.usedSprite.SetPosition(float64(x), float64(y))
	}
}

func (n *Pickup) GetPosition() *common.Vector {
	return &common.Vector{
		X: n.sprite.x,
		Y: n.sprite.y,
	}
}

func (n *Pickup) Draw(screen *ebiten.Image) {
	if !n.isUsed {
		n.sprite.Draw(screen)
	} else {
		if n.usedSprite != nil {
			n.usedSprite.Draw(screen)
		}
	}
}

func (n *Pickup) Update(delta int64, state *State) {
	// no op ? animate ?
}
