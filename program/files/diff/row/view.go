package row

import (
	"program/consts"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	style := lipgloss.NewStyle().Width(m.width).Height(1).MaxHeight(1)
	style = getStyle(m, style)

	strs := consts.TrimRight(m.text, m.width)

	return style.Render(string(strs))
}

func getStyle(m Model, style lipgloss.Style) lipgloss.Style {
	if m.descriptor == '@' {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#2563EB")).Inherit(style)
	}

	if m.descriptor == 'D' {
		return lipgloss.NewStyle().Bold(true).Inherit(style)
	}

	if m.descriptor == '+' {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#4ADE80")).Inherit(style)
	}

	if m.descriptor == '-' {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#F87171")).Inherit(style)
	}

	return style
}
