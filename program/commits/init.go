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
	initialList := list.InitialModel(getHeight(height), []log.Model{}, 0, "No Commits Found")

	m := Model{
		list: initialList,

		width:  getWidth(width),
		height: getHeight(height),
	}

	m.list.Children = []log.Model{log.EmptyInitialModel(m.width, "No commits found")}
	m.list.SetCreateChild(func(name string) *log.Model {
		created := log.EmptyInitialModel(m.width, name)
		return &created
	})
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

func getCommitLogs(m snapshot) []log.Model {
	output, err := git.Exec("log", `--pretty=format:%h%n%D%n%s`)
	if err != nil {
		return []log.Model{}
	}

	head, _ := git.Exec("rev-parse", "--short", "HEAD")

	var logs []log.Model
	commits := strings.Split(output, "\n")

	index := 0

	for len(logs) < m.listNewSize && index < len(commits) {
		branchesStr := commits[index+1]
		branchesStr = strings.TrimPrefix(branchesStr, "HEAD -> ")

		hash := commits[index]
		branches := []string{}
		if len(branchesStr) > 0 {
			branches = strings.Split(branchesStr, ", ")
		}
		desc := commits[index+2]

		if filterFn(hash, desc, m.listTextInputValue) {
			logs = append(logs, log.InitialModel(m.width, hash, branches, desc, strings.TrimSpace(head)))
		}
		index += 3
	}

	if len(logs) == 0 {
		logs = append(logs, log.EmptyInitialModel(m.width, "No Changes Made"))
	}

	return logs
}

func filterFn(hash string, desc string, text string) bool {
	return strings.Contains(strings.ToLower(hash), strings.ToLower(text)) ||
		strings.Contains(strings.ToLower(desc), strings.ToLower(text))
}

func getWidth(width int) int {
	return width
}

func getHeight(height int) int {
	return height - 2
}

type snapshot struct {
	total int

	listNewSize        int
	listTextInputValue string

	width  int
	height int
}

func (m Model) getSnapshot() snapshot {
	return snapshot{
		total: m.total,

		listNewSize:        m.list.NewSize(),
		listTextInputValue: m.list.TextInput.Value(),

		width:  m.width,
		height: m.height,
	}
}
