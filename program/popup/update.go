package popup

import (
	"program/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.TerminalMsg:
		m.Width = GetWidth(msg.Width)
		m.Height = GetHeight(msg.Height)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "n", "N":
			m.Visible = false
			return m, nil

		case "y", "Y":
			m.Fn()
			m.Visible = false
			return m, nil

		case "ctrl+c", "q":
			return m, tea.Quit

		default:
			return m, nil
		}
	}

	return m, nil
}
