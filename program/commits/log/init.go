package log

import (
	"omzgit/roller"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Active bool

	Hash     string
	branches []string
	Desc     roller.Model

	width int
}

func InitialModel(width int, hash string, branches []string, desc string) Model {
	return Model{
		Active: false,

		Hash:     hash,
		branches: branches,
		Desc:     roller.InitialModel(getWidth(width), desc),

		width: getWidth(width),
	}
}

func EmptyInitialModel(width int, emptyMsg string) Model {
	return Model{
		Hash: "-------",
		Desc: roller.InitialModel(getWidth(width), emptyMsg),

		width: getWidth(width),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width
}
