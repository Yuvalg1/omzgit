package diff

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	block := m.viewport.Block

	m.viewport.SetContent(m.Content)
	m.viewport.SetBlock(block[0], block[1], m.width, m.height)

	return lipgloss.NewStyle().Width(m.width).Height(m.height).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#414B53")).
		Render(m.viewport.View())
}
