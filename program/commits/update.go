package commits

import (
	"strconv"

	"omzgit/env"
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
			current := "HEAD~" + strconv.Itoa(m.list.ActiveRow+1)

			return m, popups.Cmd("pick", current, "choose a reset type for "+m.list.GetCurrent().Hash, func() map[string]picker.Pick {
				return map[string]picker.Pick{
					"s": picker.GetPick("reset", "--soft", current),
					"h": picker.GetPick("reset", "--hard", current),
					"m": picker.GetPick("reset", "--mixed", current),
				}
			})

		case "?":
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

		default:
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[log.Model])

			return m, tea.Batch(cmd, m.CokeCmd())
		}
	}

	return m, nil
}
