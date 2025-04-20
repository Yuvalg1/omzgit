package row

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	descriptor byte
	text       string

	Width int
}

func InitialModel(text string, isDesc bool, width int) Model {
	return Model{
		descriptor: getDescriptor(text, isDesc),
		text:       text,

		Width: width,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getDescriptor(text string, isDesc bool) byte {
	if isDesc {
		return 'D'
	}

	if strings.HasPrefix(text, "@@") {
		return '@'
	}

	if strings.HasPrefix(text, "+") {
		return '+'
	}

	if strings.HasPrefix(text, "-") {
		return '-'
	}

	return ' '
}
