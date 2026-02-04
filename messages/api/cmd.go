package api

import tea "github.com/charmbracelet/bubbletea"

type Msg struct {
	Response tea.Cmd
}

func Cmd(response tea.Cmd) tea.Cmd {
	return func() tea.Msg {
		return Msg{Response: response}
	}
}
