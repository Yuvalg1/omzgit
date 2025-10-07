package branch

import (
	"omzgit/consts"
	"omzgit/default/colors"
	"omzgit/default/colors/bg"
	"omzgit/default/colors/gray"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	rest := " " + m.diff + " " + m.lastUpdated + " "
	title := consts.TrimRight(m.Name, m.width-len(rest))

	borderStyle := lipgloss.NewStyle().
		Background(colors.GetColor(m.Active, bg.C[2], bg.C[0])).
		Border(lipgloss.MarkdownBorder(), false, false, false, true).
		BorderBackground(bg.C[0]).
		BorderForeground(colors.GetColor(m.Active, gray.C[0], bg.C[2])).
		Width(m.width - 1)

	return borderStyle.Inherit(m.getTitleStyle()).Render(title +
		lipgloss.NewStyle().Width(m.width-1-lipgloss.Width(title)).AlignHorizontal(lipgloss.Right).Render(rest))
}

func getColor(current bool) lipgloss.Color {
	if current {
		return colors.Yellow
	}

	return colors.Blue
}

func (m Model) getTitleStyle() lipgloss.Style {
	color := getColor(m.Current)

	if m.Active {
		return lipgloss.NewStyle().Foreground(color).Background(bg.C[2])
	}
	return lipgloss.NewStyle().Foreground(color).Background(bg.C[0])
}
