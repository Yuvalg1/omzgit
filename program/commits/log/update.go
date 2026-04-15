package log

import (
	"omzgit/clipboard"
	"omzgit/env"
	"omzgit/git"
	"omzgit/messages/refresh"
	"omzgit/program/popups"
	"omzgit/roller"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)

		res, cmd := m.Desc.Update(msg)
		m.Desc = res

		return m, cmd

	case refresh.Msg:
		m.Active = true

		return m, nil

	case roller.Msg:
		res, cmd := m.Desc.Update(msg)
		m.Desc = res
		return m, cmd

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case env.Commits.Checkout.Msg:
			output, err := git.Exec("checkout", m.Hash)
			if err != nil {
				return m, popups.Cmd("alert", "Checkout Error!", output, func(name string) {})
			}

			m.Current = true
			return m, refresh.Cmd()

		case env.Commits.CheckoutForce.Msg:
			return m, popups.Cmd("discard", "force checkout", m.Hash, func() tea.Cmd {
				output, err := git.Exec("checkout", "-f", m.Hash)
				if err != nil {
					return popups.Cmd("alert", "Force Checkout Error!", output, func(name string) {})
				}

				m.Current = true
				return refresh.Cmd()
			})

		case env.Commits.Up.Msg, env.Commits.Up.AltMsg, env.Commits.Down.Msg, env.Commits.Down.AltMsg, env.Goto.Top.Msg, env.Commits.Bottom.Msg, env.Commits.Search.Msg, env.Commits.Refresh.Msg:
			m.Active = !m.Active
			res, cmd := m.Desc.Update(msg)

			m.Desc = res

			return m, cmd

		case env.Commits.Yank.Msg:
			clipboard.Copy(m.Hash)
			return m, nil

		case env.Commits.CherryPick.Msg:
			output, err := git.Exec("cherry-pick", m.Hash)
			if err != nil {
				popups.Cmd("alert", "Cherry Pick Error!", output, func() {})
			}
			return m, refresh.Cmd()

		default:
			return m, nil
		}
	}

	return m, nil
}
