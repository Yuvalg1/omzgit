package branches

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	style := lipgloss.NewStyle().Width(m.width).Height(m.height)
	var branches string
	for _, element := range m.branches {
		branches += "\n" + element.View()
	}

	return style.Render(m.Title + branches)
}
