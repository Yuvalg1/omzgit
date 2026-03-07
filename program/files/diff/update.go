package diff

import (
	"omzgit/env"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)

		msg.Width = getWidth(msg.Width)
		msg.Height = getHeight(msg.Height)

		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case env.Files.Add.Msg, env.Files.AddAll.Msg:
			if m.Staged {
				return m, nil
			}

			m.Staged = true
			m.content = m.getDiff()
			return m, nil

		case env.Files.Reset.Msg, env.Files.ResetAll.Msg:
			if !m.Staged {
				return m, nil
			}
			m.Staged = false
			m.content = m.getDiff()
			return m, nil

		case env.Files.Up.Msg, env.Files.Up.AltMsg, env.Files.Down.Msg, env.Files.Down.AltMsg:
			m.content = m.getDiff()
			return m, nil

		default:
			return m, nil
		}
	}

	return m, nil
}
