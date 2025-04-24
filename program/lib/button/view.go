package button

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#888888"))
	return style.Render(m.name)
}
