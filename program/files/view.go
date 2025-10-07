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
		Height(m.Height).Width(1).Render("")

	return lipgloss.JoinHorizontal(lipgloss.Top, m.list.View(), middle, m.Diffs[m.list.ActiveRow].View())
}
