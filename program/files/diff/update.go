package diff

import (
	"program/messages"
	"program/program/files/diff/name"
	"program/program/files/diff/row"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.TerminalMsg:
		m.Width = msg.Width
		m.Height = msg.Height

		var cmds []tea.Cmd
		msg.Width = msg.Width - 1

		res, cmd := m.Name.Update(msg)
		m.Name = res.(name.Model)

		cmds = append(cmds, cmd)

		for index, element := range m.Content {
			res, cmd := element.Update(msg)
			m.Content[index] = res.(row.Model)
			cmds = append(cmds, cmd)
		}

		return m, tea.Batch(cmds...)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case "a":
			if m.staged {
				return m, nil
			}

			m.staged = true
			m.Content = m.GetContent()
			return m, nil

		case "A":
			if m.staged {
				return m, nil
			}

			m.staged = true
			m.Content = m.GetContent()
			return m, nil

		case "r":
			if !m.staged {
				return m, nil
			}
			m.staged = false
			m.Content = m.GetContent()
			return m, nil

		case "R":
			if !m.staged {
				return m, nil
			}

			m.staged = false
			m.Content = m.GetContent()
			return m, nil

		case "up", "down":
			m.Content = m.GetContent()
			return m, nil

		default:
			return m, nil
		}
	}

	return m, nil
}
