package diff

import (
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	block := m.viewport.Block

	m.viewport.SetContent(m.Content)
	m.viewport.SetBlock(block[0], block[1], m.width-3, m.height)

	return lipgloss.NewStyle().
		Background(bg.C[0]).
		Border(lipgloss.NormalBorder()).
		BorderBackground(bg.C[0]).
		BorderForeground(bg.C[4]).
		Height(m.height).
		Width(m.width - 3).
		Render(m.viewport.View())
}
