package conflict

import (
	"omzgit/env"
	"omzgit/git"
	"omzgit/messages/refresh"
	"omzgit/popups/conflict/content"
	"omzgit/program/popups"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case popups.Msg:
		m.visible = true
		m.path = msg.Name
		m.getContent()

		return m, nil

	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)

		if m.width > CUTOFF {
			msg.Width = m.getOurAxis()
			msg.Height = m.height - 1
		} else {
			msg.Width = m.width
			msg.Height = m.getOurAxis()
		}

		res1, cmd1 := m.ours.Update(msg)
		m.ours = res1.(content.Model)

		if m.width > CUTOFF {
			msg.Width = m.getTheirAxis()
			msg.Height = m.height - 1
		} else {
			msg.Width = m.width
			msg.Height = m.getTheirAxis()
		}

		res2, cmd2 := m.theirs.Update(msg)
		m.theirs = res2.(content.Model)

		return m, tea.Batch(cmd1, cmd2)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "esc":
			m.visible = false
			return m, nil

		case env.Files.Ours.Msg:
			git.Exec("checkout", "--ours", m.path)
			git.Exec("add", m.path)
			m.visible = false
			return m, refresh.Cmd()

		case env.Files.Theirs.Msg:
			git.Exec("checkout", "--theirs", m.path)
			git.Exec("add", m.path)
			m.visible = false
			return m, refresh.Cmd()

		default:
			res1, cmd1 := m.ours.Update(msg)
			m.ours = res1.(content.Model)

			res2, cmd2 := m.theirs.Update(msg)
			m.theirs = res2.(content.Model)

			return m, tea.Batch(cmd1, cmd2)
		}
	}

	return m, nil
}
