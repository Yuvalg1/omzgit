package row

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	descriptor byte
	text       string

	width  int
	height int
}

func InitialModel(text string, isDesc bool, width int) Model {
	return Model{
		descriptor: getDescriptor(text, isDesc),
		text:       text,

		width:  GetWidth(width),
		height: GetWidth(1),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func GetWidth(width int) int {
	return width - 1
}

func GetHeight(height int) int {
	return 2
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
