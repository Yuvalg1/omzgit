package input

import (
	"omzgit/default/colors"
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	CallbackFn    func(string)
	Name          string
	textinput     textinput.Model
	visible       bool
	withoutSpaces bool

	Width  int
	Height int
}

func InitialModel(fn func(string), width int, height int, withoutSpaces bool) Model {
	ti := textinput.New()
	ti.CharLimit = 50
	ti.Focus()
	ti.Width = getWidth(width) - 5
	ti.PlaceholderStyle = ti.PlaceholderStyle.Background(bg.C[0])
	ti.TextStyle = lipgloss.NewStyle().Foreground(colors.Purple).Background(bg.C[0])

	return Model{
		CallbackFn:    fn,
		Name:          "",
		textinput:     ti,
		visible:       false,
		withoutSpaces: withoutSpaces,

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
