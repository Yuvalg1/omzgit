package files

import (
	"omzgit/git"
	"omzgit/lib/list"
	"omzgit/messages"
	"omzgit/program/files/diff"
	"omzgit/program/files/row"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.RefreshMsg:
		m.list.SetContent(GetFilesChanged(m.width))

		current := m.list.GetCurrent()

		if current == nil {
			return m, nil
		}

		current.Active = true

		m.diffs[m.list.ActiveRow] = diff.InitialModel(*m.list.GetCurrent(), m.width, m.height)

		return m, nil

	case messages.TickMsg:
		m.list.SetContent(GetFilesChanged(m.width))

		current := m.list.GetCurrent()

		if current == nil {
			return m, m.TickCmd()
		}

		current.Active = true

		m.diffs[m.list.ActiveRow] = diff.InitialModel(*m.list.GetCurrent(), m.width, m.height)

		return m, m.TickCmd()

	case tea.WindowSizeMsg:
		var cmds []tea.Cmd

		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)

		msg.Width = getWidth(msg.Width)
		msg.Height = getHeight(msg.Height)

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
			git.Exec("add", "--all")

			var cmds []tea.Cmd
			cmds = append(cmds, m.CokeCmd())

			res, cmd := m.list.UpdateContent(msg)
			m.list = res
			cmds = append(cmds, cmd)

			for index, element := range m.diffs {
				res, cmd := element.Update(msg)
				m.diffs[index] = res.(diff.Model)

				cmds = append(cmds, cmd)
			}

			return m, tea.Batch(cmds...)

		case "esc":
			m.list.TextInput.SetValue("")
			m.list.SetContent(GetFilesChanged(m.width))
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[row.Model])
			return m, tea.Batch(cmd, m.CokeCmd())

		case "c":
			return m, m.PopupCmd("commit", "Commit", "Commit Message	", func() tea.Cmd { return nil })

		case "d":
			res, cmd := m.list.UpdateCurrent(msg)
			m.list = res
			return m, cmd

		case "D":
			return m, m.PopupCmd("discard", "discard", "All Files", func() tea.Cmd {
				git.Exec("reset", "--hard")
				return nil
			})

		case "R":
			git.Exec("reset")

			var cmds []tea.Cmd
			cmds = append(cmds, m.CokeCmd())

			res, cmd := m.list.UpdateContent(msg)
			m.list = res
			cmds = append(cmds, cmd)

			for index, element := range m.diffs {
				res, cmd := element.Update(msg)
				m.diffs[index] = res.(diff.Model)

				cmds = append(cmds, cmd)
			}

			return m, tea.Batch(cmds...)

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
