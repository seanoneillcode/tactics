package explore

import (
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

const (
	characterMoveTime   = 250.0 // milliseconds
	characterMoveAmount = common.TileSize / characterMoveTime
)

type Character struct {
	sprite *common.Sprite
	pos    *common.Position
	// movement
	isMoving  bool
	velocity  *common.Position
	moveTime  int64
	goalPos   *common.Position
	lastInput string
	Direction *common.Direction
	// animation
	isAFrame bool
}

func NewCharacter(imageFileName string) *Character {
	return &Character{
		sprite: common.NewSprite(imageFileName),
		Direction: &common.Direction{
			X: 0,
			Y: 1,
		},
	}
}

func (c *Character) Draw(camera *Camera) {
	if !c.isMoving {
		if c.Direction.Y == -1 {
			c.sprite.SetFrame(1)
		}
		if c.Direction.Y == 1 {
			c.sprite.SetFrame(0)
		}
		if c.Direction.X == -1 {
			c.sprite.SetFrame(2)
		}
		if c.Direction.X == 1 {
			c.sprite.SetFrame(7)
		}
	} else {
		var frameA int
		var frameB int
		if c.Direction.Y == 1 {
			frameA = 4
			frameB = 5
		}
		if c.Direction.Y == -1 {
			frameA = 8
			frameB = 9
		}
		if c.Direction.X == 1 {
			frameA = 6
			frameB = 7
		}
		if c.Direction.X == -1 {
			frameA = 2
			frameB = 3
		}
		if c.isAFrame {
			c.sprite.SetFrame(frameA)
		} else {
			c.sprite.SetFrame(frameB)
		}
	}
	c.sprite.Draw(camera)
}

func (c *Character) Update(delta int64) {
	if c.isMoving {
		c.moveTime = c.moveTime - delta
		if c.moveTime < 0 {
			c.pos = c.goalPos
			c.isMoving = false
			c.velocity = &common.Position{}
		}
		c.pos = c.pos.Add(c.velocity.Mul(float64(delta)))

		c.isAFrame = c.moveTime > (characterMoveTime / 2)
	}
	c.sprite.SetPosition(c.pos.X, c.pos.Y)
}

func (c *Character) TryToMove(dirX int, dirY int, state *State) {
	// check can move
	c.Direction = &common.Direction{
		X: dirX,
		Y: dirY,
	}

	tile := common.WorldToTile(c.pos)
	tile.X = tile.X + dirX
	tile.Y = tile.Y + dirY
	ti := state.Map.Level.GetTileInfo(tile)
	if ti.tileData.IsBlock {
		if ti.link == nil {
			return
		}
	}
	if ti.npc != nil {
		return
	}
	if ti.enemy != nil {
		return
	}
	if ti.pickup != nil {
		return
	}
	// perform move
	c.isAFrame = !c.isAFrame
	c.isMoving = true
	c.moveTime = characterMoveTime
	c.goalPos = c.pos.Add(common.PositionFromDirection(c.Direction).Mul(common.TileSizeF))
	c.velocity = common.PositionFromDirection(c.Direction).Mul(characterMoveAmount)
}

func (c *Character) SetPosition(pos *common.Position) {
	offset := c.goalPos.Sub(c.pos)
	c.goalPos = offset.Add(pos)
	c.pos = pos
}
