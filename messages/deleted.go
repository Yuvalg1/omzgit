package messages

import tea "github.com/charmbracelet/bubbletea"

type RefreshMsg struct{}

type Refresher interface {
	RefreshCmd() tea.Cmd
}
