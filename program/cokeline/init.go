package cokeline

import (
	"omzgit/default/colors"
	"omzgit/default/colors/bg"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Left   string
	Center string
	Right  string

	width int
}

func InitialModel(width int, height int, titles []string) Model {
	return Model{
		Left: lipgloss.NewStyle().Background(colors.Yellow).Bold(true).Foreground(bg.C[0]).Padding(0, 1).Render("Files"),

		width: getWidth(width),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width
}
