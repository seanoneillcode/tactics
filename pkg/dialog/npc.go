package dialog

func (d *Dialog) Reset() {
	d.currentLineIndex = 0
}

func (d *Dialog) HasNextLine() bool {
	return d.currentLineIndex+1 < len(d.lines)
}

// NextLine moves to the next line and returns true if the dialog is done
func (d *Dialog) NextLine() {
	d.currentLineIndex = d.currentLineIndex + 1
}

func (d *Dialog) GetCurrentLine() (string, string) {
	line := d.lines[d.currentLineIndex]
	return line.name, line.text
}

type NpcDialog struct {
	currentDialogName string
	currentDialog     *Dialog
	dialogs           map[string]*Dialog
}

func (nd NpcDialog) GetCurrentDialog() *Dialog {
	return nd.currentDialog
}

type Dialog struct {
	currentLineIndex int
	lines            []*Line
}

type Line struct {
	name string
	text string
}

func GetNpcDialog(name string) *NpcDialog {
	nd := dialogData[name]
	nd.currentDialog = nd.dialogs[nd.currentDialogName] // todo move this
	return nd
}
