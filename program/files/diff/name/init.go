package name

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Name string

	width int
}

func InitialModel(path string, width int) Model {
	parts := strings.Split(path, "/")

	return Model{
		Name:  parts[len(parts)-1],
		width: width,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
