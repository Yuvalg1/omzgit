package diff

import (
	"program/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	block := m.viewport.Block

	m.viewport.SetContent(m.Content)
	m.viewport.SetBlock(block[0], block[1], m.width-3, m.height)

	return lipgloss.NewStyle().Width(m.width - 3).Height(m.height).
		Background(bg.C[0]).
		Border(lipgloss.NormalBorder()).
		BorderForeground(bg.C[4]).
		BorderBackground(bg.C[0]).
		Render(m.viewport.View())
}
