package input

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var borderColor = lipgloss.Color("#CCCCCC")

func (m Model) View() string {
	titleStyle := lipgloss.NewStyle().Width(m.Width)
	borderStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, true, true).Width(m.Width - 2)

	padding := m.Width - len(m.Name) - 4
	parity := len(m.Name) % 2
	return titleStyle.Render("┌"+strings.Repeat("─", padding/2)+" "+m.Name+" "+strings.Repeat("─", padding/2+parity)+"┐") + "\n" + borderStyle.Render(m.textinput.View())
}
