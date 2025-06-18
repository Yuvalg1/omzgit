package messages

import tea "github.com/charmbracelet/bubbletea"

type ApiMsg struct {
	Response tea.Cmd
}

type Apier interface {
	ApiCmd() tea.Cmd
}
