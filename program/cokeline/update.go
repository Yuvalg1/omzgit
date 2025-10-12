package cokeline

import (
	"omzgit/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.CokeMsg:
		m.Center = msg.Center
		m.Right = msg.Right

		m.Primary = msg.Primary
		return m, nil

	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "b":
			m.Left = "Branches"
		case "c":
			m.Left = "Commits"
		case "f":
			m.Left = "Files"
		default:
			return m, nil
		}
	}

	return m, nil
}
