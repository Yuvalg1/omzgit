package async

import (
	"omzgit/default/style"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	return style.Bg.
		Border(lipgloss.NormalBorder(), true).
		Height(m.height - 2).
		Width(m.width - 2).
		Render(
			lipgloss.NewStyle().Width(m.width-2-lipgloss.Width(m.spinner.View())).Render(m.title+"... ") +
				m.spinner.View())
}
