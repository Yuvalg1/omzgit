package list

import (
	"omzgit/default/style"

	"github.com/charmbracelet/lipgloss"
)

func (m Model[T]) View() string {
	var fileStrings string
	diff := max(m.ActiveRow-m.height+2, 0)

	children := m.getFilteredChildren()

	if len(children) == 0 && !m.TextInput.Focused() {
		fileStrings = m.Children[len(m.Children)-1].View()
	}

	for i := range min(len(children), m.height-1) {
		fileStrings += children[i+diff].View() + "\n"
	}

	fileStrings = fileStrings[:max(len(fileStrings)-1, 0)]

	return style.Bg.Width(m.width).Height(m.height).Render(m.getTextInput() + "\n" + fileStrings)
}

func (m Model[T]) getTextInput() string {
	if !m.TextInput.Focused() {
		return ""
	}

	return lipgloss.NewStyle().Width(m.width).Render("Search " + m.TextInput.View())
}
