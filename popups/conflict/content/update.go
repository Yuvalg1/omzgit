package content

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Content.Width = getWidth(msg.Width)
		m.Content.Height = getHeight(msg.Height)

		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "j", "down":
			m.Content.ScrollDown(1)
			return m, nil

		case "k", "up":
			m.Content.ScrollUp(1)
			return m, nil

		case "pgdown":
			m.Content.PageDown()
			return m, nil

		case "pgup":
			m.Content.PageUp()
			return m, nil

		default:
			return m, nil
		}
	}

	return m, nil
}
