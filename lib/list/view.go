package list

import (
	"omzgit/default/colors/bg"
	"omzgit/default/colors/gray"
	"omzgit/default/style"

	"github.com/charmbracelet/lipgloss"
)

func (m Model[T]) View() string {
	var fileStrings string
	diff := max(m.ActiveRow-m.innerOffset, 0)

	if m.innerOffset >= m.ActiveRow {
		m.innerOffset = m.ActiveRow
	}

	if len(m.Children) == 0 {
		fileStrings = (*m.createChildFn(m.emptyMsg)).View()
	}

	for i := range min(max(len(m.Children)-diff, 0), m.height-1) {
		fileStrings += m.Children[i+diff].View() + "\n"
	}

	fileStrings = fileStrings[:max(len(fileStrings)-1, 0)]

	modeStyle := lipgloss.NewStyle().Background(gray.C[2]).Foreground(bg.C[0])
	endStyle := lipgloss.NewStyle().Background(bg.C[0]).Foreground(gray.C[2])

	mode := m.mode
	if mode != "" {
		mode = " " + mode + " " + endStyle.Render("")
	}

	modeRender := modeStyle.Render(mode)

	return style.Bg.Height(m.height).Render(modeRender + style.Bg.Render(m.getTextInput(m.width-lipgloss.Width(modeRender))) + "\n" + fileStrings)
}

func (m Model[T]) getTextInput(width int) string {
	if !m.TextInput.Focused() {
		basename := lipgloss.NewStyle().Background(gray.C[2]).Foreground(bg.C[0]).Padding(0, 1).Render(m.basename)
		connector := lipgloss.NewStyle().Background(bg.C[0]).Foreground(gray.C[2]).Render("")

		return lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Right).
			Width(width).
			Render(connector + basename)
	}

	return " " + m.TextInput.View()
}
