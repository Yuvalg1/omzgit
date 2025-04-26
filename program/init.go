package program

import (
	"program/consts"
	"program/program/cokeline"
	"program/program/popup"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	ActiveTab int
	cokeline  cokeline.Model
	Tabs      []tea.Model
	pickMode  bool
	Popup     popup.Model

	Height int
	Width  int
}

type ExtendedModel struct {
	Tab   tea.Model
	Title string
}

func InitialModel(tabs []ExtendedModel, width int, height int) Model {
	return Model{
		ActiveTab: consts.FILES - 1,
		cokeline:  cokeline.InitialModel(width, height, getCokes(tabs)),
		Tabs:      getTabs(tabs),
		Popup:     popup.InitialModel(func() {}, "", getWidth(width), getHeight(height)),
		pickMode:  false,

		Width:  getWidth(width),
		Height: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	cmds := make([]tea.Cmd, 0, consts.PAGES+1)

	for _, element := range m.Tabs {
		cmds = append(cmds, element.Init())
	}

	return tea.Batch(cmds...)
}

func getWidth(width int) int {
	return width
}

func getHeight(height int) int {
	return height
}

func getTabs(extended []ExtendedModel) []tea.Model {
	var tabs []tea.Model

	for _, element := range extended {
		tabs = append(tabs, element.Tab)
	}

	return tabs
}

func getCokes(tabs []ExtendedModel) []string {
	var cokes []string

	for _, element := range tabs {
		cokes = append(cokes, element.Title)
	}

	return cokes
}
