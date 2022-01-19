package dialog

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
	nd.currentDialog = nd.dialogs[nd.currentDialogName] // todo move this
	nd.currentDialog.Reset()
	return nd
}
