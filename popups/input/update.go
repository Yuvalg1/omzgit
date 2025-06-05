package input

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
		m.CallbackFn = msg.Fn.(func(string))
		m.Name = msg.Name
		m.visible = true
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "esc":
			m.visible = false
			return m, nil

		case "enter":
			m.CallbackFn(m.textinput.Value())
			m.visible = false
			return m, nil

		case "ctrl+c", "q":
			return m, tea.Quit

		default:
			res, cmd := m.textinput.Update(msg)
			m.textinput = res
			return m, cmd
		}
	}

	return m, nil
}
