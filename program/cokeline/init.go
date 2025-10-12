package cokeline

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Left   string
	Center string
	Right  string

	Primary bool

	width int
}

func InitialModel(width int, height int, titles []string) Model {
	return Model{
		Left: "Files",

		width: getWidth(width),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width
}
