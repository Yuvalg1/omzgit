package files

import (
	"slices"

	"omzgit/env"
	"omzgit/git"
	"omzgit/lib/list"
	"omzgit/messages/mode"
	"omzgit/messages/refresh"
	"omzgit/messages/tick"
	"omzgit/popups/help"
	"omzgit/program/files/diff"
	"omzgit/program/files/row"
	"omzgit/program/popups"
	"omzgit/roller"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case refresh.Msg:
		snapshot := m.getSnapshot()
		return m, list.Cmd(func() []row.Model { return GetFilesChanged(snapshot) }, m.list.ActiveRow, "", m.CokeCmd())

	case list.Msg[row.Model]:
		m.list.Children = msg.Children
		m.total = msg.Total
		m.list.ActiveRow = min(len(m.list.Children)-1, msg.Active)
		m.list.Children[m.list.ActiveRow].Active = true

		width, height := m.getDiffAxis()
		m.diff = diff.InitialModel(*m.list.GetCurrent(), width, height)

		res, cmd := m.list.Update(msg.Msg)
		m.list = res.(list.Model[row.Model])

		return m, tea.Batch(cmd, msg.Cmd)

	case tick.Msg:
		snapshot := m.getSnapshot()
		return m, list.Cmd(func() []row.Model { return GetFilesChanged(snapshot) }, m.list.ActiveRow, "", tick.Cmd(m.list.Children[m.list.ActiveRow].Roller.Offset))

	case roller.Msg:
		res, cmd := m.list.UpdateCurrent(msg)
		m.list = res

		return m, cmd

	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)

		dWidth, dHeight := m.getDiffAxis()
		m.diff = diff.InitialModel(*m.list.GetCurrent(), dWidth, dHeight)

		lWidth, lHeight := m.getFilesAxis()
		msg.Width = lWidth
		msg.Height = lHeight

		res, cmd := m.list.Update(msg)
		m.list = res.(list.Model[row.Model])

		return m, cmd

	case mode.Msg:
		res, cmd := m.list.Update(msg)
		m.list = res.(list.Model[row.Model])
		return m, cmd

	case tea.KeyMsg:
		if m.list.TextInput.Focused() {
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[row.Model])

			return m, cmd
		}

		switch keypress := msg.String(); keypress {
		case env.Files.AddAll.Msg:
			index := slices.IndexFunc(m.list.Children, func(row row.Model) bool { return row.Conflict })

			if index == -1 {
				git.Exec("add", "--all")
			} else {
				return m, popups.Cmd("discard", "add", "All Files", func() tea.Cmd {
					git.Exec("add", "--all")
					return m.updateChildren(msg)
				})
			}

			return m, tea.Batch(m.updateChildren(msg), m.CokeCmd())

		case env.Files.Commit.Msg:
			return m, popups.Cmd("commit", "Commit", "Commit Message	", func() tea.Cmd { return nil })

		case env.Files.DiscardAll.Msg:
			return m, popups.Cmd("discard", "discard", "All Files", func() tea.Cmd {
				git.Exec("reset", "--hard")
				return nil
			})

		case env.Files.ResetAll.Msg:
			git.Exec("reset")

			res1, cmd1 := m.list.UpdateContent(msg)
			m.list = res1

			res2, cmd2 := m.diff.Update(msg)
			m.diff = res2.(diff.Model)

			return m, tea.Batch(cmd1, cmd2)

		case env.Files.Refresh.Msg, env.Files.Search.Msg:
			m.list.SetContent(GetFilesChanged(m.getSnapshot()))
			m.total = len(m.list.Children)

			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[row.Model])

			width, height := m.getDiffAxis()
			m.diff = diff.InitialModel(*m.list.GetCurrent(), width, height)

			return m, tea.Batch(cmd, m.CokeCmd())

		case "?":
			return m, popups.Cmd("help", "", "", func() []env.Option {
				return append(help.GetEnvOptions(env.Files), help.GetEnvOptions(env.Program)...)
			})

		case env.Files.PgDown.Msg, env.Files.PgUp.Msg:
			res, cmd := m.diff.Update(msg)
			m.diff = res.(diff.Model)
			return m, cmd

		default:
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[row.Model])

			width, height := m.getDiffAxis()
			m.diff = diff.InitialModel(*m.list.GetCurrent(), width, height)

			return m, tea.Batch(cmd, m.CokeCmd())
		}
	}

	return m, nil
}

func (m *Model) updateChildren(msg tea.Msg) tea.Cmd {
	res1, cmd1 := m.list.UpdateContent(msg)
	m.list = res1

	res2, cmd2 := m.diff.Update(msg)
	m.diff = res2.(diff.Model)

	return tea.Batch(cmd1, cmd2)
}
