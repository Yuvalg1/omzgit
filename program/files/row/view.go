package row

import (
	"program/consts"

	"github.com/charmbracelet/lipgloss"
)

var addedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7FBBB3"))

var resetStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7A8478"))

var deletedStyle = lipgloss.NewStyle().Strikethrough(true)

func (m Model) View() string {
	path := m.Path
	path = consts.TrimRight(path, m.width-2)

	return getActiveStyle(2, m.Staged, m.Active).Render(m.status+" ") + getDeletedStyle(m.width-2, m.Staged, m.Active, m.status).Render(path)
}

func getActiveStyle(width int, added bool, active bool) lipgloss.Style {
	current := getStyle(width, added)
	if !active {
		return current.Background(lipgloss.Color("#272E33"))
	}

	return current.Background(lipgloss.Color("#384B55"))
}

func getStyle(width int, added bool) lipgloss.Style {
	style := lipgloss.NewStyle().Width(width)

	if added {
		return addedStyle.Width(width).Inherit(style)
	}

	return resetStyle.Width(width).Inherit(style)
}

func getDeletedStyle(width int, added bool, active bool, status string) lipgloss.Style {
	current := getActiveStyle(width, added, active)

	if status == "D" {
		return current.Strikethrough(true)
	}

	return current
}
