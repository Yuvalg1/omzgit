package program

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	style := lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).Height(m.Height - 2).Width(m.Width - 2)

	if m.Popup.Visible {
		return style.Render(m.Popup.View())
	}

	return style.Render(m.Tabs[m.ActiveTab].View())
}
