package input

import (
	"omzgit/consts"
	"omzgit/default/colors/bg"
	"omzgit/default/style"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	titleStyle := style.Bg.Width(m.Width).Background(bg.C[0])
	borderStyle := lipgloss.NewStyle().
		Background(bg.C[0]).
		Border(lipgloss.NormalBorder(), false, true, true).
		BorderBackground(bg.C[0]).
		Width(m.Width - 2)

	return titleStyle.Render(consts.PadTitle(m.Name, m.Width) + borderStyle.Render(m.textinput.View()))
}
