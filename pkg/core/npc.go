package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/go-tactics/pkg/common"
)

type Npc struct {
	character *Character
	npcDialog *NpcDialog
}

func NewNpc(name string) *Npc {
	return &Npc{
		character: NewCharacter(name + ".png"),
		npcDialog: GetNpcDialog(name),
	}
}

func (n *Npc) Draw(screen *ebiten.Image) {
	n.character.Draw(screen)
}

func (n *Npc) Update(delta int64, state *State) {
	n.character.Update(delta)
}

func (n *Npc) SetPosition(x int, y int) {
	n.character.pos = common.VectorFromInt(x, y)
}

func (n *Npc) GetPosition() *common.Vector {
	return n.character.pos
}

func (n *Npc) GetCurrentDialog() *Dialog {
	return n.npcDialog.GetCurrentDialog()
}
