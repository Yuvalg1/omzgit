package branches

import (
	"omzgit/git"
	"omzgit/lib/list"
	"omzgit/messages"
	"omzgit/program/branches/branch"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.RefreshMsg:
		m.list.SetContent(getBranches(m.width, m.height))

		m.list.ActiveRow = slices.IndexFunc(m.list.Children, func(branch branch.Model) bool { return branch.Current })
		current := m.list.GetCurrent()

		if current == nil {
			return m, nil
		}

		current.Active = true
		return m, nil

	case messages.TerminalMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)

		msg.Width = getWidth(msg.Width)
		msg.Height = getHeight(msg.Height)

		res, cmd := m.list.Update(msg)
		m.list = res.(list.Model[branch.Model])

		return m, cmd

	case tea.KeyMsg:
		if m.list.TextInput.Focused() {
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[branch.Model])
			return m, cmd
		}

		switch keypress := msg.String(); keypress {

		case "b":
			return m, m.PopupCmd("input", "Name", "Enter A new Branch Name", func(name string) {
				git.Exec("checkout", "-b", name)
			})

		case "c":
			output, err := git.Exec("checkout", m.list.GetCurrent().Name)
			if err == nil {
				current := slices.IndexFunc(m.list.Children, func(branch branch.Model) bool { return branch.Current })

				if current != -1 {
					m.list.Children[current].Current = false
				}

				m.list.Children[m.list.ActiveRow].Current = true
				return m, nil
			}
			return m, m.PopupCmd("alert", "Alert!", output, func(name string) {})

		case "d":
			return m, m.PopupCmd("discard", "delete", m.list.GetCurrent().Name, func() tea.Cmd {
				output, err := git.Exec("branch", "-d", m.list.GetCurrent().Name)
				if err == nil {
					return nil
				}

				return m.PopupCmd("alert", "Delete Error", output, func(name string) {})
			})

		case "D":
			return m, m.PopupCmd("discard", "force delete", m.list.GetCurrent().Name, func() tea.Cmd {
				output, err := git.Exec("branch", "-D", m.list.GetCurrent().Name)
				if err == nil {
					return nil
				}

				return m.PopupCmd("alert", "Force Delete Error", output, func(name string) {})
			})

		case "esc":
			m.list.TextInput.SetValue("")
			m.list.SetContent(getBranches(m.width, m.height))

			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[branch.Model])

			return m, tea.Batch(cmd, m.CokeCmd())

		default:
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[branch.Model])

			return m, tea.Batch(cmd, m.CokeCmd())
		}
	}

	return m, nil
}
