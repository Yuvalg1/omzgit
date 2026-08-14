package picker

import (
	"omzgit/program/popups"
	"omzgit/suspend"

	tea "github.com/charmbracelet/bubbletea"
)

type Pick struct {
	Desc     string
	Callback func() tea.Cmd
}

type Model struct {
	name    string
	title   string
	visible bool
	options map[string]Pick

	width  int
	height int
}

func InitialModel(width int, height int) Model {
	return Model{
		visible: false,

		width:  getWidth(width),
		height: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getHeight(height int) int {
	return 4
}

func getWidth(width int) int {
	return min(34, width-2-width%2)
}

func (m Model) GetVisible() bool {
	return m.visible
}

func GetPick(command ...string) Pick {
	return Pick{
		Desc: command[1], Callback: func() tea.Cmd {
			return popups.Cmd("async", "", command[1], func() tea.Cmd {
				return suspend.Cmd(command...)
			})
		},
	}
}
