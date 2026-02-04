package refresh

import tea "github.com/charmbracelet/bubbletea"

type Msg struct{}

func Cmd() tea.Cmd {
	return func() tea.Msg {
		return Msg{}
	}
}
