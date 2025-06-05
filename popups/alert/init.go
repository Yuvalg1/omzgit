package alert

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	error   string
	visible bool

	Width  int
	Height int
}

func InitialModel(width int, height int) Model {
	return Model{
		error:   "",
		visible: false,

		Width:  getWidth(width),
		Height: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getHeight(height int) int {
	return 5
}

func getWidth(width int) int {
	return min(34, width-2)
}

func (m Model) GetVisible() bool {
	return m.visible
}
