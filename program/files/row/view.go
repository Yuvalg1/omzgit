package row

import (
	"program/consts"

	"github.com/charmbracelet/lipgloss"
)

var addedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7CE38B"))

var resetStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FA7970"))

var deletedStyle = lipgloss.NewStyle().Strikethrough(true)

func (m Model) View() string {
	path := m.Path
	path = consts.TrimRight(path, m.width-3)

	return getActiveStyle(m.Staged, m.Active).Width(m.width - 1).Render(m.status + " " +
		getDeletedStyle(m.Staged, m.Active, m.status).Render(path))
}

func getActiveStyle(added bool, active bool) lipgloss.Style {
	current := getStyle(added)
	if !active {
		return current.Background(lipgloss.Color("#21262D"))
	}

	return current.Background(lipgloss.Color("#3A444B"))
}

func getStyle(added bool) lipgloss.Style {
	if added {
		return addedStyle
	}

	return resetStyle
}

func getDeletedStyle(added bool, active bool, status string) lipgloss.Style {
	current := getActiveStyle(added, active)

	if status == "D" {
		return current.Strikethrough(true)
	}

	return current
}
