package cokeline

import tea "github.com/charmbracelet/bubbletea"

type Msg struct {
	Left   string
	Center string
	Right  string

	Primary bool
}

func Cmd(center string, right string, primary bool) tea.Cmd {
	return func() tea.Msg {
		return Msg{
			Center:  center,
			Right:   right,
			Primary: primary,
		}
	}
}
