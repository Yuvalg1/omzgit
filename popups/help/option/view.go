package option

import (
	"omzgit/consts"
	"omzgit/default/colors"
	"omzgit/default/colors/bg"
	"omzgit/default/colors/gray"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	msg := lipgloss.NewStyle().
		Background(colors.GetColor(m.Active, bg.C[2], bg.C[0])).
		Foreground(colors.Green).
		Width(8).
		Render(m.Msg)

	description := consts.TrimRight(m.Roller.View(), m.width-lipgloss.Width(msg)-1)

	borderStyle := lipgloss.NewStyle().
		Background(colors.GetColor(m.Active, bg.C[2], bg.C[0])).
		Border(lipgloss.MarkdownBorder(), false, false, false, true).
		BorderBackground(bg.C[0]).
		BorderForeground(colors.GetColor(m.Active, gray.C[0], bg.C[2])).
		Width(m.width - 1)

	return borderStyle.Render(msg +
		lipgloss.NewStyle().
			Background(colors.GetColor(m.Active, bg.C[2], bg.C[0])).
			Foreground(gray.C[2]).
			Render(description))
}
