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
	isMoving  bool
	vx        float64
	vy        float64
	moveTime  int64
	goalX     float64
	goalY     float64
	lastInput string
}

func NewCharacter(imageFileName string) *Character {
	return &Character{
		sprite: NewSprite(imageFileName),
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
	c.sprite.SetPosition(c.x, c.y)
}

func (c *Character) TryToMove(dirX int, dirY int, state *State) {
	// check can move
	tileX, tileY := common.WorldToTile(c.x, c.y)
	tileX = tileX + dirX
	tileY = tileY + dirY
	ti := state.Level.GetTileInfo(tileX, tileY)
	if ti.tileData.isBlock {
		return
	}
	if len(ti.npcs) > 0 {
		return
	}
	// perform move
	c.isMoving = true
	c.moveTime = characterMoveTime
	c.goalX = c.x + float64(dirX*common.TileSize)
	c.goalY = c.y + float64(dirY*common.TileSize)
	c.vx = float64(dirX) * (characterMoveAmount)
	c.vy = float64(dirY) * (characterMoveAmount)
}
