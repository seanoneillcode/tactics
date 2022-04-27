package dialog

import (
	"github.com/seanoneillcode/go-tactics/pkg/core"
)

const letterSpeed = 40

type dialogState struct {
	currentLineIndex int
	lines            []*lineState
	index            int
	timer            int64
	bufferedText     string
	formattedText    string
	currentName      string
	order            int
	names            []string
}

type lineState struct {
	Name     string
	Text     string
	FullText string
}

func NewDialogState(data *core.DialogData) *dialogState {
	d := &dialogState{}
	var lines []*lineState
	for _, l := range data.Lines {
		lines = append(lines, &lineState{
			Name:     l.Name,
			Text:     l.Text,
			FullText: fullText(l.Name, l.Text),
		})
	}
	d.lines = lines
	names := map[string]bool{}
	for _, line := range d.lines {
		if ok, _ := names[line.Name]; !ok {
			names[line.Name] = true
		}
	}
	d.names = []string{}
	for name := range names {
		d.names = append(d.names, name)
	}
	line := d.lines[d.currentLineIndex]
	d.currentName = line.Name
	d.formattedText = core.GetFormattedValue(line.FullText)
	return d
}

func (d *dialogState) Reset() {
	d.currentLineIndex = 0
	d.bufferedText = ""
	d.index = 0
	d.timer = 0
	d.order = 0
	names := map[string]bool{}
	for _, line := range d.lines {
		if ok, _ := names[line.Name]; !ok {
			names[line.Name] = true
		}
	}
	d.names = []string{}
	for name := range names {
		d.names = append(d.names, name)
	}
	line := d.lines[d.currentLineIndex]
	d.currentName = line.Name
	d.formattedText = core.GetFormattedValue(line.FullText)
}

func (d *dialogState) IsBuffering() bool {
	return d.bufferedText != d.formattedText
}

func (d *dialogState) SkipBuffer() {
	d.bufferedText = d.formattedText
}

func (d *dialogState) HasNextLine() bool {
	return d.currentLineIndex+1 < len(d.lines)
}

// NextLine moves to the next line and returns true if the dialog is done
func (d *dialogState) NextLine() {
	d.currentLineIndex = d.currentLineIndex + 1
	d.bufferedText = ""
	d.index = 0
	d.timer = 0
	l := d.lines[d.currentLineIndex]
	if l.Name != d.currentName {
		d.order = d.order + 1
		if d.order == len(d.names) {
			d.order = 0
		}
	}
	d.currentName = l.Name
	d.formattedText = core.GetFormattedValue(l.FullText)
}

func (d *dialogState) GetNameOrder() int {
	return d.order
}

func (d *dialogState) GetNextLinesForName() []string {
	var f []string
	for index, line := range d.lines {
		if index >= d.currentLineIndex {
			if d.currentName == line.Name {
				f = append(f, core.GetFormattedValue(line.FullText))
			} else {
				break
			}
		}
	}
	return f
}

func (d *dialogState) GetCurrentLine() (string, string) {
	return d.currentName, d.bufferedText
}

func fullText(name string, text string) string {
	return name + ":\n" + text
}

func (d *dialogState) Update(delta int64) {
	if d.bufferedText != d.formattedText {
		d.timer = d.timer + delta
		for d.timer > letterSpeed {
			d.timer = d.timer - letterSpeed
			d.index = d.index + 1
			if d.index == len(d.formattedText)-1 {
				d.bufferedText = d.formattedText
				break
			} else {
				d.bufferedText = d.formattedText[:d.index]
			}
		}
	}
}
