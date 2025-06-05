package alert

import (
	"program/consts"

	"github.com/charmbracelet/lipgloss"
)

var borderColor = lipgloss.Color("#CCCCCC")

func (m Model) View() string {
	titleStyle := lipgloss.NewStyle().Width(m.Width).Foreground(lipgloss.Color("#FA7970"))
	borderStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, true, true).Width(m.Width - 2).BorderForeground(lipgloss.Color("#FA7970")).Foreground(lipgloss.Color("#FA7970"))

	return titleStyle.Render(consts.PadTitle("Alert!", m.Width) + borderStyle.Render(m.error))
}
