package logs

import (
	"fmt"
	"strings"

	"omzgit/git"
	"omzgit/lib/list"
	"omzgit/program/cokeline"
	"omzgit/program/logs/log"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list list.Model[log.Model]

	width  int
	height int
}

func InitialModel(width int, height int, title string) Model {
	initialList := list.InitialModel(getHeight(height), []log.Model{log.EmptyInitialModel(getWidth(width), "No logs found")}, 0, "No Logs Found")

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
		fmt.Sprint(m.list.ActiveRow+1, "/", len(m.list.Children)),

		m.list.Children[m.list.ActiveRow].Current,
	)
}

func (m Model) getCommitLogs() []log.Model {
	output, err := git.Exec("log", `--pretty=format:%h%n%D%n%s`)
	if err != nil {
		return []log.Model{}
	}

	head, _ := git.Exec("rev-parse", "--short", "HEAD")

	var logs []log.Model
	commitLogs := strings.Split(output, "\n")
	index := 0

	for len(logs) < m.list.NewSize() && index < len(commitLogs) {
		branchesStr := commitLogs[index+1]
		branchesStr = strings.TrimPrefix(branchesStr, "HEAD -> ")

		hash := commitLogs[index]
		branches := []string{}
		if len(branchesStr) > 0 {
			branches = strings.Split(branchesStr, ", ")
		}
		desc := commitLogs[index+2]

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
