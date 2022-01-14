package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/dialog"
)

type Player struct {
	character    *Character
	lastInput    string
	ActiveDialog *dialog.NpcDialog
}

func NewPlayer() *Player {
	return &Player{
		character: NewCharacter("player.png"),
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.character.Draw(screen)
}

func (p *Player) Update(delta int64, state *State) {
	p.character.Update(delta)

	if p.ActiveDialog != nil {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			isDone := p.ActiveDialog.NextLine()
			if isDone {
				p.ActiveDialog = nil
			}
		}
		return
	}

	if !p.character.isMoving {
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
			p.character.TryToMove(inputX, inputY, state)

			// check for dialogs
			tileX, tileY := common.WorldToTile(p.character.x, p.character.y)
			ti := state.Map.Level.GetTileInfo(inputX+tileX, tileY+inputY)
			if ti.npc != nil {
				p.ActiveDialog = ti.npc.GetNpcDialog()
			}
			if ti.link != nil {
				state.Map.StartTransition(ti.link)
			}
		}
	}
}

func (p *Player) EnterLevel(level *Level) {
	for _, link := range level.links {
		p.character.x = float64(link.x)
		p.character.y = float64(link.y)
		return
	}
}

func (p *Player) SetPosition(x float64, y float64) {
	p.character.SetPosition(x, y)
}
