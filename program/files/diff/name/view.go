package name

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), false, true, true, true).Padding(0, 1).MaxWidth(m.width - 2).Height(1)

	name := m.Name

	return lipgloss.PlaceHorizontal(m.width, lipgloss.Center, style.Render(name))
}
