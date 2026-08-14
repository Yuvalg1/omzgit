package row

import (
	"omzgit/default/colors"
	"omzgit/default/colors/bg"
	"omzgit/default/colors/gray"
	"omzgit/default/style"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	return lipgloss.NewStyle().
		Background(colors.GetColor(m.Active, bg.C[2], bg.C[0])).
		Border(lipgloss.MarkdownBorder(), false, false, false, true).
		BorderBackground(bg.C[0]).
		BorderForeground(colors.GetColor(m.Active, gray.C[0], bg.C[2])).
		Foreground(getForeground(m.Conflict, m.Staged)).
		Width(m.width - 1).
		Render(m.status + " " +
			m.getStrikethroughStyle().Render(m.Roller.View()))
}

func (m Model) getStrikethroughStyle() lipgloss.Style {
	current := style.Bg.
		Background(colors.GetColor(m.Active, bg.C[2], bg.C[0])).
		Foreground(getForeground(m.Conflict, m.Staged))
	if m.status == "D" {
		return current.Strikethrough(true)
	}

	return current
}

func getForeground(conflict bool, staged bool) lipgloss.Color {
	if conflict {
		return colors.Orange
	}

	return colors.GetColor(staged, colors.Green, colors.Red)
}
