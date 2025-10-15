package log

import (
	"strings"

	"omzgit/git"
	"omzgit/messages"
	"omzgit/roller"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Active  bool
	Current bool

	Hash     string
	branches []string
	Desc     roller.Model
	tip      string

	width int
}

func InitialModel(width int, hash string, branches []string, desc string, head string) Model {
	return Model{
		Active:  false,
		Current: hash == head,

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

func (m Model) PopupCmd(pType string, placeholder string, title string, fn any) tea.Cmd {
	return func() tea.Msg {
		return messages.PopupMsg{
			Fn:   fn,
			Name: title,
			Type: pType,
			Verb: placeholder,
		}
	}
}

func (m Model) RefreshCmd() tea.Cmd {
	return func() tea.Msg {
		return messages.RefreshMsg{}
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
