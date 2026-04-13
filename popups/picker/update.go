package picker

import (
	"omzgit/program/popups"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)
		return m, nil

	case popups.Msg:
		m.visible = true
		m.title = msg.Name
		m.name = msg.Verb
		m.options = msg.Fn.(func() map[string]Pick)()
		return m, nil

	case tea.KeyMsg:
		if pick, ok := m.options[msg.String()]; ok {
			m.visible = false
			return m, pick.Callback()
		}
		switch keypress := msg.String(); keypress {
		case "esc":
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
