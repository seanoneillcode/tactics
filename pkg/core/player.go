package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/dialog"
	"log"
)

type Player struct {
	Character      *Character
	lastInput      string
	ActiveDialog   *dialog.Dialog
	CharacterState *CharacterState
	// not sure about this
	isSleeping bool
}

func NewPlayer() *Player {
	return &Player{
		Character:      NewCharacter("player.png"),
		CharacterState: NewCharacterState(),
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Character.Draw(screen)
}

func (p *Player) Update(delta int64, state *State) {
	p.Character.Update(delta)

	if p.ActiveDialog != nil {
		p.ActiveDialog.Update(delta)
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			if p.ActiveDialog.IsBuffering() {
				p.ActiveDialog.SkipBuffer()
			} else {
				if p.ActiveDialog.HasNextLine() {
					p.ActiveDialog.NextLine()
				} else {
					p.ActiveDialog.Reset()
					p.ActiveDialog = nil
				}
			}
		}
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

			// check for dialogs
			if ti.npc != nil {
				p.ActiveDialog = ti.npc.GetCurrentDialog()
			}
			// check for links
			if ti.link != nil {
				state.Map.StartTransition(ti.link)
			}
			// check for pickups
			if ti.pickup != nil && !ti.pickup.isUsed {
				ti.pickup.isUsed = true
				p.Pickup(ti.pickup)
			}
			// check for actions
			if ti.action != nil {
				if ti.action.name == "bed" {
					p.SetSleep(true)
					p.CharacterState.Health = p.CharacterState.MaxHealth
				}
			}
		}
	}

}

func (p *Player) EnterLevel(level *Level) {
	for _, link := range level.links {
		p.Character.pos = common.VectorFromInt(link.x, link.y)
		return
	}
}

func (p *Player) SetPosition(pos *common.VectorF) {
	p.Character.SetPosition(pos)
}

func (p *Player) Pickup(pickup *Pickup) {
	p.CharacterState.Items = append(p.CharacterState.Items, NewItem(pickup.itemName))
	log.Printf("picked up %v", pickup.itemName)
}

func (p *Player) SetSleep(b bool) {
	p.isSleeping = b
}
