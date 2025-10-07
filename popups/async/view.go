package async

import (
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	return lipgloss.NewStyle().
		Background(bg.C[0]).
		Border(lipgloss.NormalBorder(), true).
		BorderBackground(bg.C[0]).
		BorderForeground(bg.C[4]).
		Height(m.height - 2).
		Render(
			lipgloss.NewStyle().Width(m.width-2-lipgloss.Width(m.spinner.View())).Background(bg.C[0]).Render(m.title+"... ") +
				lipgloss.NewStyle().Background(bg.C[0]).Render(m.spinner.View()))
}
