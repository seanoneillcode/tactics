package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	Name string
	Text string
}

func (l *Line) FullText() string {
	return l.Name + "\n" + l.Text
}

func (d *Dialog) Reset() {
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
	d.formattedText = GetFormattedValue(line.Text)
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
	if line.Name != d.currentName {
		d.order = d.order + 1
		if d.order == len(d.names) {
			d.order = 0
		}
	}
	d.currentName = line.Name
	d.formattedText = GetFormattedValue(line.Text)
}

func (d *Dialog) GetNameOrder() int {
	return d.order
}

func (d *Dialog) GetNextLinesForName() []string {
	var f []string
	for index, line := range d.lines {
		if index >= d.currentLineIndex {
			if d.currentName == line.Name {
				f = append(f, GetFormattedValue(line.FullText()))
			} else {
				break
			}
		}
	}
	return f
}

func (d *Dialog) GetCurrentLine() *Line {
	return &Line{
		Name: d.currentName,
		Text: d.bufferedText,
	}
}

func (d *Dialog) Update(delta int64, state *State) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if d.IsBuffering() {
			d.SkipBuffer()
		} else {
			if d.HasNextLine() {
				d.NextLine()
			} else {
				d.Reset()
				state.ActiveDialog = nil
			}
		}
	}

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

func GetFormattedValue(value string) string {
	return getFormattedValue(value, maxRunePerLine)
}

func GetFormattedValueMax(value string, max int) string {
	return getFormattedValue(value, max)
}

func getFormattedValue(value string, max int) string {
	var lines []string

	potentialLines := strings.Split(value, "\n")

	for _, potentialLine := range potentialLines {
		words := strings.Split(potentialLine, " ")

		var line []string
		count := 0
		for _, word := range words {

			if count+len(word) > max {
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
