package files

import (
	"fmt"
	"program/border"

	"github.com/charmbracelet/lipgloss"
)

var leftText = "Files"

func (m Model) View() string {
	style := lipgloss.NewStyle().Border(lipgloss.DoubleBorder(), false, true, true).Width(m.Width)
	var fileStrings string

	diff := max(m.ActiveRow-m.Height+1, 0)

	for i := range min(len(m.files), m.Height) {
		fileStrings += m.files[i+diff].View() + "\n"
	}

	fileStrings = fileStrings[:len(fileStrings)-1]

	title := fmt.Sprintf("%s: %d/%d", m.title, m.ActiveRow+1, len(m.files))

	return border.GetTopBorder(title, m.Width) +
		style.Render(
			lipgloss.JoinHorizontal(lipgloss.Top, fileStrings, m.Diffs[m.ActiveRow].View()),
		)
}
