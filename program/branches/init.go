package branches

import (
	"fmt"
	"slices"
	"strings"

	"omzgit/default/colors"
	"omzgit/default/colors/bg"
	"omzgit/default/colors/gray"
	"omzgit/git"
	"omzgit/lib/list"
	"omzgit/messages"
	"omzgit/program/branches/branch"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

func (m Model) CokeCmd() tea.Cmd {
	return func() tea.Msg {
		return messages.CokeMsg{
			Center: lipgloss.NewStyle().
				Background(bg.C[2]).
				Foreground(m.getCurrentBranchColor()).
				Padding(0, 1).
				Render(m.list.Children[m.list.ActiveRow].Name),
			Right: lipgloss.NewStyle().
				Background(gray.C[1]).
				Padding(0, 1).
				Render(fmt.Sprint(m.list.ActiveRow+1, "/", len(m.list.Children))),
		}
	}
}

func (m Model) getCurrentBranchColor() lipgloss.Color {
	if m.list.Children[m.list.ActiveRow].Current {
		return lipgloss.Color(colors.Yellow)
	}
	return lipgloss.Color(colors.Blue)
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
	output, err := git.Exec("branch")
	if err != nil {
		return []branch.Model{}
	}

	branches := strings.Split(string(output), "\n")
	branches = branches[:len(branches)-1]

	var models []branch.Model
	for _, element := range branches {
		models = append(models, branch.InitialModel(width, height, element, getDefaultBranch(), false))
	}

	return models
}

func getDefaultBranch() string {
	output, err := git.Exec("symbolic-ref", "--short", "refs/remotes/origin/HEAD")
	if err != nil {
		return ""
	}

	return string(output)[:len(string(output))-1]
}
