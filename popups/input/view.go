package input

import (
	"program/consts"

	"github.com/charmbracelet/lipgloss"
)

var borderColor = lipgloss.Color("#CCCCCC")

func (m Model) View() string {
	titleStyle := lipgloss.NewStyle().Width(m.Width).Foreground(lipgloss.Color("#FFFF66"))
	borderStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, true, true).Width(m.Width - 2).BorderForeground(lipgloss.Color("#FFFF66")).Foreground(lipgloss.Color("#02FFE4"))

	return titleStyle.Render(consts.PadTitle(m.Name, m.Width) + borderStyle.Render(m.textinput.View()))
}
