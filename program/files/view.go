package files

import (
	"github.com/charmbracelet/lipgloss"
)

var leftText = "Files"

func (m Model) View() string {
	style := lipgloss.NewStyle().Width(m.Width)
	var fileStrings string

	diff := max(m.ActiveRow-m.Height+1, 0)

	for i := range min(len(m.files), m.Height) {
		fileStrings += m.files[i+diff].View() + "\n"
	}

	fileStrings = fileStrings[:len(fileStrings)-1]

	filesStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(lipgloss.Color("#414B53")).
		Background(lipgloss.Color("#21262D")).
		Height(m.Height).
		Width(m.Width/2 - 1)

	return style.Render(
		lipgloss.JoinHorizontal(lipgloss.Top, filesStyle.Render(fileStrings), m.Diffs[m.ActiveRow].View()),
	)
}
