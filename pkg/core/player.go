package core

import (
	"fmt"
	"github.com/seanoneillcode/go-tactics/pkg/input"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Player struct {
	isActive  bool
	Character *Character
	lastInput string

	ActiveShop *ShopData

	// not sure about this
	isSleeping  bool
	playerState string
}

func NewPlayer() *Player {
	p := &Player{
		isActive:  true,
		Character: NewCharacter("player.png"),
	}
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
	if state.DialogHandler.IsActive {
		return
	}
	if !p.Character.isMoving {
		var inputX = 0
		var inputY = 0
		if input.IsLeftPressed() {
			inputX = inputX - 1
		}
		if input.IsRightPressed() {
			inputX = inputX + 1
		}
		if input.IsUpPressed() {
			inputY = inputY - 1
		}
		if input.IsDownPressed() {
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
				state.DialogHandler.SetActiveDialog(ti.npc.GetCurrentDialog())
			}
			if ti.link != nil {
				state.Map.StartTransition(ti.link)
			}
			if ti.pickup != nil && !ti.pickup.isUsed {
				ti.pickup.isUsed = true
				state.TeamState.Pickup(ti.pickup)
				state.DialogHandler.SetActiveDialog(getPickupDialog(ti.pickup.itemName))
			}
			if ti.action != nil {
				if ti.action.name == "bed" {
					p.SetSleep(true)
					state.TeamState.RestoreHealth()
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
	if input.IsMenuPressed() {
		state.UI.Open(MenuUI)
		p.isActive = false
	}
}

func (p *Player) Activate() {
	p.isActive = true
}

func (p *Player) EnterLevel(level *Level) {
	for _, link := range level.links {
		p.Character.pos = common.PositionFromInt(link.x, link.y)
		return
	}
}

func (p *Player) SetPosition(pos *common.Position) {
	p.Character.SetPosition(pos)
}

func (p *Player) SetSleep(b bool) {
	p.isSleeping = b
}

func getPickupDialog(name string) *DialogData {
	d := &DialogData{
		Lines: []*Line{
			{
				Text: fmt.Sprintf("picked up a '%v'", name),
			},
		},
	}
	return d
}
