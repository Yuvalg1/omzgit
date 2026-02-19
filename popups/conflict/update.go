package conflict

import (
	"omzgit/program/popups"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case popups.Msg:
		m.visible = true
		m.ours = getContent(msg.Verb)
		m.theirs = getContent(msg.Verb)

		return m, nil

	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)

		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		default:
			return m, nil
		}
	}

	return m, nil
}
