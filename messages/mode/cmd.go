package mode

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Msg struct {
	Mode string
}

func Cmd(mode string) tea.Cmd {
	return func() tea.Msg {
		return Msg{Mode: mode}
	}
}
