package popup

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Fn      func()
	Name    string
	Visible bool

	Width  int
	Height int
}

func InitialModel(fn func(), name string, width int, height int) Model {
	return Model{
		Fn:      fn,
		Name:    name,
		Visible: false,

		Width:  getWidth(width),
		Height: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getHeight(height int) int {
	return height
}

func getWidth(width int) int {
	return width
}
