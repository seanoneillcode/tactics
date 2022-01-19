package dialog

import (
	"strings"
)

const letterSpeed = 20
const maxRunePerLine = 40

type Dialog struct {
	currentLineIndex int
	lines            []*Line
	index            int
	timer            int64
	bufferedText     string
	formattedText    string
	currentName      string
	order            int
	names            []string
}

type Line struct {
	name string
	text string
}

func (d *Dialog) Reset() {
	d.currentLineIndex = 0
	d.bufferedText = ""
	d.index = 0
	d.timer = 0
	d.order = 0
	names := map[string]bool{}
	for _, line := range d.lines {
		if ok, _ := names[line.name]; !ok {
			names[line.name] = true
		}
	}
	d.names = []string{}
	for name := range names {
		d.names = append(d.names, name)
	}
	line := d.lines[d.currentLineIndex]
	d.currentName = line.name
	d.formattedText = getFormattedValue(line.text)
}

func (d *Dialog) IsBuffering() bool {
	return d.bufferedText != d.formattedText
}

func (d *Dialog) SkipBuffer() {
	d.bufferedText = d.formattedText
}

func (d *Dialog) HasNextLine() bool {
	return d.currentLineIndex+1 < len(d.lines)
}

// NextLine moves to the next line and returns true if the dialog is done
func (d *Dialog) NextLine() {
	d.currentLineIndex = d.currentLineIndex + 1
	d.bufferedText = ""
	d.index = 0
	d.timer = 0
	line := d.lines[d.currentLineIndex]
	if line.name != d.currentName {
		d.order = d.order + 1
		if d.order == len(d.names) {
			d.order = 0
		}
	}
	d.currentName = line.name
	d.formattedText = getFormattedValue(line.text)
}

func (d *Dialog) GetNameOrder() int {
	return d.order
}

func (d *Dialog) GetAllFormattedLines() []string {
	var f []string
	for _, line := range d.lines {
		f = append(f, getFormattedValue(line.text))
	}
	return f
}

func (d *Dialog) GetCurrentText() string {
	return d.bufferedText
}

func (d *Dialog) Update(delta int64) {
	if d.bufferedText != d.formattedText {
		d.timer = d.timer + delta
		if d.timer > letterSpeed {
			d.timer = d.timer - letterSpeed
			d.index = d.index + 1
			if d.index == len(d.formattedText)-1 {
				d.bufferedText = d.formattedText
			} else {
				d.bufferedText = d.formattedText[:d.index]
			}
		}
	}
}

func getFormattedValue(value string) string {
	var lines []string

	potentialLines := strings.Split(value, "\n")

	for _, potentialLine := range potentialLines {
		words := strings.Split(potentialLine, " ")

		var line []string
		count := 0
		for _, word := range words {

			if count+len(word) > maxRunePerLine {
				compoundLine := strings.Join(line, " ")
				lines = append(lines, compoundLine)
				line = []string{}
				count = 0
			}
			count = count + len(word) + 1
			line = append(line, word)

		}
		lines = append(lines, strings.Join(line, " "))
	}

	return strings.Join(lines, "\n")
}
