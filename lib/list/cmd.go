package list

import tea "github.com/charmbracelet/bubbletea"

type Msg struct {
	Count int
}

func Cmd() tea.Cmd {
	return func() tea.Msg {
		return Msg{}
	}
}
