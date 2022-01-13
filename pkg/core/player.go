package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	character *Character
	lastInput string
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
		}
	}
}
