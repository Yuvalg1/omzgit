package cokeline

import (
	"omzgit/env"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case Msg:
		m.Center = msg.Center
		m.Right = msg.Right

		m.Primary = msg.Primary
		return m, nil

	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case env.Goto.Branches.Msg:
			m.Left = "Branches"
		case env.Goto.Commits.Msg:
			m.Left = "Commits"
		case env.Goto.Files.Msg:
			m.Left = "Files"
		default:
			return m, nil
		}
	}

	return m, nil
}
