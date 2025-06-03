package row

import (
	"program/consts"

	"github.com/charmbracelet/lipgloss"
)

const (
	addedColor = "#7CE38B"
	resetColor = "#FA7970"
)

var deletedStyle = lipgloss.NewStyle().Strikethrough(true)

func (m Model) View() string {
	path := m.Path
	path = consts.TrimRight(path, m.width-3)

	return getStyle(m.Staged, m.Active).Width(m.width - 1).Render(m.status + " " +
		getStrikethroughStyle(m.Staged, m.Active, m.status).Render(path))
}

func getStyle(added bool, active bool) lipgloss.Style {
	color := getColor(added)
	if !active {
		return lipgloss.NewStyle().Foreground(lipgloss.Color(color))
	}

	return lipgloss.NewStyle().Background(lipgloss.Color(color)).Foreground(lipgloss.Color("#21262D"))
}

func getColor(added bool) string {
	if added {
		return addedColor
	}

	return resetColor
}

func getStrikethroughStyle(added bool, active bool, status string) lipgloss.Style {
	current := getStyle(added, active)

	if status == "D" {
		return current.Strikethrough(true)
	}

	return current
}
