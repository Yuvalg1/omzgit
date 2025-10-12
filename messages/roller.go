package messages

import tea "github.com/charmbracelet/bubbletea"

type RollerMsg struct {
	Id string
}

type Rollerer interface {
	RollerCmd() tea.Cmd
}
