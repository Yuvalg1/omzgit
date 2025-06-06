package messages

import tea "github.com/charmbracelet/bubbletea"

type Func struct {
	Func any
}

type PopupMsg struct {
	Fn   any
	Name string
	Type string
	Verb string
}

type Popuper[F any] interface {
	PopupCmd(string, string, string, F) tea.Cmd
}
