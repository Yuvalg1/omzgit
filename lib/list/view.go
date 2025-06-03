package list

import "github.com/charmbracelet/lipgloss"

func (m Model[T]) View() string {
	var fileStrings string
	diff := max(m.ActiveRow-m.height+1, 0)

	children := m.getFilteredChildren()

	for i := range min(len(children), m.height-1) {
		fileStrings += children[i+diff].View() + "\n"
	}

	if len(fileStrings) == 0 {
		fileStrings = "No Files Found\n"
	}

	fileStrings = fileStrings[:max(len(fileStrings)-1, 0)]

	return m.getTextInput() + "\n" + fileStrings
}

func (m Model[T]) getTextInput() string {
	if !m.TextInput.Focused() {
		return ""
	}

	style := lipgloss.NewStyle().Width(m.width)
	return style.Render("Search " + m.TextInput.View())
}
