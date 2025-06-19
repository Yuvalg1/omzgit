package alert

import (
	"omzgit/consts"
	"omzgit/default/style"

	"github.com/charmbracelet/lipgloss"
)

var borderColor = lipgloss.Color("#CCCCCC")

func (m Model) View() string {
	titleStyle := style.Bg.Width(m.Width).Foreground(lipgloss.Color("#FA7970"))
	borderStyle := style.Bg.Border(lipgloss.NormalBorder(), false, true, true).Width(m.Width - 2).BorderForeground(lipgloss.Color("#FA7970")).Foreground(lipgloss.Color("#FA7970"))

	return titleStyle.Render(consts.PadTitle(m.verb, m.Width) + borderStyle.Render(m.error))
}
