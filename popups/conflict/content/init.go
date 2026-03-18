package content

import (
	"strings"

	"omzgit/popups/conflict/chunk"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Content   viewport.Model
	ours      bool
	index     int
	conflicts []chunk.Model
	sum       int

	activeConflict int
}

func InitialModel(width int, height int, ours bool) Model {
	viewport := viewport.New(getWidth(width), getHeight(height))

	return Model{
		Content:   viewport,
		ours:      ours,
		conflicts: []chunk.Model{},
		index:     -1,

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
	sum := 0
	lines := 0

	for _, element := range m.conflicts {
		if element.Conflict {
			sum++
		}

		if m.index == sum-1 && element.Conflict {
			element.Active = true
			m.Content.SetYOffset(max(lines-3, 0))
		}
		content += element.View() + "\n"
		lines += strings.Count(lipgloss.NewStyle().Width(element.Width).Render(element.Content), "\n")
	}
	m.sum = sum
	m.Content.SetContent(content)
}
