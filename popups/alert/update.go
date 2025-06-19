package alert

import (
	"omzgit/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.TerminalMsg:
		m.Width = getWidth(msg.Width)
		m.Height = getHeight(msg.Height)
		return m, nil

	case messages.PopupMsg:
		m.error = msg.Name
		m.visible = true
		m.verb = msg.Verb
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			return m, tea.Quit

		default:
			m.visible = false
			return m, nil
		}
	}

	return m, nil
}
