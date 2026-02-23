package discard

import (
	"omzgit/messages/refresh"
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
		m.CallbackFn = msg.Fn.(func() tea.Cmd)
		m.Name = msg.Name
		m.visible = true
		m.verb = msg.Verb
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "esc", "n", "N":
			m.visible = false
			return m, nil

		case "y", "Y", "enter":
			callbackCmd := m.CallbackFn()
			if callbackCmd != nil {
				return m, callbackCmd
			}
			m.visible = false
			return m, refresh.Cmd()

		case "ctrl+c", "q":
			return m, tea.Quit

		default:
			return m, nil
		}
	}

	return m, nil
}
