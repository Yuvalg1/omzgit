package messages

import tea "github.com/charmbracelet/bubbletea"

type TickMsg struct{}

type Ticker interface {
	TickCmd() tea.Cmd
}
