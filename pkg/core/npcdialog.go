package core

import "log"

type NpcDialog struct {
	currentDialogName string
	currentDialog     *Dialog
	dialogs           map[string]*Dialog
}

func (nd NpcDialog) GetCurrentDialog() *Dialog {
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
