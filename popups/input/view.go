package input

import (
	"omzgit/consts"
	"omzgit/default/colors"
	"omzgit/default/style"

	"github.com/charmbracelet/lipgloss"
)

var borderColor = lipgloss.Color("#CCCCCC")

func (m Model) View() string {
	titleStyle := style.Bg.Width(m.Width).Foreground(colors.Yellow)
	borderStyle := style.Bg.
		Border(lipgloss.NormalBorder(), false, true, true).
		BorderForeground(colors.Yellow).
		Width(m.Width - 2).
		Foreground(colors.Blue)

	return titleStyle.Render(consts.PadTitle(m.Name, m.Width) + borderStyle.Render(m.textinput.View()))
}
