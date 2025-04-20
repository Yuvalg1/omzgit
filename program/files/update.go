package files

import (
	"os/exec"
	"program/messages"
	"program/program/files/diff"
	"program/program/files/row"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.DeletedMsg:
		m.files = GetFilesChanged((m.Width - 2) / 2)
		m.files[m.ActiveRow].Active = true

		for index, element := range m.files {
			newDiff := diff.InitialModel(element.Path, element.Staged, (m.Width-2)/2, m.Height)
			newDiff.Content = newDiff.GetContent()

			m.Diffs[index] = newDiff
		}
		return m, nil

	case messages.TickMsg:
		m.files = GetFilesChanged((m.Width - 2) / 2)
		m.files[m.ActiveRow].Active = true
		m.Diffs[m.ActiveRow].Content = m.Diffs[m.ActiveRow].GetContent()
		return m, m.TickCmd()

	case messages.TerminalMsg:
		var cmds []tea.Cmd

		m.Width = msg.Width
		m.Height = msg.Height

		msg.Width = msg.Width / 2

		for index, element := range m.Diffs {
			res, cmd := element.Update(msg)
			m.Diffs[index] = res.(diff.Model)

			cmds = append(cmds, cmd)
		}

		for index, element := range m.files {
			res, cmd := element.Update(msg)

			m.files[index] = res.(row.Model)
			cmds = append(cmds, cmd)
		}

		return m, tea.Batch(cmds...)

	case messages.PopupMsg:
		m = InitialModel(m.Width, m.Height)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "A":
			if !gitAddAll() {
				return m, nil
			}

			var cmds []tea.Cmd

			for index, element := range m.files {
				res, cmd := element.Update(msg)
				m.files[index] = res.(row.Model)

				cmds = append(cmds, cmd)
			}

			for index, element := range m.Diffs {
				res, cmd := element.Update(msg)
				m.Diffs[index] = res.(diff.Model)

				cmds = append(cmds, cmd)
			}

			return m, tea.Batch(cmds...)

		case "esc":
			m.files = GetFilesChanged((m.Width - 2) / 2)
			m.files[m.ActiveRow].Active = true
			return m, nil

		case "d":
			res, cmd := m.files[m.ActiveRow].Update(msg)
			m.files[m.ActiveRow] = res.(row.Model)
			m.files = GetFilesChanged((m.Width - 2) / 2)
			return m, cmd

		case "D":
			return m, m.PopupCmd("All Files", func() {
				gitRestoreAll()
			})

		case "g":
			cmds := move(m, msg, m.ActiveRow, 0)
			m.ActiveRow = 0
			return m, tea.Batch(cmds...)

		case "G":
			cmds := move(m, msg, m.ActiveRow, len(m.files)-1)
			m.ActiveRow = len(m.files) - 1
			return m, tea.Batch(cmds...)

		case "j", "down":
			curr := m.ActiveRow
			next := (m.ActiveRow + 1 + len(m.files)) % len(m.files)

			cmds := move(m, msg, curr, next)
			m.ActiveRow = next

			return m, tea.Batch(cmds...)

		case "k", "up":
			curr := m.ActiveRow
			next := (m.ActiveRow - 1 + len(m.files)) % len(m.files)

			cmds := move(m, msg, curr, next)
			m.ActiveRow = next

			return m, tea.Batch(cmds...)

		case "R":
			if !gitResetAll() {
				return m, nil
			}

			var cmds []tea.Cmd

			for index, element := range m.files {
				res, cmd := element.Update(msg)
				m.files[index] = res.(row.Model)
				cmds = append(cmds, cmd)
			}

			for index, element := range m.Diffs {
				res, cmd := element.Update(msg)
				m.Diffs[index] = res.(diff.Model)

				cmds = append(cmds, cmd)
			}

			return m, tea.Batch(cmds...)

		default:
			res1, cmd1 := m.files[m.ActiveRow].Update(msg)
			m.files[m.ActiveRow] = res1.(row.Model)

			res2, cmd2 := m.Diffs[m.ActiveRow].Update(msg)
			m.Diffs[m.ActiveRow] = res2.(diff.Model)

			return m, tea.Batch(cmd1, cmd2)
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

func move(m Model, msg tea.Msg, curr int, next int) []tea.Cmd {
	res1, cmd1 := m.files[curr].Update(msg)
	m.files[curr] = res1.(row.Model)

	res2, cmd2 := m.files[next].Update(msg)
	m.files[next] = res2.(row.Model)

	m.Diffs[next].Width = m.Width / 2
	m.Diffs[next].Height = m.Height

	res3, cmd3 := m.Diffs[next].Update(msg)
	m.Diffs[next] = res3.(diff.Model)

	return []tea.Cmd{cmd1, cmd2, cmd3}
}
