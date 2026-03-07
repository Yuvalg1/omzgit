package option

import (
	"omzgit/messages/refresh"
	"omzgit/roller"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case refresh.Msg:
		m.Active = true
		return m, nil

	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)

		res, cmd := m.Roller.Update(msg)
		m.Roller = res

		return m, cmd

	case roller.Msg:
		res, cmd := m.Roller.Update(msg)
		m.Roller = res

		return m, cmd

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "j", "k", "down", "up", "g", "G", "/", "esc":
			m.Active = !m.Active

			m.Roller.Width = m.width - 8
			res, cmd := m.Roller.Update(msg)
			m.Roller = res

			return m, cmd

		default:
			return m, nil
		}
	}

	return m, nil
}
