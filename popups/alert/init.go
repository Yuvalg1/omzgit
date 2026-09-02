package alert

import (
	"omzgit/program/popups"

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

func InitialModel(width int, height int, error chan string) Model {
	viewport := viewport.New(getWidth(width), getHeight(height))
	value := <-error
	visible := value != ""

	return Model{
		error:    value,
		viewport: viewport,
		visible:  visible,
		verb:     "Error!",

		maxHeight: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	if m.error != "" {
		return popups.Cmd("alert", m.verb, m.error, func(name string) {})
	}

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
