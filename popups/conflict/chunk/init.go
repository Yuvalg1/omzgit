package chunk

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Conflict bool
	Active   bool
	Content  string
	ours     bool

	Width int
}

func InitialModel(conflict bool, ours bool, width int) Model {
	return Model{
		Conflict: conflict,
		Active:   false,
		Content:  "",
		ours:     ours,

		Width: getWidth(width),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) SetContent(content string) {
	m.Content = content
}

func (m *Model) Append(row string) {
	m.Content += row + "\n"
}

func getWidth(width int) int {
	return width
}
