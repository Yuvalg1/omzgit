package option

import (
	"omzgit/env"
	"omzgit/roller"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Active bool
	Msg    string
	Roller roller.Model

	width int
}

func EmptyInitialModel(width int, height int) Model {
	return Model{
		Active: true,
		Roller: roller.InitialModel(getWidth(width), "no options availible"),
		Msg:    "Msg",

		width: getWidth(width),
	}
}

func InitialModel(width int, option env.Option) Model {
	msg := option.Msg

	if option.AltMsg != "" {
		msg += "/" + option.AltMsg
	}

	return Model{
		Active: false,
		Roller: roller.InitialModel(getWidth(width), option.Description),
		Msg:    msg,

		width: getWidth(width),
	}
}

func (m Model) Init() tea.Cmd {
	return m.Roller.Init()
}

func getWidth(width int) int {
	return width - 2
}
