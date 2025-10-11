package list

import (
	"omzgit/default/style"

	"github.com/charmbracelet/lipgloss"
)

func (m Model[T]) View() string {
	var fileStrings string
	diff := max(m.ActiveRow-m.innerOffset, 0)

	if m.innerOffset >= m.ActiveRow {
		m.innerOffset = m.ActiveRow
	}

	children := m.getFilteredChildren()

	if len(children) == 0 {
		fileStrings = (*m.createChildFn(m.emptyMsg)).View()
	}

	for i := range min(max(len(children)-diff, 0), m.height-1) {
		fileStrings += children[i+diff].View() + "\n"
	}

	fileStrings = fileStrings[:max(len(fileStrings)-1, 0)]

	return style.Bg.Height(m.height).Render(m.getTextInput() + "\n" + fileStrings)
}

func (m Model[T]) getTextInput() string {
	if !m.TextInput.Focused() {
		return ""
	}

	return lipgloss.NewStyle().Render("Search " + m.TextInput.View())
}
