package branch

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Name        string
	lastUpdated string
	current     bool
	diff        string

	width  int
	height int
}

func InitialModel(width int, height int, name string, lastUpdated string, diff string) Model {
	return Model{
		Name:        name[2:],
		lastUpdated: lastUpdated,
		current:     strings.Contains(name[:2], "*"),
		diff:        diff,

		width:  getWidth(width),
		height: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width
}

func getHeight(height int) int {
	return 1
}
