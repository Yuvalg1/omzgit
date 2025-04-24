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

	return style.Render(
		lipgloss.JoinHorizontal(lipgloss.Top, fileStrings, m.Diffs[m.ActiveRow].View()),
	)
}
