package branch

import (
	"omzgit/clipboard"
	"omzgit/env"
	"omzgit/git"
	"omzgit/messages/refresh"
	"omzgit/popups/picker"
	"omzgit/program/popups"
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
			output, err := git.Exec("checkout", m.Roller.Name)
			if err != nil {
				return m, popups.Cmd("alert", "Checkout Error!", output, func(name string) {})
			}

			m.Active = true
			m.lastUpdated = m.getLastUpdatedDate()
			m.diff = m.getBranchDiff()
			m.Current = true
			return m, nil

		case env.Branches.CheckoutForce.Msg:
			return m, popups.Cmd("discard", "force checkout", m.Roller.Name, func() tea.Cmd {
				output, err := git.Exec("checkout", "-f", m.Roller.Name)
				if err != nil {
					return popups.Cmd("alert", "Force Checkout Error!", output, func(name string) {})
				}

				m.Current = true
				return nil
			})

		case env.Branches.Delete.Msg:
			return m, popups.Cmd("discard", "delete", m.Roller.Name, func() tea.Cmd {
				output, err := git.Exec("branch", "-d", m.Roller.Name)
				if err == nil {
					return nil
				}

				return popups.Cmd("alert", "Delete Error!", output, func(name string) {})
			})

		case env.Branches.DeleteForce.Msg:
			return m, popups.Cmd("discard", "force delete", m.Roller.Name, func() tea.Cmd {
				output, err := git.Exec("branch", "-D", m.Roller.Name)
				if err == nil {
					return nil
				}

				return popups.Cmd("alert", "Force Delete Error!", output, func(name string) {})
			})

		case env.Branches.Rebase.Msg:
			return m, popups.Cmd("async", "", "rebasing", func() tea.Cmd {
				output, err := git.Exec("rebase", m.Roller.Name)
				if err == nil {
					return nil
				}

				return popups.Cmd("alert", "Rebase Error!", output, func() tea.Cmd { return nil })
			})

		case env.Branches.RebaseOptions.Msg:
			return m, popups.Cmd("pick", m.Roller.Name, "choose a rebase option for "+m.Roller.Name, func() map[string]picker.Pick {
				return map[string]picker.Pick{
					"a": picker.GetPick("rebase", "--abort"),
					"c": picker.GetPick("rebase", "--continue"),
					"s": picker.GetPick("rebase", "--skip"),
					"q": picker.GetPick("rebase", "--quit"),
				}
			})

		case env.Branches.Merge.Msg:
			return m, popups.Cmd("async", "", "merging", func() tea.Cmd {
				output, err := git.Exec("merge", m.Roller.Name)
				if err == nil {
					return nil
				}

				return popups.Cmd("alert", "Merge Error!", output, func() tea.Cmd { return nil })
			})

		case env.Branches.MergeOptions.Msg:
			return m, popups.Cmd("pick", m.Roller.Name, "choose a merge option for "+m.Roller.Name, func() map[string]picker.Pick {
				return map[string]picker.Pick{
					"a": picker.GetPick("merge", "--abort"),
					"c": picker.GetPick("merge", "--continue"),
					"C": picker.GetPick("merge", "--no-commit", m.Roller.Name),
					"s": picker.GetPick("merge", "--squash", m.Roller.Name),
					"V": picker.GetPick("merge", "--no-verify", m.Roller.Name),
				}
			})

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
