package messages

import tea "github.com/charmbracelet/bubbletea"

type ModeMsg struct {
	Mode string
}

type Moderer interface {
	ModeCmd(title string) tea.Cmd
}
