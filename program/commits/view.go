package commits

import (
	"program/border"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	style := lipgloss.NewStyle().Border(lipgloss.DoubleBorder(), false, true, true).Width(m.width - 2).Height(m.height - 2)
	return border.GetTopBorder(m.title, m.width-2) + style.Render(m.title)
}
