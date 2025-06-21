package input

import (
	"omzgit/consts"
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	titleStyle := lipgloss.NewStyle().
		Background(bg.C[0]).
		Foreground(bg.C[4]).
		Width(m.Width)

	borderStyle := lipgloss.NewStyle().
		Background(bg.C[0]).
		Border(lipgloss.NormalBorder(), false, true, true).
		BorderBackground(bg.C[0]).
		BorderForeground(bg.C[4]).
		Width(m.Width - 2)

	return titleStyle.Render(consts.PadTitle(m.Name, m.Width) + borderStyle.Render(m.textinput.View()))
}
