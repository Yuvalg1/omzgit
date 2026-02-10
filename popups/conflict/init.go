package conflict

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	width  int
	height int
}

func InitialModel(width int, height int) Model {
	return Model{
		width:  width,
		height: height,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return min(width-4, 40)
}

func getHeight(height int) int {
	return min(height-4, 20)
}
