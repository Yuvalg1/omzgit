package branches

import (
	"os/exec"
	"program/lib/list"
	"program/messages"
	"program/program/branches/branch"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.TerminalMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)

		var cmds []tea.Cmd
		for index, element := range m.list.Children {
			res, cmd := element.Update(msg)
			cmds = append(cmds, cmd)
			m.list.Children[index] = res.(branch.Model)
		}

		return m, tea.Batch(cmds...)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "esc":
			m.list.TextInput.SetValue("")
			m.list.SetContent(getBranches(m.width, m.height))

			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[branch.Model])

			return m, cmd

		case "c":
			if gitCheckout(m.list.GetCurrent().Name) {
				current := slices.IndexFunc(m.list.Children, func(branch branch.Model) bool { return branch.Current })

				if current != -1 {
					m.list.Children[current].Current = false
				}

				m.list.Children[m.list.ActiveRow].Current = true
			}
			return m, nil

		default:
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[branch.Model])

			return m, cmd
		}
	}

	return m, nil
}

func gitCheckout(branch string) bool {
	cmd := exec.Command("git", "checkout", branch)

	_, err := cmd.Output()

	return err != nil
}
