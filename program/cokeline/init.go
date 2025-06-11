package cokeline

import (
	"fmt"
	"program/default/colors"
	"program/default/colors/status"

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
		Left:   lipgloss.NewStyle().Background(colors.Blue).Padding(0, 1).Render("Files"),
		Center: lipgloss.NewStyle().Background(status.Line[0]).Render(titles[0]),
		Right:  lipgloss.NewStyle().Background(status.Line[0]).Render(fmt.Sprint("0/", len(titles))),

		width: getWidth(width),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width
}
