package button

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	name string
}

func InitialModel(name string) Model {
	return Model{
		name: name,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
