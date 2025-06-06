package branches

import (
	"os/exec"
	"program/lib/list"
	"program/messages"
	"program/program/branches/branch"
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Title string
	list  list.Model[branch.Model]

	width  int
	height int
}

func InitialModel(width int, height int, title string) Model {
	branches := getBranches(getWidth(width), getHeight(height))

	initialActive := slices.IndexFunc(branches, func(branch branch.Model) bool { return branch.Current })
	branches[initialActive].Active = true

	initialList := list.InitialModel(getWidth(width), getHeight(height), branches, initialActive, "No Branches Found")
	initialList.SetCreateChild(func(name string) *branch.Model {
		created := branch.InitialModel(getWidth(width), getHeight(height), getDefaultBranch(), "", true)
		return &created
	})
	initialList.SetFilterFn(func(branch branch.Model, text string) bool {
		return strings.Contains(branch.Name, text)
	})
	return Model{
		list: initialList,

		width:  getWidth(width),
		height: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width
}

func getHeight(height int) int {
	return height - 2
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

func getBranches(width int, height int) []branch.Model {
	cmd := exec.Command("git", "branch")

	stdout, err := cmd.Output()
	if err != nil {
		return []branch.Model{}
	}

	branches := strings.Split(string(stdout), "\n")
	branches = branches[:len(branches)-1]

	var models []branch.Model
	for _, element := range branches {
		models = append(models, branch.InitialModel(width, height, element, getDefaultBranch(), false))
	}

	return models
}

func getDefaultBranch() string {
	cmd := exec.Command("git", "symbolic-ref", "--short", "refs/remotes/origin/HEAD")

	stdout, err := cmd.Output()
	if err != nil {
		return ""
	}

	return string(stdout)[:len(string(stdout))-1]
}
