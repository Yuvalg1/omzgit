package diff

import (
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	return lipgloss.NewStyle().
		Background(bg.C[0]).
		Border(lipgloss.NormalBorder()).
		BorderBackground(bg.C[0]).
		BorderForeground(bg.C[4]).
		Height(m.viewport.Height).
		Width(m.viewport.Width).
		Render(m.viewport.View())
}
