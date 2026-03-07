package content

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Content  viewport.Model
	ours     bool
	conflict int
}

func InitialModel(width int, height int, ours bool) Model {
	viewport := viewport.New(getWidth(width), getHeight(height))

	return Model{
		Content:  viewport,
		ours:     ours,
		conflict: -1,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width - 2
}

func getHeight(height int) int {
	return height - 1
}

func (m *Model) SetContent(content string) {
	m.Content.SetContent(content)
}
