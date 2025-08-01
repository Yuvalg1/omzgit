package cokeline

import (
	"omzgit/default/colors"
	"omzgit/default/colors/bg"
	"omzgit/messages"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.CokeMsg:
		m.Center = msg.Center
		m.Right = msg.Right
		return m, nil

	case messages.TerminalMsg:
		m.width = getWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "b":
			m.Left = lipgloss.NewStyle().Padding(0, 1).Bold(true).Background(colors.Pink).Foreground(bg.C[0]).Render("Branches")
		case "c":
			m.Left = lipgloss.NewStyle().Padding(0, 1).Bold(true).Background(colors.Red).Foreground(bg.C[0]).Render("Commits")
		case "f":
			m.Left = lipgloss.NewStyle().Padding(0, 1).Bold(true).Background(colors.Yellow).Foreground(bg.C[0]).Render("Files")
		default:
			return m, nil
		}
	}

	return m, nil
}
