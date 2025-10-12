package messages

import tea "github.com/charmbracelet/bubbletea"

type TickMsg struct {
	RollOffset int
}

type Ticker interface {
	TickCmd() tea.Cmd
}
