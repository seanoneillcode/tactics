package explore

import (
	"strings"
)

const maxRunePerLine = 40

type DialogHandler struct {
	ActiveDialog *DialogData
	IsActive     bool
}

func NewDialogHandler() *DialogHandler {
	return &DialogHandler{
		ActiveDialog: nil,
		IsActive:     false,
	}
}

func (r *DialogHandler) SetActiveDialog(dialog *DialogData) {
	r.ActiveDialog = dialog
	r.IsActive = true
}

func (r *DialogHandler) CloseDialog() {
	r.ActiveDialog = nil
	r.IsActive = false
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
