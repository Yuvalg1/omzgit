package branches

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Title  string
	width  int
	height int
}

func InitialModel(width int, height int, title string) Model {
	return Model{
		Title: title,

		width:  getWidth(width),
		height: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width
}

func getHeight(height int) int {
	return height
}
