package list

import tea "github.com/charmbracelet/bubbletea"

type Msg[T tea.Model] struct {
	Children []T
	Active   int
	Total    int

	Msg tea.Msg
	Cmd tea.Cmd
}

func Cmd[T tea.Model](callback func() []T, active int, msg tea.Msg, cmd tea.Cmd) tea.Cmd {
	return func() tea.Msg {
		children := callback()
		return Msg[T]{
			Children: callback(),
			Active:   active,
			Total:    len(children),
			Msg:      msg,
			Cmd:      cmd,
		}
	}
}
