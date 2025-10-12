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
		Foreground(colors.GetColor(m.Staged, colors.Green, colors.Red)).
		Width(m.width - 1).
		Render(m.status + " " +
			getStrikethroughStyle(m.Active, m.Staged, m.status).Render(m.Roller.View()))
}

func getStrikethroughStyle(active bool, staged bool, status string) lipgloss.Style {
	current := style.Bg.
		Background(colors.GetColor(active, bg.C[2], bg.C[0])).
		Foreground(colors.GetColor(staged, colors.Green, colors.Red))
	if status == "D" {
		return current.Strikethrough(true)
	}

	return current
}
