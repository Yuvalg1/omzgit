package files

import (
	"github.com/charmbracelet/lipgloss"
)

var leftText = "Files"

func (m Model) View() string {
	style := lipgloss.NewStyle().Width(m.Width)

	filesStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(lipgloss.Color("#414B53")).
		Height(m.Height).
		Width(m.Width/2 - 1)

	return style.Render(
		lipgloss.JoinHorizontal(lipgloss.Top, filesStyle.Render(m.list.View()), m.Diffs[m.list.ActiveRow].View()),
	)
}
