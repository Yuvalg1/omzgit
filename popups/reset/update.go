package reset

import (
	"strings"

	"omzgit/git"
	"omzgit/messages/refresh"
	"omzgit/program/popups"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)
		return m, nil

	case popups.Msg:
		m.visible = true
		m.hash = msg.Verb
		m.name = msg.Name
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "esc":
			m.visible = false
			return m, nil

		case "h", "m", "s":
			output, err := git.Exec("reset", m.options[keypress[0]], m.name)
			if err != nil {
				return m, popups.Cmd("alert", "reset", strings.TrimSpace(output), func() {})
			}
			m.visible = false

			return m, refresh.Cmd()

		case "ctrl+c", "q":
			return m, tea.Quit

		default:
			return m, nil
		}
	}

	return m, nil
}
