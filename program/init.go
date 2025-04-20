package program

import (
	"program/consts"
	"program/program/popup"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Tabs      map[string]tea.Model
	ActiveTab string
	Popup     popup.Model

	Height int
	Width  int
}

func InitialModel(tabs map[string]tea.Model, width int, height int) Model {
	return Model{
		ActiveTab: consts.FILES,
		Tabs:      tabs,
		Popup:     popup.InitialModel(func() {}, "", width, height),

		Height: height,
		Width:  width,
	}
}

func (m Model) Init() tea.Cmd {
	cmds := make([]tea.Cmd, 0, consts.PAGES+1)

	for _, element := range m.Tabs {
		cmds = append(cmds, element.Init())
	}

	return tea.Batch(cmds...)
}
