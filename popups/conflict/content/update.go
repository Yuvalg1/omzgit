package content

import (
	"omzgit/popups/conflict/chunk"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Content.Width = getWidth(msg.Width)
		m.Content.Height = getHeight(msg.Height)

		msg.Width = getWidth(msg.Width)
		msg.Height = getHeight(msg.Height)

		cmds := []tea.Cmd{}
		for index, element := range m.conflicts {
			res, cmd := element.Update(msg)
			m.conflicts[index] = res.(chunk.Model)
			cmds = append(cmds, cmd)
		}
		m.Refresh()

		return m, tea.Batch(cmds...)

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

		case "n":
			if m.sum > 0 {
				m.index = (m.index + 1) % m.sum
			}
			m.Refresh()
			return m, nil

		case "N":
			if m.sum > 0 {
				m.index = (m.index - 1 + m.sum) % m.sum
			}
			m.Refresh()
			return m, nil

		case "o", "t":
			return m, Cmd(m.index, m.ours)

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
