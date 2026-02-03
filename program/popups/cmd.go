package popups

import tea "github.com/charmbracelet/bubbletea"

type Msg struct {
	Fn   any
	Name string
	Type string
	Verb string
}

func Cmd(pType string, name string, verb string, fn any) tea.Cmd {
	return func() tea.Msg {
		return Msg{
			Type: pType,
			Name: name,
			Verb: verb,
			Fn:   fn,
		}
	}
}
