package cokeline

import (
	"program/program/lib/button"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	style := lipgloss.NewStyle().Width(m.width).Height(1).Padding(0, 0, 1)

	tabTitleStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#A7C080")).
		Padding(0, 1).
		Foreground(lipgloss.Color("#1E2326")).
		Bold(true)

	tabTitle := tabTitleStyle.Render(m.Cokes[m.ActiveCoke])

	endStyle := lipgloss.NewStyle().Width(m.width - 4 - lipgloss.Width(tabTitle) - len(m.title)).Align(lipgloss.Right)
	endButtons := endStyle.Render(button.InitialModel("[]/h").View())

	return style.Render(tabTitle + " " + m.title + endButtons)
}
