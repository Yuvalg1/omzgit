package files

import (
	"slices"

	"omzgit/git"
	"omzgit/lib/list"
	"omzgit/messages/refresh"
	"omzgit/messages/tick"
	"omzgit/program/files/diff"
	"omzgit/program/files/row"
	"omzgit/program/popups"
	"omzgit/roller"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case refresh.Msg:
		m.list.SetContent(m.GetFilesChanged())
		m.list.GetCurrent().Active = true

		m.diffs = getDiffs(m.list.Children, m.width, m.height)

		return m, nil

	case tick.Msg:
		m.list.SetContent(m.GetFilesChanged())

		current := m.list.GetCurrent()

		if current == nil {
			return m, tick.Cmd(m.list.Children[m.list.ActiveRow].Roller.Offset)
		}

		res, cmd := m.list.UpdateCurrent(msg)
		m.list = res

		m.diffs[m.list.ActiveRow] = diff.InitialModel(*m.list.GetCurrent(), m.width, m.height)

		return m, tea.Batch(cmd, tick.Cmd(m.list.Children[m.list.ActiveRow].Roller.Offset))

	case roller.Msg:
		res, cmd := m.list.UpdateCurrent(msg)
		m.list = res

		return m, cmd

	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)

		msg.Width = getWidth(msg.Width)
		msg.Height = getHeight(msg.Height)

		var cmds []tea.Cmd
		for index, element := range m.diffs {
			res, cmd := element.Update(msg)
			m.diffs[index] = res.(diff.Model)

			cmds = append(cmds, cmd)
		}

		res, cmd := m.list.Update(msg)
		m.list = res.(list.Model[row.Model])
		cmds = append(cmds, cmd)

		return m, tea.Batch(cmds...)

	case tea.KeyMsg:
		if m.list.TextInput.Focused() {
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[row.Model])

			return m, cmd
		}

		switch keypress := msg.String(); keypress {
		case "a", "r":
			res1, cmd1 := m.list.UpdateCurrent(msg)
			m.list = res1

			res2, cmd2 := m.diffs[m.list.ActiveRow].Update(msg)
			m.diffs[m.list.ActiveRow] = res2.(diff.Model)

			return m, tea.Batch(cmd1, cmd2, m.CokeCmd())

		case "A":
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

		case "c":
			return m, popups.Cmd("commit", "Commit", "Commit Message	", func() tea.Cmd { return nil })

		case "D":
			return m, popups.Cmd("discard", "discard", "All Files", func() tea.Cmd {
				git.Exec("reset", "--hard")
				return nil
			})

		case "R":
			git.Exec("reset")

			cmds := []tea.Cmd{m.CokeCmd()}

			res, cmd := m.list.UpdateContent(msg)
			m.list = res
			cmds = append(cmds, cmd)

			for index, element := range m.diffs {
				res, cmd := element.Update(msg)
				m.diffs[index] = res.(diff.Model)

				cmds = append(cmds, cmd)
			}

			return m, tea.Batch(cmds...)

		case "esc", "/":
			m.list.SetContent(m.GetFilesChanged())

			res1, cmd1 := m.list.Update(msg)
			m.list = res1.(list.Model[row.Model])

			m.diffs = getDiffs(m.list.Children, m.width, m.height)

			res2, cmd2 := m.diffs[m.list.ActiveRow].Update(msg)
			m.diffs[m.list.ActiveRow] = res2.(diff.Model)

			return m, tea.Batch(cmd1, cmd2, m.CokeCmd())

		default:
			res1, cmd1 := m.list.Update(msg)
			m.list = res1.(list.Model[row.Model])

			res2, cmd2 := m.diffs[m.list.ActiveRow].Update(msg)
			m.diffs[m.list.ActiveRow] = res2.(diff.Model)

			return m, tea.Batch(cmd1, cmd2, m.CokeCmd())
		}
	}

	return m, nil
}

func (m *Model) updateChildren(msg tea.Msg) tea.Cmd {
	cmds := []tea.Cmd{m.CokeCmd()}

	res, cmd := m.list.UpdateContent(msg)
	m.list = res
	cmds = append(cmds, cmd)

	for index, element := range m.diffs {
		res, cmd := element.Update(msg)
		m.diffs[index] = res.(diff.Model)

		cmds = append(cmds, cmd)
	}

	return tea.Batch(cmds...)
}
