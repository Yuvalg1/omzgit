package log

import (
	"strings"

	"omzgit/git"
	"omzgit/roller"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Active bool

	Hash     string
	branches []string
	Desc     roller.Model
	tip      string

	width int
}

func InitialModel(width int, hash string, branches []string, desc string) Model {
	return Model{
		Active: false,

		Hash:     hash,
		branches: branches,
		Desc:     roller.InitialModel(getWidth(width), desc),
		tip:      getBranchTip(branches, hash),

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

func getBranchTip(branches []string, hash string) string {
	if len(branches) == 0 {
		return ""
	}

	output, err := git.Exec("rev-parse", "--short", branches[len(branches)-1])

	if err == nil && hash == strings.TrimSpace(output) {
		return branches[0]
	}

	return ""
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width
}
