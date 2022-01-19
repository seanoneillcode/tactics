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
	pos    *common.VectorF
	// movement
	isMoving  bool
	velocity  *common.VectorF
	moveTime  int64
	goalPos   *common.VectorF
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
			c.pos = c.goalPos
			c.isMoving = false
			c.velocity = &common.VectorF{}
		}
		c.pos = c.pos.Add(c.velocity.Mul(float64(delta)))
	}
	c.sprite.SetPosition(c.pos.X, c.pos.Y)
}

func (c *Character) TryToMove(dirX int, dirY int, state *State) {
	// check can move
	tileX, tileY := common.WorldToTile(c.pos)
	tileX = tileX + dirX
	tileY = tileY + dirY
	ti := state.Map.Level.GetTileInfo(tileX, tileY)
	if ti.tileData.isBlock {
		if ti.link == nil {
			return
		}
	}
	if ti.npc != nil {
		return
	}
	// perform move
	c.isMoving = true
	c.moveTime = characterMoveTime
	c.goalPos = c.pos.Add(common.VectorFromInt(dirX, dirY).Mul(common.TileSizeF))
	c.velocity = common.VectorFromInt(dirX, dirY).Mul(characterMoveAmount)
}

func (c *Character) SetPosition(pos *common.VectorF) {
	offset := c.goalPos.Sub(c.pos)
	c.goalPos = offset.Add(pos)
	c.pos = pos
}
