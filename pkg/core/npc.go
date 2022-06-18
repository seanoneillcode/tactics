package core

import (
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

func (n *Npc) Draw(camera *Camera) {
	n.character.Draw(camera)
}

func (n *Npc) Update(delta int64, state *State) {
	n.character.Update(delta)
}

func (n *Npc) SetPosition(x int, y int) {
	n.character.pos = common.PositionFromInt(x, y)
}

func (n *Npc) GetPosition() *common.Position {
	return n.character.pos
}

func (n *Npc) GetCurrentDialog() *DialogData {
	return n.npcDialog.GetCurrentDialog()
}
