package branch

import (
	"program/consts"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	rest := " " + m.diff + " " + m.lastUpdated + " "

	titleStyle := m.getTitleStyle()
	title := consts.TrimRight(m.Name, m.width-lipgloss.Width(rest))

	return titleStyle.Render(title + lipgloss.NewStyle().Width(m.width-lipgloss.Width(title)).AlignHorizontal(lipgloss.Right).Render(rest))
}

func getColor(current bool) lipgloss.Color {
	if current {
		return lipgloss.Color("#FFFF66")
	}

	return lipgloss.Color("#02FFE4")
}

func (m Model) getTitleStyle() lipgloss.Style {
	color := getColor(m.Current)

	if m.Active {
		return lipgloss.NewStyle().Background(color).Foreground(lipgloss.Color("#21262D"))
	}
	return lipgloss.NewStyle().Foreground(color)
}
