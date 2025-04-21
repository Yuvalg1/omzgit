package name

import (
	"program/consts"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), false, true, true, true).Padding(0, 1).MaxWidth(m.width - 2).Height(1).Bold(true)

	return lipgloss.PlaceHorizontal(m.width, lipgloss.Center, style.Render(consts.TrimRight(m.Name, m.width-6)))
}
