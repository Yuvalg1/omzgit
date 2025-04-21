package program

import (
	"program/consts"
	"program/program/popup"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Tabs      []tea.Model
	ActiveTab int
	Popup     popup.Model
	title     string
	pickMode  bool

	Height int
	Width  int
}

func InitialModel(tabs []tea.Model, width int, height int) Model {
	return Model{
		ActiveTab: consts.FILES - 1,
		Tabs:      tabs,
		Popup:     popup.InitialModel(func() {}, "", width, height),
		title:     "Files Changed",
		pickMode:  false,

		Width:  GetWidth(width),
		Height: GetHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	cmds := make([]tea.Cmd, 0, consts.PAGES+1)

	for _, element := range m.Tabs {
		cmds = append(cmds, element.Init())
	}

	return tea.Batch(cmds...)
}

func GetWidth(width int) int {
	return width
}

func GetHeight(height int) int {
	return height
}
