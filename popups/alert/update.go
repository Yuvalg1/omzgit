package alert

import (
	"omzgit/clipboard"
	"omzgit/env"
	"omzgit/program/popups"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = getWidth(msg.Width)
		m.Height = getHeight(msg.Height)
		return m, nil

	case popups.Msg:
		m.error = msg.Name
		m.visible = true
		m.verb = msg.Verb
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case env.Alert.CtrlC.Msg, env.Alert.Quit.Msg:
			return m, tea.Quit

		case env.Alert.Yank.Msg:
			clipboard.Copy(m.error)
			m.visible = false
			return m, nil

		default:
			m.visible = false
			return m, nil
		}
	}

	return m, nil
}
