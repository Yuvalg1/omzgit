package content

import tea "github.com/charmbracelet/bubbletea"

type Msg struct {
	Index int
	Ours  bool
}

func Cmd(index int, ours bool) tea.Cmd {
	return func() tea.Msg {
		return Msg{
			Index: index,
			Ours:  ours,
		}
	}
}
