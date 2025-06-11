package discard

import (
	"program/messages"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.TerminalMsg:
		m.Width = getWidth(msg.Width)
		m.Height = getHeight(msg.Height)
		return m, nil

	case messages.PopupMsg:
		m.CallbackFn = msg.Fn.(func() bool)
		m.Name = msg.Name
		m.visible = true
		m.verb = msg.Verb
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "esc", "n", "N":
			m.visible = false
			return m, nil

		case "y", "Y":
			if !m.CallbackFn() {
				return m, m.PopupCmd("alert", cases.Title(language.English).String(m.verb)+" Error!", "Could not "+m.verb+" requested item.", func() bool { return true })
			}
			m.visible = false
			return m, m.RefreshCmd()

		case "ctrl+c", "q":
			return m, tea.Quit

		default:
			return m, nil
		}
	}

	return m, nil
}
