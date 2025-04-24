package branches

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	style := lipgloss.NewStyle().Width(m.width).Height(m.height)
	return style.Render(m.Title)
}
