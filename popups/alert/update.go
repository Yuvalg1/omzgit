package alert

import (
	"strings"

	"omzgit/clipboard"
	"omzgit/env"
	"omzgit/popups/help"
	"omzgit/program/popups"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.maxHeight = getHeight(msg.Height)
		m.viewport.Width = getWidth(msg.Width)

		error := m.getContentStyle().
			Render(m.error)

		m.viewport.Height = min(
			getHeight(msg.Height),
			strings.Count(error, "\n")+1,
		)

		m.viewport.SetContent(error)

		return m, nil

	case popups.Msg:
		m.error = msg.Name[:len(msg.Name)-1]
		error := m.getContentStyle().
			Render(m.error)

		m.viewport.Height = min(
			m.maxHeight,
			strings.Count(error, "\n")+1,
		)

		m.viewport.SetContent(error)
		m.visible = true
		m.verb = msg.Verb
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case env.Alert.CtrlC.Msg, env.Alert.Quit.Msg:
			return m, tea.Quit

		case "?":
			return m, popups.Cmd("help", "", "", func() ([]env.Option, func() tea.Cmd) {
				return help.GetEnvOptions(env.Alert),
					func() tea.Cmd {
						return popups.Cmd("alert", m.verb, m.error, func() tea.Cmd { return nil })
					}
			})

		case env.Alert.Up.Msg, env.Alert.Up.Msg:
			m.viewport.ScrollUp(1)
			return m, nil

		case env.Alert.Down.Msg, env.Alert.Down.Msg:
			m.viewport.ScrollDown(1)
			return m, nil

		case env.Alert.PgDown.Msg:
			m.viewport.PageDown()
			return m, nil

		case env.Alert.PgUp.Msg:
			m.viewport.PageUp()
			return m, nil

		case env.Alert.Yank.Msg:
			clipboard.Copy(m.error)
			return m, nil

		default:
			m.visible = false
			return m, nil
		}
	}

	return m, nil
}
