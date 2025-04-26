package cokeline

import (
	"program/program/lib/button"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	style := lipgloss.NewStyle().Width(m.width).Height(1).Padding(0, 0, 1)

	tabTitleStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#77BDFB")).
		Padding(0, 1).
		Foreground(lipgloss.Color("#21262D")).
		Bold(true)

	tabTitle := tabTitleStyle.Render(m.Cokes[m.ActiveCoke])

	endStyle := lipgloss.NewStyle().Background(lipgloss.Color("#21262D")).Width(m.width - lipgloss.Width(tabTitle) - lipgloss.Width(m.title)).Align(lipgloss.Right)
	endButtons := endStyle.Render(button.InitialModel("[]").View())

	return style.Render(tabTitle + m.title + endButtons)
}
