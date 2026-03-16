package files

import (
	"strings"

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

	if m.width > CUTOFF {
		return lipgloss.JoinHorizontal(lipgloss.Top, m.list.View(), middle, m.diff.View())
	}

	line := lipgloss.NewStyle().Background(bg.C[0]).Foreground(bg.C[4]).Render(strings.Repeat("─", m.width))

	return lipgloss.JoinVertical(lipgloss.Top, m.list.View(), line, m.diff.View())
}
