package program

import (
	"program/consts"
	"program/messages"
	"program/popups/discard"
	"program/popups/input"
	"program/program/cokeline"
	"program/program/popup"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	ActiveTab int
	cokeline  cokeline.Model
	Tabs      []tea.Model
	mode      string
	Popup     popup.Model[popup.InnerModel]

	Height int
	Width  int
}

type ExtendedModel struct {
	Tab   tea.Model
	Title string
}

func InitialModel(tabs []ExtendedModel, width int, height int) Model {
	initialPopups := popup.InitialModel[popup.InnerModel]("discard")

	initialInput := input.InitialModel(func(name string) {}, "", getWidth(width), getHeight(height))
	initialPopups.AddPopup("input", initialInput)

	initialDiscard := discard.InitialModel(func() {}, "", getWidth(width), getHeight(height))
	initialPopups.AddPopup("discard", initialDiscard)

	return Model{
		ActiveTab: consts.FILES - 1,
		cokeline:  cokeline.InitialModel(width, height, getCokes(tabs)),
		Tabs:      getTabs(tabs),
		Popup:     initialPopups,
		mode:      "",

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

func (m Model) ModeCmd(mode string) tea.Cmd {
	return func() tea.Msg {
		return messages.ModeMsg{Mode: mode}
	}
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
