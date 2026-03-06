package conflict

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	visible bool

	ours   string
	theirs string

	width  int
	height int
}

func InitialModel(width int, height int) Model {
	return Model{
		visible: false,

		width:  width,
		height: height,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) GetVisible() bool {
	return m.visible
}

func getContent(path string) string {
	return path + "apt apt"
}

func getWidth(width int) int {
	return width
}

func getHeight(height int) int {
	return height - 2
}
