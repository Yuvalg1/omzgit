package branch

import (
	"program/consts"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	rest := " " + m.lastUpdated + " " + m.diff + " "

	titleStyle := getTitleStyle(m.current, m.width-lipgloss.Width(rest))
	title := titleStyle.Render(consts.TrimRight(m.Name, m.width-lipgloss.Width(rest)))

	return title + rest
}

func getTitleStyle(current bool, width int) lipgloss.Style {
	style := lipgloss.NewStyle().Width(width)
	if current {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Inherit(style)
	}
	return style
}
