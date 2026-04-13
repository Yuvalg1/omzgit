package content

import tea "github.com/charmbracelet/bubbletea"

type Msg struct {
	Index   int
	Command string
}

func Cmd(index int, command string) tea.Cmd {
	return func() tea.Msg {
		return Msg{
			Index:   index,
			Command: command,
		}
	}
}
