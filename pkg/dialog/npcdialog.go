package dialog

// NextLine moves to the next line and returns true if the dialog is done
func (n *NpcDialog) NextLine() bool {
	cd := n.dialogs[n.currentDialog]

	// if we reached the end
	if cd.currentLineIndex+1 == len(cd.lines) {
		cd.currentLineIndex = 0
		return true
	}

	cd.currentLineIndex = cd.currentLineIndex + 1
	return false
}

func (n *NpcDialog) GetCurrentLine() (string, string) {
	cd := n.dialogs[n.currentDialog]
	line := cd.lines[cd.currentLineIndex]
	return line.name, line.text
}

type NpcDialog struct {
	currentDialog string
	dialogs       map[string]*Dialog
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
	return dialogData[name]
}

// the map is organized by npc name
var dialogData = map[string]*NpcDialog{
	"dave": {
		// each dialog has a key, this can be changed by events
		// i.e. player kills a king -> change to king based context dialog
		dialogs: map[string]*Dialog{
			"": {
				lines: []*Line{
					{
						name: "dave",
						text: "my name is dave.",
					},
					{
						name: "dave",
						text: "welcome to the game!",
					},
				},
			},
		},
	},
}
