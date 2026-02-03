package branch

import (
	"omzgit/clipboard"
	"omzgit/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case *messages.RefreshMsg:
		m.Active = true
		return m, nil

	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)

		res, cmd := m.Roller.Update(msg)
		m.Roller = res

		return m, cmd

	case messages.RollerMsg:
		res, cmd := m.Roller.Update(msg)
		m.Roller = res

		return m, cmd

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			m.Active = true
			m.lastUpdated = m.getLastUpdatedDate()
			m.diff = m.getBranchDiff()
			return m, nil

		case "j", "k", "down", "up", "g", "G", "/", "esc":
			m.Active = !m.Active
			m.lastUpdated = m.getLastUpdatedDate()
			m.diff = m.getBranchDiff()

			m.Roller.Width = m.width - len(m.diff) - len(m.lastUpdated) - 3
			res, cmd := m.Roller.Update(msg)
			m.Roller = res

			return m, cmd

		case "y":
			clipboard.Copy(m.Roller.Name)
			return m, nil

		default:
			return m, nil
		}
	}

	return m, nil
}
