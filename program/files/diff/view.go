package diff

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	style := lipgloss.NewStyle().Height(m.Height).Width(m.Width-1).
		Border(lipgloss.DoubleBorder(), false, false, false, true)

	content := ""

	size := min(m.Height-2, len(m.Content)-2)

	for i := range size {
		content += "\n" + m.Content[i].View()
	}

	return style.Render(m.Name.View() + content)
}
