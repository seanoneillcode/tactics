package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

const (
	characterMoveTime   = 200.0 // milliseconds
	characterMoveAmount = common.TileSize / characterMoveTime
)

type Character struct {
	sprite *Sprite
	x      float64
	y      float64
	// movement
	isMoving bool
	vx       float64
	vy       float64
	moveTime int64
	goalX    float64
	goalY    float64
}

func NewCharacter() *Character {
	return &Character{
		sprite: NewSprite(),
	}
}

func (c *Character) Draw(screen *ebiten.Image) {
	c.sprite.Draw(screen)
}

func (c *Character) Update(delta int64) {
	if c.isMoving {
		c.moveTime = c.moveTime - delta
		if c.moveTime < 0 {
			c.x = c.goalX
			c.y = c.goalY
			c.isMoving = false
			c.vx = 0
			c.vy = 0
		}
		c.x = c.x + (c.vx * float64(delta))
		c.y = c.y + (c.vy * float64(delta))
	}
	if !c.isMoving {
		var inputX = 0
		var inputY = 0

		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			inputX = inputX - 1
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			inputX = inputX + 1
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
			inputY = inputY - 1
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
			inputY = inputY + 1
		}
		if inputX != 0 || inputY != 0 {
			c.isMoving = true
			c.moveTime = characterMoveTime
			c.goalX = c.x + float64(inputX*common.TileSize)
			c.goalY = c.y + float64(inputY*common.TileSize)
			c.vx = float64(inputX) * (characterMoveAmount)
			c.vy = float64(inputY) * (characterMoveAmount)
		}
	}
	c.sprite.x = c.x
	c.sprite.y = c.y
}
