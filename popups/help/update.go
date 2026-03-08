package help

import (
	"omzgit/env"
	"omzgit/lib/list"
	"omzgit/messages/mode"
	"omzgit/messages/refresh"
	"omzgit/popups/help/option"
	"omzgit/program/popups"
	"omzgit/roller"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case refresh.Msg:
		m.list.SetContent(m.getOptions())
		m.list.GetCurrent().Active = true

		return m, nil

	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)

		msg.Width = getWidth(msg.Width)
		msg.Height = getHeight(msg.Height)

		res, cmd := m.list.Update(msg)
		m.list = res.(list.Model[option.Model])

		return m, cmd

	case popups.Msg:
		m.defaultOptions = msg.Fn.(func() []env.Option)()
		m.list.SetContent(m.getOptions())
		m.list.Children[m.list.ActiveRow].Active = true
		m.visible = true
		return m, nil

	case roller.Msg:
		res, cmd := m.list.UpdateCurrent(msg)
		m.list = res

		return m, cmd

	case mode.Msg:
		res, cmd := m.list.Update(msg)
		m.list = res.(list.Model[option.Model])
		return m, cmd

	case tea.KeyMsg:
		if m.list.TextInput.Focused() {
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[option.Model])

			return m, cmd
		}

		switch keypress := msg.String(); keypress {

		case "esc":
			m.list.SetContent(m.getOptions())
			m.visible = false

			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[option.Model])

			return m, cmd

		case "/":
			m.list.SetContent(m.getOptions())

			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[option.Model])

			return m, cmd

		case "ctrl+c", "q":
			return m, tea.Quit

		default:
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[option.Model])

			return m, cmd
		}
	}

	return m, nil
}
