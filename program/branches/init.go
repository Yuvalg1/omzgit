package branches

import (
	"fmt"
	"strings"

	"omzgit/git"
	"omzgit/lib/list"
	"omzgit/program/branches/branch"
	"omzgit/program/cokeline"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Title  string
	list   list.Model[branch.Model]
	remote bool
	total  int

	width  int
	height int
}

func InitialModel(width int, height int, title string) Model {
	initialList := list.InitialModel(getHeight(height), []branch.Model{}, 0, "No Branches Found")

	m := Model{
		list:   initialList,
		remote: false,

		width:  getWidth(width),
		height: getHeight(height),
	}

	m.list.Children = []branch.Model{branch.EmptyInitialModel(m.width, getHeight(height), "No Branches Found", "")}
	m.list.SetCreateChild(func(name string) *branch.Model {
		created := branch.EmptyInitialModel(m.width, getHeight(height), name, "")
		return &created
	})

	m.list.SetFilterFn(filterFn)

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
	return cokeline.Cmd(
		m.list.Children[m.list.ActiveRow].Roller.Name,
		fmt.Sprintf("%d/%d", m.list.ActiveRow+1, m.total),
		m.list.Children[m.list.ActiveRow].Current,
	)
}

func getBranches(m snapshot) []branch.Model {
	args := []string{"branch"}

	if m.remote {
		args = append(args, "--remote")
	}

	output, err := git.Exec(args...)
	if err != nil {
		return []branch.Model{}
	}

	branches := strings.Split(output, "\n")
	branches = branches[:len(branches)-1]

	index := 0

	var models []branch.Model
	for len(models) < m.listNewSize && index < len(branches) {
		branch := branch.InitialModel(m.width, branches[index], getDefaultBranch())

		if filterFn(branch, m.listTextInputValue) {
			models = append(models, branch)
		}

		index++
	}

	if len(models) == 0 {
		models = append(models, branch.EmptyInitialModel(m.width, getHeight(m.height), "No Branches Found", ""))
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

func filterFn(branch branch.Model, text string) bool {
	return strings.Contains(strings.ToLower(branch.Roller.Name), strings.ToLower(text))
}

type snapshot struct {
	remote bool
	total  int

	listNewSize        int
	listTextInputValue string

	width  int
	height int
}

func (m Model) getSnapshot() snapshot {
	return snapshot{
		remote: m.remote,
		total:  m.total,

		listNewSize:        m.list.NewSize(),
		listTextInputValue: m.list.TextInput.Value(),

		width:  m.width,
		height: m.height,
	}
}
