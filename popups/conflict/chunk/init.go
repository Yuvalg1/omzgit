package chunk

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Conflict bool
	Active   bool
	content  string
	ours     bool

	width int
}

func InitialModel(conflict bool, ours bool, width int) Model {
	return Model{
		Conflict: conflict,
		Active:   false,
		content:  "",
		ours:     ours,

		width: getWidth(width),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) SetContent(content string) {
	m.content = content
}

func (m *Model) Append(row string) {
	style := lipgloss.NewStyle().Width(m.width)

	m.content += style.Render(row) + "\n"
}

func getWidth(width int) int {
	return width
}
