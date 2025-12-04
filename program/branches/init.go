package branches

import (
	"fmt"
	"slices"
	"strings"

	"omzgit/git"
	"omzgit/lib/list"
	"omzgit/messages"
	"omzgit/program/branches/branch"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Title  string
	list   list.Model[branch.Model]
	remote bool

	width  int
	height int
}

func InitialModel(width int, height int, title string) Model {
	branches := getBranches(getWidth(width), getHeight(height), false)

	initialActive := slices.IndexFunc(branches, func(branch branch.Model) bool { return branch.Current })
	branches[initialActive].Active = true

	initialList := list.InitialModel(getHeight(height), branches, initialActive, "No Branches Found")

	initialList.SetCreateChild(func(name string) *branch.Model {
		created := branch.EmptyInitialModel(getWidth(width), getHeight(height), name, "")
		return &created
	})

	initialList.SetFilterFn(func(branch branch.Model, text string) bool {
		return strings.Contains(strings.ToLower(branch.Roller.Name), strings.ToLower(text))
	})

	m := Model{
		list:   initialList,
		remote: false,

		width:  getWidth(width),
		height: getHeight(height),
	}

	m.list.SetGetContentFn(func() []branch.Model {
		return getBranches(m.width, m.height, m.remote)
	})

	return m
}

func (m Model) Init() tea.Cmd {
	return m.list.Children[m.list.ActiveRow].Init()
}

func getWidth(width int) int {
	return width
}

func getHeight(height int) int {
	return height - 2
}

func (m Model) CokeCmd() tea.Cmd {
	return func() tea.Msg {
		return messages.CokeMsg{
			Center: m.list.Children[m.list.ActiveRow].Roller.Name,
			Right:  fmt.Sprint(m.list.ActiveRow+1, "/", len(m.list.Children)),

			Primary: m.list.Children[m.list.ActiveRow].Current,
		}
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

func getBranches(width int, height int, remote bool) []branch.Model {
	args := []string{"branch"}

	if remote {
		args = append(args, "--remote")
	}

	output, err := git.Exec(args...)
	if err != nil {
		return []branch.Model{}
	}

	branches := strings.Split(string(output), "\n")
	branches = branches[:len(branches)-1]

	var models []branch.Model
	for _, element := range branches {
		models = append(models, branch.InitialModel(width, height, element, getDefaultBranch()))
	}

	return models
}

func getDefaultBranch() string {
	output, err := git.Exec("symbolic-ref", "--short", "refs/remotes/origin/HEAD")
	if err != nil {
		return ""
	}

	return output[:len(output)-1]
}
