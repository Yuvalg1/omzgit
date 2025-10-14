package log

import (
	"omzgit/consts"
	"omzgit/default/colors"
	"omzgit/default/colors/bg"
	"omzgit/default/colors/gray"
	"omzgit/default/style"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	hash := lipgloss.NewStyle().
		Background(colors.GetColor(m.Active, bg.C[2], bg.C[0])).
		Foreground(colors.Yellow).
		Render(m.Hash + " ")

	current := ""
	if len(m.branches) >= 1 {
		current = " " + lipgloss.NewStyle().
			Background(colors.Aqua).
			Padding(0, 1).
			Foreground(bg.C[0]).
			Render(m.branches[0])
	}

	desc := lipgloss.NewStyle().
		Background(colors.GetColor(m.Active, bg.C[2], bg.C[0])).
		Foreground(colors.Purple).
		Render(consts.TrimRight(m.Desc.View(), max(m.width-1-lipgloss.Width(hash)-lipgloss.Width(current), 0)))

	return style.Bg.Width(m.width).
		Background(colors.GetColor(m.Active, bg.C[2], bg.C[0])).
		Border(lipgloss.MarkdownBorder(), false, false, false, true).
		BorderBackground(bg.C[0]).
		BorderForeground(colors.GetColor(m.Active, gray.C[0], bg.C[2])).
		Width(m.width - 1).
		Render(hash + desc +
			lipgloss.NewStyle().
				Align(lipgloss.Right).
				Background(colors.GetColor(m.Active, bg.C[2], bg.C[0])).
				Width(m.width-1-lipgloss.Width(desc)-lipgloss.Width(hash)).
				Render(current))
}
