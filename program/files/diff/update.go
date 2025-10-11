package diff

import (
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

		case "a", "A":
			if m.Staged {
				return m, nil
			}

			m.Staged = true
			m.content = m.getDiffStaged()
			return m, nil

		case "r", "R":
			if !m.Staged {
				return m, nil
			}
			m.Staged = false
			m.content = m.getDiffStaged()
			return m, nil

		case "j", "k", "up", "down":
			m.content = m.getDiffStaged()
			return m, nil

		default:
			return m, nil
		}
	}

	return m, nil
}
