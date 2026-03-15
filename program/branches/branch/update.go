package branch

import (
	"omzgit/clipboard"
	"omzgit/env"
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
		case env.Branches.Checkout.Msg, env.Branches.Checkout.AltMsg:
			m.Active = true
			m.lastUpdated = m.getLastUpdatedDate()
			m.diff = m.getBranchDiff()
			return m, nil

		case env.Branches.Down.Msg, env.Branches.Down.AltMsg, env.Branches.Up.Msg, env.Branches.Up.AltMsg, env.Goto.Top.Msg, env.Branches.Bottom.Msg, env.Branches.Search.Msg:
			m.Active = !m.Active
			m.lastUpdated = m.getLastUpdatedDate()
			m.diff = m.getBranchDiff()

			m.Roller.Width = m.width - len(m.diff) - len(m.lastUpdated) - 3
			res, cmd := m.Roller.Update(msg)
			m.Roller = res

			return m, cmd

		case env.Branches.Yank.Msg:
			clipboard.Copy(m.Roller.Name)
			return m, nil

		default:
			return m, nil
		}
	}

	return m, nil
}
