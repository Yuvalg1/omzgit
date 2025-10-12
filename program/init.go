package program

import (
	"omzgit/messages"
	"omzgit/popups/alert"
	"omzgit/popups/async"
	"omzgit/popups/commit"
	"omzgit/popups/discard"
	"omzgit/popups/input"
	"omzgit/program/cokeline"
	"omzgit/program/popups"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	ActiveTab int
	cokeline  cokeline.Model
	Tabs      []tea.Model
	mode      string
	Popup     popups.Model[popups.InnerModel]

	Height int
	Width  int
}

type ExtendedModel struct {
	Tab   tea.Model
	Title string
}

func InitialModel(tabs []ExtendedModel, width int, height int) Model {
	initialPopups := popups.InitialModel[popups.InnerModel]("discard")

	initialInput := input.InitialModel(func(name string) {}, getWidth(width), getHeight(height), true)
	initialPopups.AddPopup("input", initialInput)

	initialDiscard := discard.InitialModel(func() tea.Cmd { return nil }, getWidth(width), getHeight(height))
	initialPopups.AddPopup("discard", initialDiscard)

	initialAlert := alert.InitialModel(getWidth(width), getHeight(height))
	initialPopups.AddPopup("alert", initialAlert)

	initialCommit := commit.InitialModel(getWidth(width), getHeight(height), "commit")
	initialPopups.AddPopup("commit", initialCommit)

	initialAsync := async.InitialModel(width, height, "fetching")
	initialPopups.AddPopup("async", initialAsync)

	return Model{
		ActiveTab: 0,
		cokeline:  cokeline.InitialModel(width, height, getCokes(tabs)),
		Tabs:      getTabs(tabs),
		Popup:     initialPopups,
		mode:      "",

		Width:  getWidth(width),
		Height: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	cmds := make([]tea.Cmd, 0, len(m.Tabs)+1)

	for _, element := range m.Tabs {
		cmds = append(cmds, element.Init())
	}

	cmds = append(cmds, m.Popup.Init())

	return tea.Batch(cmds...)
}

func (m Model) ModeCmd(mode string) tea.Cmd {
	return func() tea.Msg {
		return messages.ModeMsg{Mode: mode}
	}
}

func (m Model) PopupCmd(pType string, verb string, name string, callbackFn func() tea.Cmd) tea.Cmd {
	return func() tea.Msg {
		return messages.PopupMsg{
			Fn:   callbackFn,
			Type: pType,
			Name: name,
			Verb: verb,
		}
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
