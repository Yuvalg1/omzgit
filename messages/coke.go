package messages

import tea "github.com/charmbracelet/bubbletea"

type CokeMsg struct {
	Title string
}

type Cokerer interface {
	CokeCmd(title string) tea.Cmd
}
