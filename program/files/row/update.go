package row

import (
	"omzgit/git"
	"omzgit/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)

		msg.Width = getWidth(msg.Width)

		res, cmd := m.Roller.Update(msg)
		m.Roller = res

		return m, cmd

	case messages.RefreshMsg:
		m.Active = true

		return m, nil

	case messages.TickMsg:
		m.Active = true
		m.Roller.Offset = msg.RollOffset

		return m, nil

	case messages.RollerMsg:
		res, cmd := m.Roller.Update(msg)
		m.Roller = res
		return m, cmd

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "a":
			if !m.Staged {
				_, err := git.Exec("add", m.Roller.Name)
				m.Staged = err == nil
			}
			return m, nil

		case "A":
			m.Staged = true
			return m, nil

		case "enter":
			m.Active = true
			return m, nil

		case "j", "k", "down", "up", "g", "G", "/", "esc":
			m.Active = !m.Active

			res, cmd := m.Roller.Update(msg)
			m.Roller = res

			return m, cmd

		case "d":
			return m, m.PopupCmd("discard", "discard", m.Roller.Name, func() tea.Cmd {
				if m.Staged {
					_, err := git.Exec("reset", "--", m.Roller.Name)
					m.Staged = err != nil
				}
				git.Exec("restore", m.Roller.Name)
				return nil
			})

		case "r":
			if m.Staged {
				_, err := git.Exec("reset", "--", m.Roller.Name)
				m.Staged = err != nil
			}
			return m, nil

		case "R":
			m.Staged = false
			return m, nil

		default:
			res, cmd := m.Roller.Update(msg)
			m.Roller = res

			return m, cmd
		}
	}

	return m, nil
}
