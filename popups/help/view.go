package help

import (
	"omzgit/consts"
	"omzgit/default/colors"
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	titleStyle := lipgloss.NewStyle().
		Background(bg.C[0]).
		Foreground(colors.Green)

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, true).
		BorderBackground(bg.C[0]).
		BorderForeground(colors.Green).
		Width(m.width)

	return titleStyle.Render(consts.PadTitle("help", m.width+2) +
		borderStyle.Render(m.list.View()))
}
