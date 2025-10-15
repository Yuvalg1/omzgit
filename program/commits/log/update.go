package log

import (
	"omzgit/git"
	"omzgit/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)

		res, cmd := m.Desc.Update(msg)
		m.Desc = res

		return m, cmd

	case messages.RefreshMsg:
		m.Active = true

		return m, nil

	case messages.RollerMsg:
		res, cmd := m.Desc.Update(msg)
		m.Desc = res
		return m, cmd

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "j", "k", "down", "up", "g", "G", "/", "esc":
			m.Active = !m.Active
			res, cmd := m.Desc.Update(msg)

			m.Desc = res

			return m, cmd

		case "ctrl+p":
			output, err := git.Exec("cherry-pick", m.Hash)
			if err != nil {
				m.PopupCmd("alert", "cherry pick error", output, func() {})
			}
			return m, m.RefreshCmd()

		default:
			return m, nil
		}
	}

	return m, nil
}
