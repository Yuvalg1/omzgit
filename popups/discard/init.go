package discard

import (
	"program/messages"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	CallbackFn func() bool
	Name       string
	visible    bool
	verb       string

	Width  int
	Height int
}

func InitialModel(fn func() bool, width int, height int) Model {
	return Model{
		CallbackFn: fn,
		Name:       "",
		visible:    false,
		verb:       "",

		Width:  getWidth(width),
		Height: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) PopupCmd(ptype string, placeholder string, title string, fn func() tea.Cmd) tea.Cmd {
	return func() tea.Msg {
		return messages.PopupMsg{
			Fn:   fn,
			Type: ptype,
			Name: title,
			Verb: placeholder,
		}
	}
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
