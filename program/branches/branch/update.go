package branch

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)
		return m, nil

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
			return m, nil

		default:
			return m, nil
		}
	}

	return m, nil
}
