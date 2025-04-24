package cokeline

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Cokes      []string
	ActiveCoke int
	title      string

	width int
}

func InitialModel(width int, height int, cokes []string) Model {
	return Model{
		Cokes: cokes,

		width: getWidth(width),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width
}
