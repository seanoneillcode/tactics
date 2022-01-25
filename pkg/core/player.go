package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/seanoneillcode/go-tactics/pkg/common"
	"github.com/seanoneillcode/go-tactics/pkg/dialog"
)

type Player struct {
	Character    *Character
	lastInput    string
	ActiveDialog *dialog.Dialog
}

func NewPlayer() *Player {
	return &Player{
		Character: NewCharacter("player.png"),
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

			// check for dialogs
			tileX, tileY := common.WorldToTile(p.Character.pos)
			ti := state.Map.Level.GetTileInfo(inputX+tileX, tileY+inputY)
			if ti.npc != nil {
				p.ActiveDialog = ti.npc.GetCurrentDialog()
			}
			if ti.link != nil {
				state.Map.StartTransition(ti.link)
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
