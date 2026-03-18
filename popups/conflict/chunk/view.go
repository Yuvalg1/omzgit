package chunk

import (
	"omzgit/default/colors"
	"omzgit/default/colors/bg"
	"omzgit/default/colors/gray"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	ourColor := colors.GetColor(m.ours, colors.Red, colors.Green)

	return lipgloss.NewStyle().
		Background(colors.GetColor(m.Active, bg.C[2], bg.C[0])).
		Foreground(colors.GetColor(m.Conflict, ourColor, gray.C[2])).
		Width(m.Width).
		Render(m.Content[:len(m.Content)-1])
}
