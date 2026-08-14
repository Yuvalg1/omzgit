package suspend

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Msg struct {
	Command []string
}

func Cmd(command ...string) tea.Cmd {
	return func() tea.Msg {
		return Msg{Command: command}
	}
}
