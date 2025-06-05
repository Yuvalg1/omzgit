package files

import (
	"os/exec"
	"program/lib/list"
	"program/messages"
	"program/program/files/diff"
	"program/program/files/row"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.DeletedMsg:
		m.list.SetContent(GetFilesChanged(m.Width))

		current := m.list.GetCurrent()

		if current == nil {
			return m, nil
		}

		current.Active = true

		m.Diffs[m.list.ActiveRow].Content = m.Diffs[m.list.ActiveRow].GetContent()
		return m, nil

	case messages.TickMsg:
		m.list.SetContent(GetFilesChanged(m.Width))

		current := m.list.GetCurrent()

		if current == nil {
			return m, m.TickCmd()
		}

		current.Active = true

		m.Diffs[m.list.ActiveRow].Content = m.Diffs[m.list.ActiveRow].GetContent()
		return m, m.TickCmd()

	case messages.TerminalMsg:
		var cmds []tea.Cmd

		m.Width = getWidth(msg.Width)
		m.Height = getHeight(msg.Height)

		msg.Width = getWidth(msg.Width)
		msg.Height = getHeight(msg.Height)

		for index, element := range m.Diffs {
			res, cmd := element.Update(msg)
			m.Diffs[index] = res.(diff.Model)

			cmds = append(cmds, cmd)
		}

		res, cmd := m.list.Update(msg)
		m.list = res.(list.Model[row.Model])
		cmds = append(cmds, cmd)

		return m, tea.Batch(cmds...)

	case messages.PopupMsg:
		m = InitialModel(m.Width, m.Height+2)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "A":
			if !gitAddAll() {
				return m, nil
			}

			var cmds []tea.Cmd
			cmds = append(cmds, m.CokeCmd(""))

			res, cmd := m.list.UpdateContent(msg)
			m.list = res
			cmds = append(cmds, cmd)

			for index, element := range m.Diffs {
				res, cmd := element.Update(msg)
				m.Diffs[index] = res.(diff.Model)

				cmds = append(cmds, cmd)
			}

			return m, tea.Batch(cmds...)

		case "esc":
			m.list.TextInput.SetValue("")
			m.list.SetContent(GetFilesChanged(m.Width))
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[row.Model])
			return m, cmd

		case "d":
			res, cmd := m.list.UpdateCurrent(msg)
			m.list = res

			m.list.SetContent(GetFilesChanged(m.Width))
			return m, cmd

		case "D":
			return m, m.PopupCmd("discard", "All Files", func() {
				gitRestoreAll()
			})

		case "R":
			if !gitResetAll() {
				return m, nil
			}

			var cmds []tea.Cmd
			cmds = append(cmds, m.CokeCmd(""))

			res, cmd := m.list.UpdateContent(msg)
			m.list = res
			cmds = append(cmds, cmd)

			for index, element := range m.Diffs {
				res, cmd := element.Update(msg)
				m.Diffs[index] = res.(diff.Model)

				cmds = append(cmds, cmd)
			}

			return m, tea.Batch(cmds...)

		default:
			res1, cmd1 := m.list.Update(msg)
			m.list = res1.(list.Model[row.Model])

			res2, cmd2 := m.Diffs[m.list.ActiveRow].Update(msg)
			m.Diffs[m.list.ActiveRow] = res2.(diff.Model)

			return m, tea.Batch(cmd1, cmd2, m.CokeCmd(""))
		}
	}

	return m, nil
}

func gitAddAll() bool {
	cmd := exec.Command("git", "add", "--all")
	_, err := cmd.Output()

	return err == nil
}

func gitResetAll() bool {
	cmd := exec.Command("git", "reset")
	_, err := cmd.Output()

	return err == nil
}

func gitRestoreAll() {
	cmd := exec.Command("git", "reset", "--hard")
	cmd.Output()
}
