package row

import (
	"omzgit/git"
	"omzgit/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.TerminalMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)

		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "a":
			if !m.Staged {
				m.Staged = git.Exec("add", m.Path)
			}
			return m, nil

		case "A":
			m.Staged = true
			return m, nil

		case "enter":
			m.Active = true
			return m, nil

		case "g", "G", "/", "esc":
			m.Active = !m.Active
			return m, nil

		case "j", "k", "down", "up":
			m.Active = !m.Active
			return m, nil

		case "d":
			return m, m.PopupCmd("discard", "discard", m.Path, func() tea.Cmd {
				if m.Staged {
					m.Staged = !git.Exec("reset", "--", m.Path)
				}
				git.Exec("restore", m.Path)
				return nil
			})

		case "r":
			if m.Staged {
				m.Staged = !git.Exec("reset", "--", m.Path)
			}
			return m, nil

		case "R":
			m.Staged = false
			return m, nil

		default:
			return m, nil
		}
	}

	return m, nil
}
