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

		ourWidth, ourHeight := m.getOurAxis()
		msg.Width = ourWidth
		msg.Height = ourHeight

		res1, cmd1 := m.ours.Update(msg)
		m.ours = res1.(content.Model)

		theirWidth, theirHeight := m.getTheirAxis()
		msg.Width = theirWidth
		msg.Height = theirHeight

		res2, cmd2 := m.theirs.Update(msg)
		m.theirs = res2.(content.Model)

		return m, tea.Batch(cmd1, cmd2)

	case content.Msg:
		m.resolve(msg.Index, msg.Ours)
		m.ours.Refresh()
		m.theirs.Refresh()
		return m, nil

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

		case "o":
			res, cmd := m.ours.Update(msg)
			m.ours = res.(content.Model)
			return m, cmd

		case "t":
			res, cmd := m.theirs.Update(msg)
			m.theirs = res.(content.Model)
			return m, cmd

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
