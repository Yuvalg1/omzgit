package commits

import (
	"fmt"
	"strings"

	"omzgit/git"
	"omzgit/lib/list"
	"omzgit/messages"
	"omzgit/program/commits/log"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list list.Model[log.Model]

	width  int
	height int
}

func InitialModel(width int, height int, title string) Model {
	logs := getCommitLogs(getWidth(width))
	initialList := list.InitialModel(getHeight(height), logs, 0, "No Commits Found")

	initialList.SetCreateChild(func(name string) *log.Model {
		created := log.EmptyInitialModel(getWidth(width), name)
		return &created
	})
	initialList.SetFilterFn(func(row log.Model, text string) bool {
		return strings.Contains(strings.ToLower(row.Hash), strings.ToLower(text)) ||
			strings.Contains(strings.ToLower(row.Desc.Name), strings.ToLower(text))
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

func (m Model) CokeCmd() tea.Cmd {
	return func() tea.Msg {
		return messages.CokeMsg{
			Center: m.list.Children[m.list.ActiveRow].Hash,
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

func getCommitLogs(width int) []log.Model {
	output, err := git.Exec("log", `--pretty=format:%h%n%D%n%s`)
	if err != nil {
		return []log.Model{}
	}

	head, _ := git.Exec("rev-parse", "--short", "HEAD")

	var logs []log.Model
	commits := strings.Split(output, "\n")

	for i := 0; i < len(commits); i += 3 {
		branchesStr := commits[i+1]
		branchesStr = strings.TrimPrefix(branchesStr, "HEAD -> ")

		hash := commits[i]
		branches := []string{}
		if len(branchesStr) > 0 {
			branches = strings.Split(branchesStr, ", ")
		}
		desc := commits[i+2]

		logs = append(logs, log.InitialModel(width, hash, branches, desc, strings.TrimSpace(head)))
	}

	return logs
}

func getWidth(width int) int {
	return width
}

func getHeight(height int) int {
	return height - 2
}
