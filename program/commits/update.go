package commits

import (
	"omzgit/env"
	"omzgit/git"
	"omzgit/lib/list"
	"omzgit/messages/mode"
	"omzgit/messages/refresh"
	"omzgit/popups/help"
	"omzgit/popups/picker"
	"omzgit/program/commits/log"
	"omzgit/program/popups"
	"omzgit/roller"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case refresh.Msg:
		snapshot := m.getSnapshot()
		return m, list.Cmd(func() []log.Model { return getCommitLogs(snapshot) }, m.list.ActiveRow, "", m.CokeCmd())

	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)

		msg.Width = m.width
		msg.Height = m.height

		res, cmd := m.list.Update(msg)
		m.list = res.(list.Model[log.Model])

		return m, cmd

	case list.Msg[log.Model]:
		m.list.Children = msg.Children
		m.total = msg.Total
		m.list.ActiveRow = min(len(m.list.Children)-1, msg.Active)
		m.list.Children[m.list.ActiveRow].Active = true

		res, cmd := m.list.Update(msg.Msg)
		m.list = res.(list.Model[log.Model])

		return m, tea.Batch(cmd, msg.Cmd)

	case roller.Msg:
		res, cmd := m.list.UpdateCurrent(msg)
		m.list = res

		return m, cmd

	case mode.Msg:
		res, cmd := m.list.Update(msg)
		m.list = res.(list.Model[log.Model])
		return m, cmd

	case tea.KeyMsg:
		if m.list.TextInput.Focused() {
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[log.Model])

			return m, cmd
		}

		switch keypress := msg.String(); keypress {
		case env.Commits.Reset.Msg:
			current := m.list.GetCurrent().Hash

			return m, popups.Cmd("pick", current, "choose a reset type for "+current, func() map[string]picker.Pick {
				return map[string]picker.Pick{
					"s": picker.GetPick("reset", "--soft", current),
					"h": picker.GetPick("reset", "--hard", current),
					"m": picker.GetPick("reset", "--mixed", current),
				}
			})

		case env.Program.Help.Msg, env.Program.Help.AltMsg:
			return m, popups.Cmd("help", "", "", func() ([]env.Option, func() tea.Cmd) {
				return append(help.GetEnvOptions(env.Commits), help.GetEnvOptions(env.Program)...),
					func() tea.Cmd {
						return nil
					}
			})

		case env.Commits.Refresh.Msg, env.Commits.Search.Msg:
			m.list.SetContent(getCommitLogs(m.getSnapshot()))
			m.total = len(m.list.Children)

			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[log.Model])

			return m, tea.Batch(cmd, m.CokeCmd())

		case env.Commits.CherryPick.Msg:
			return m, popups.Cmd("pick", "Cherry Pick Options", "choose a cherry-pick option", func() map[string]picker.Pick {
				return map[string]picker.Pick{
					"p": {Desc: "cherry-pick", Callback: func() tea.Cmd {
						return popups.Cmd("input", "Cherry Pick", "enter a cherry pick message", func(message string) tea.Cmd {
							output, err := git.Exec("cherry-pick", message)
							if err != nil {
								return popups.Cmd("alert", "Cherry Pick Error!", output, func(name string) {})
							}

							return nil
						})
					}},
					"a": picker.GetPick("cherry-pick", "--abort"),
					"c": picker.GetPick("cherry-pick", "--continue"),
					"s": picker.GetPick("cherry-pick", "--skip"),
					"q": picker.GetPick("cherry-pick", "--quit"),
				}
			})

		default:
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[log.Model])

			return m, tea.Batch(cmd, m.CokeCmd())
		}
	}

	return m, nil
}
