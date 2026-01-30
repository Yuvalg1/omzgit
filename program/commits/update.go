package commits

import (
	"slices"
	"strconv"

	"omzgit/git"
	"omzgit/lib/list"
	"omzgit/messages"
	"omzgit/program/commits/log"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.RefreshMsg:
		m.list.SetContent(getCommitLogs(m.width))

		res, cmd := m.list.Update(msg)
		m.list = res.(list.Model[log.Model])

		return m, cmd

	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)

		msg.Width = m.width
		msg.Height = m.height

		res, cmd := m.list.Update(msg)
		m.list = res.(list.Model[log.Model])

		return m, cmd

	case messages.RollerMsg:
		res, cmd := m.list.UpdateCurrent(msg)
		m.list = res

		return m, cmd

	case tea.KeyMsg:
		if m.list.TextInput.Focused() {
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[log.Model])
			return m, cmd
		}

		switch keypress := msg.String(); keypress {
		case "c":
			output, err := git.Exec("checkout", m.list.GetCurrent().Hash)
			if err != nil {
				return m, m.PopupCmd("alert", "Alert!", output, func(name string) {})
			}

			current := slices.IndexFunc(m.list.Children, func(log log.Model) bool { return log.Current })

			if current != -1 {
				m.list.Children[current].Current = false
			}

			m.list.Children[m.list.ActiveRow].Current = true
			return m, nil

		case "C":
			return m, m.PopupCmd("discard", "force checkout", m.list.GetCurrent().Hash, func() tea.Cmd {
				output, err := git.Exec("checkout", "-f", m.list.GetCurrent().Hash)
				if err != nil {
					return m.PopupCmd("alert", "checkout error", output, func(name string) {})
				}

				current := slices.IndexFunc(m.list.Children, func(log log.Model) bool { return log.Current })

				if current != -1 {
					m.list.Children[current].Current = false
				}

				m.list.Children[m.list.ActiveRow].Current = true
				return nil
			})

		case "r":
			return m, m.PopupCmd("reset", m.list.GetCurrent().Hash, "HEAD~"+strconv.Itoa(m.list.ActiveRow+1), func() {})

		case "esc", "/":
			m.list.SetContent(getCommitLogs(m.width))

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
