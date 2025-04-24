package diff

import (
	"program/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.TerminalMsg:
		m.SetWidth(msg.Width)
		m.SetHeight(msg.Height)

		msg.Width = getWidth(msg.Width)
		msg.Height = getHeight(msg.Height)

		var cmds []tea.Cmd

		m.Content = m.getDiffStaged()
		return m, tea.Batch(cmds...)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case "a", "A":
			if m.staged {
				return m, nil
			}

			m.staged = true
			m.Content = m.getDiffStaged()
			return m, nil

		case "r", "R":
			if !m.staged {
				return m, nil
			}
			m.staged = false
			m.Content = m.getDiffStaged()
			return m, nil

		case "j", "k", "up", "down":
			m.Content = m.getDiffStaged()
			return m, nil

		default:
			return m, nil
		}
	}

	return m, nil
}
