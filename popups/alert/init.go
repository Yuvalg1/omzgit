package alert

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	error    string
	viewport viewport.Model
	visible  bool
	verb     string

	maxHeight int
}

func InitialModel(width int, height int) Model {
	viewport := viewport.New(getWidth(width), getHeight(height))

	return Model{
		error:    "",
		viewport: viewport,
		visible:  false,
		verb:     "",

		maxHeight: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getHeight(height int) int {
	return height - 8
}

func getWidth(width int) int {
	return max(width-12+width%2, 0)
}

func (m Model) GetVisible() bool {
	return m.visible
}
