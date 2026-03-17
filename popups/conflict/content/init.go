package content

import (
	"omzgit/popups/conflict/chunk"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Content   viewport.Model
	ours      bool
	conflicts []chunk.Model

	activeConflict int
}

func InitialModel(width int, height int, ours bool) Model {
	viewport := viewport.New(getWidth(width), getHeight(height))

	return Model{
		Content:   viewport,
		ours:      ours,
		conflicts: []chunk.Model{},

		activeConflict: 0,
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

func (m *Model) Append(chunk chunk.Model) {
	m.conflicts = append(m.conflicts, chunk)
}

func (m *Model) Refresh() {
	content := ""

	for _, element := range m.conflicts {
		if element.View() != "" {
			content += element.View()
		}
	}
	m.Content.SetContent(content)
}
