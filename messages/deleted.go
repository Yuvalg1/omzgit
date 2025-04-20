package messages

import tea "github.com/charmbracelet/bubbletea"

type DeletedMsg struct{}

type Deleter interface {
	DeleteCmd() tea.Cmd
}
