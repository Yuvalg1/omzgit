package input

import (
	"omzgit/default/colors"
	"omzgit/messages"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	CallbackFn func(string)
	Name       string
	textinput  textinput.Model
	visible    bool

	Width  int
	Height int
}

func InitialModel(fn func(string), width int, height int) Model {
	ti := textinput.New()
	ti.CharLimit = 50
	ti.Focus()
	ti.Width = getWidth(width)
	ti.PlaceholderStyle = ti.PlaceholderStyle.Foreground(colors.Blue)

	return Model{
		CallbackFn: fn,
		Name:       "",
		textinput:  ti,
		visible:    false,

		Width:  getWidth(width),
		Height: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) RefreshCmd() tea.Cmd {
	return func() tea.Msg {
		return messages.RefreshMsg{}
	}
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
