package discard

import (
	"program/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.TerminalMsg:
		m.Width = getWidth(msg.Width)
		m.Height = getHeight(msg.Height)
		return m, nil

	case messages.PopupMsg:
		m.CallbackFn = msg.Fn.(func())
		m.Name = msg.Name
		m.visible = true
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "n", "N":
			m.visible = false
			return m, nil

		case "y", "Y":
			m.CallbackFn()
			m.visible = false
			return m, nil

		case "ctrl+c", "q":
			return m, tea.Quit

		default:
			return m, nil
		}
	}

	return m, nil
}
