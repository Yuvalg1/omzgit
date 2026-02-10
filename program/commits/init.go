package commits

import (
	"fmt"
	"strings"

	"omzgit/git"
	"omzgit/lib/list"
	"omzgit/program/cokeline"
	"omzgit/program/commits/log"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list  list.Model[log.Model]
	total int

	width  int
	height int
}

func InitialModel(width int, height int, title string) Model {
	initialList := list.InitialModel(getHeight(height), []log.Model{log.EmptyInitialModel(getWidth(width), "No commits found")}, 0, "No Commits Found")

	initialList.SetCreateChild(func(name string) *log.Model {
		created := log.EmptyInitialModel(getWidth(width), name)
		return &created
	})
	initialList.SetFilterFn(filterFn)

	m := Model{
		list: initialList,

		width:  getWidth(width),
		height: getHeight(height),
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) CokeCmd() tea.Cmd {
	return cokeline.Cmd(
		m.list.Children[m.list.ActiveRow].Hash,
		fmt.Sprintf("%d/%d", m.list.ActiveRow+1, m.total),
		m.list.Children[m.list.ActiveRow].Current,
	)
}

func (m *Model) getCommitLogs() []log.Model {
	output, err := git.Exec("log", `--pretty=format:%h%n%D%n%s`)
	if err != nil {
		return []log.Model{}
	}

	head, _ := git.Exec("rev-parse", "--short", "HEAD")

	var logs []log.Model
	commits := strings.Split(output, "\n")

	m.total = len(commits) / 3

	index := 0

	for len(logs) < m.list.NewSize() && index < len(commits) {
		branchesStr := commits[index+1]
		branchesStr = strings.TrimPrefix(branchesStr, "HEAD -> ")

		hash := commits[index]
		branches := []string{}
		if len(branchesStr) > 0 {
			branches = strings.Split(branchesStr, ", ")
		}
		desc := commits[index+2]

		log := log.InitialModel(m.width, hash, branches, desc, strings.TrimSpace(head))

		if filterFn(log, m.list.TextInput.Value()) {
			logs = append(logs, log)
		}
		index += 3
	}

	return logs
}

func filterFn(row log.Model, text string) bool {
	return strings.Contains(strings.ToLower(row.Hash), strings.ToLower(text)) ||
		strings.Contains(strings.ToLower(row.Desc.Name), strings.ToLower(text))
}

func getWidth(width int) int {
	return width
}

func getHeight(height int) int {
	return height - 2
}
