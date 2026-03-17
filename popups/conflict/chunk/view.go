package chunk

import (
	"omzgit/default/colors"
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	return lipgloss.NewStyle().
		Background(colors.GetColor(m.Active, bg.C[2], bg.C[0])).
		Foreground(colors.GetColor(m.ours, colors.Green, colors.Red)).
		Width(m.width).
		Render(m.content)
}
