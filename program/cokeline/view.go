package cokeline

import (
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	partStyle := lipgloss.NewStyle().
		Background(bg.C[3]).
		Foreground(bg.C[0])

	tabTitle := partStyle.Render(m.Left)

	end := lipgloss.NewStyle().
		Align(lipgloss.Right).
		Inherit(partStyle).
		Render(m.Right)

	title := partStyle.
		Align(lipgloss.Center).
		Width(m.width - lipgloss.Width(end) - lipgloss.Width(tabTitle)).
		Render(m.Center)

	return lipgloss.NewStyle().
		Width(m.width).
		Height(1).
		Padding(0, 0, 1).
		Render(tabTitle + title + end)
}
