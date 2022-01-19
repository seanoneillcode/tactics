package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/dialog"
)

type Npc struct {
	character *Character
	npcDialog *dialog.NpcDialog
}

func NewNpc(name string) *Npc {
	return &Npc{
		character: NewCharacter(name + ".png"),
		npcDialog: dialog.GetNpcDialog(name),
	}
}

func (n *Npc) Draw(screen *ebiten.Image) {
	n.character.Draw(screen)
}

func (n *Npc) Update(delta int64, state *State) {
	n.character.Update(delta)
}

func (n *Npc) SetPosition(x int, y int) {
	n.character.x = float64(x)
	n.character.y = float64(y)
}

func (n *Npc) GetPosition() (float64, float64) {
	return n.character.x, n.character.y
}

func (n *Npc) GetCurrentDialog() *dialog.Dialog {
	return n.npcDialog.GetCurrentDialog()
}
