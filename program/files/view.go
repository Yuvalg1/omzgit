package files

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	var fileStrings string

	diff := max(m.ActiveRow-m.Height+1, 0)

	for i := range min(len(m.files), m.Height) {
		fileStrings += m.files[i+diff].View() + "\n"
	}

	if len(fileStrings) > 0 {
		fileStrings = fileStrings[:len(fileStrings)-1]
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, fileStrings, m.Diffs[m.ActiveRow].View())
}
