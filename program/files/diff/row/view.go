package row

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	style := lipgloss.NewStyle().Width(m.Width).Height(1).MaxHeight(1)
	style = getStyle(m, style)

	strs := m.text
	strs = strs[:min(len(strs), m.Width)]

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
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#16A34A")).Inherit(style)
	}

	if m.descriptor == '-' {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#DC2626")).Inherit(style)
	}

	return style
}
