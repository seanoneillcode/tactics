package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Player struct {
	isActive  bool
	Character *Character
	lastInput string

	ActiveShop *ShopData
	TeamState  *TeamState
	// not sure about this
	isSleeping  bool
	playerState string
}

func NewPlayer() *Player {
	p := &Player{
		isActive:  true,
		Character: NewCharacter("player.png"),
		TeamState: NewTeamState(),
	}
	p.TeamState.Characters = append(p.TeamState.Characters, NewCharacterState())
	return p
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Character.Draw(screen)
}

func (p *Player) Update(delta int64, state *State) {
	p.Character.Update(delta)

	if !p.isActive {
		return
	}
	if state.ActiveDialog != nil {
		return
	}
	if !p.Character.isMoving {
		var inputX = 0
		var inputY = 0
		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
			inputX = inputX - 1
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
			inputX = inputX + 1
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
			inputY = inputY - 1
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
			inputY = inputY + 1
		}
		if inputX != 0 && inputY != 0 {
			if p.lastInput == "x" {
				inputX = 0
			} else {
				inputY = 0
			}
		}
		if inputX != 0 || inputY != 0 {
			if inputX != 0 {
				p.lastInput = "x"
			}
			if inputY != 0 {
				p.lastInput = "y"
			}
			p.Character.TryToMove(inputX, inputY, state)

			tileX, tileY := common.WorldToTile(p.Character.pos)
			ti := state.Map.Level.GetTileInfo(inputX+tileX, tileY+inputY)

			// check for things on the tile
			if ti.npc != nil {
				state.ActiveDialog = ti.npc.GetCurrentDialog()
			}
			if ti.link != nil {
				state.Map.StartTransition(ti.link)
			}
			if ti.pickup != nil && !ti.pickup.isUsed {
				ti.pickup.isUsed = true
				p.TeamState.Pickup(ti.pickup)
			}
			if ti.action != nil {
				if ti.action.name == "bed" {
					p.SetSleep(true)
					p.TeamState.RestoreHealth()
				}
			}
			if ti.shop != nil {
				// pause everything else and open 'shop mode'
				state.Shop.Open(ti.shop)
				p.isActive = false
			}
		}
	}

	// consider this key more carefully
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.Inventory.Open(p.TeamState)
		p.isActive = false
	}
}

func (p *Player) Activate() {
	p.isActive = true
}

func (p *Player) EnterLevel(level *Level) {
	for _, link := range level.links {
		p.Character.pos = common.VectorFromInt(link.x, link.y)
		return
	}
}

func (p *Player) SetPosition(pos *common.Vector) {
	p.Character.SetPosition(pos)
}

func (p *Player) SetSleep(b bool) {
	p.isSleeping = b
}
