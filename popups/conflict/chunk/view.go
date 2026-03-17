package chunk

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	return lipgloss.NewStyle().
		Width(m.width).
		Render(m.content)
}
