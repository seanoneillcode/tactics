package core

import "log"

type NpcDialog struct {
	currentDialogName string
	currentDialog     *DialogData
	dialogs           map[string]*DialogData
}

func (nd NpcDialog) GetCurrentDialog() *DialogData {
	return nd.currentDialog
}

func GetNpcDialog(name string) *NpcDialog {
	nd := dialogData[name]
	if nd == nil {
		log.Fatalf("missing dialog for npc: %s", name)
	}
	nd.currentDialog = nd.dialogs[nd.currentDialogName] // todo move this
	return nd
}
