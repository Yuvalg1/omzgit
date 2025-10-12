package messages

import tea "github.com/charmbracelet/bubbletea"

type CokeMsg struct {
	Left   string
	Center string
	Right  string

	Primary bool
}

type Cokerer interface {
	CokeCmd() tea.Cmd
}
