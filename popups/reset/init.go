package reset

import (
	"omzgit/messages"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	name    string
	hash    string
	visible bool
	options map[byte]string

	width  int
	height int
}

func InitialModel(width int, height int) Model {
	options := map[byte]string{}
	options['s'] = "--soft"
	options['h'] = "--hard"
	options['m'] = "--mixed"

	return Model{
		options: options,
		visible: false,

		width:  getWidth(width),
		height: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) PopupCmd(pType string, placeholder string, title string, fn any) tea.Cmd {
	return func() tea.Msg {
		return messages.PopupMsg{
			Fn:   fn,
			Name: title,
			Type: pType,
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
	return 4
}

func getWidth(width int) int {
	return min(34, width-2-width%2)
}

func (m Model) GetVisible() bool {
	return m.visible
}
