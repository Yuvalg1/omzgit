package input

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	CallbackFn func(string)
	error      string
	Name       string
	textinput  textinput.Model
	visible    bool

	Width  int
	Height int
}

func InitialModel(fn func(string), name string, width int, height int) Model {
	ti := textinput.New()
	ti.CharLimit = 50
	ti.Focus()

	return Model{
		CallbackFn: fn,
		error:      "",
		Name:       name,
		textinput:  ti,
		visible:    false,

		Width:  getWidth(width),
		Height: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getHeight(height int) int {
	return 5
}

func getWidth(width int) int {
	return min(34, width-2)
}

func (m Model) GetVisible() bool {
	return m.visible
}
