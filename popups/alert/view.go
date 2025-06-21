package alert

import (
	"omzgit/consts"
	"omzgit/default/colors"
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	titleStyle := lipgloss.NewStyle().
		Background(bg.C[0]).
		Foreground(colors.Red).
		Width(m.Width)

	borderStyle := lipgloss.NewStyle().
		Background(bg.C[0]).
		Border(lipgloss.NormalBorder(), false, true, true).
		BorderBackground(bg.C[0]).
		BorderForeground(colors.Red).
		Foreground(colors.Red).
		Width(m.Width - 2)

	return titleStyle.Render(consts.PadTitle(m.verb, m.Width) + borderStyle.Render(m.error))
}
