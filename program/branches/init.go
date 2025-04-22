package branches

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	title  string
	width  int
	height int
}

func InitialModel(width int, height int) Model {
	return Model{
		title: "Branches",

		width:  width,
		height: height,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
