package branch

import (
	"omzgit/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.TerminalMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			m.Active = true
			return m, nil

		case "g", "G", "/", "esc":
			m.Active = !m.Active
			return m, nil

		case "j", "k", "down", "up":
			m.Active = !m.Active
			return m, nil

		default:
			return m, nil
		}
	}

	return m, nil
}
