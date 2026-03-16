package diff

import (
	"omzgit/env"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Width = getWidth(msg.Width)
		m.viewport.Height = getHeight(msg.Height)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case env.Files.PgDown.Msg:
			m.viewport.ScrollDown(1)
			return m, nil

		case env.Files.PgUp.Msg:
			m.viewport.ScrollUp(1)
			return m, nil

		default:
			return m, nil
		}
	}

	return m, nil
}
