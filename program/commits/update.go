package commits

import (
	"strconv"

	"omzgit/lib/list"
	"omzgit/messages"
	"omzgit/program/commits/log"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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
		switch keypress := msg.String(); keypress {

		case "r":
			return m, m.PopupCmd("reset", "reset", "HEAD~"+strconv.Itoa(m.list.ActiveRow+1), func() {})

		case "esc":
			m.list.TextInput.SetValue("")
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
