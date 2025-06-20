package files

import (
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	filesStyle := lipgloss.NewStyle().
		Background(bg.C[0]).
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderBackground(bg.C[0]).
		BorderForeground(bg.C[4]).
		Height(m.Height).
		Width(m.Width/2 - 1)

	return lipgloss.NewStyle().
		Render(lipgloss.JoinHorizontal(lipgloss.Top, filesStyle.Render(m.list.View()), m.Diffs[m.list.ActiveRow].View()))
}
