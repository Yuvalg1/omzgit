package messages

import tea "github.com/charmbracelet/bubbletea"

type PopupMsg struct {
	Fn   func()
	Name string
}

type Popuper interface {
	PopupCmd(string, func()) tea.Cmd
}
