package files

import (
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	middle := lipgloss.NewStyle().
		Background(bg.C[0]).
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderBackground(bg.C[0]).
		BorderForeground(bg.C[4]).
		Height(m.height).Width(0).Render("")

	return lipgloss.JoinHorizontal(lipgloss.Top, m.list.View(), middle, m.diffs[m.list.ActiveRow].View())
}
