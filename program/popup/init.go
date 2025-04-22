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

		Width:  GetWidth(width),
		Height: GetHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func GetHeight(height int) int {
	return height - 2
}

func GetWidth(width int) int {
	return width - 2
}
