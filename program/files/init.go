package files

import (
	"fmt"
	"strings"

	"omzgit/git"
	"omzgit/lib/list"
	"omzgit/messages/tick"
	"omzgit/program/cokeline"
	"omzgit/program/files/diff"
	"omzgit/program/files/row"

	tea "github.com/charmbracelet/bubbletea"
)

var CUTOFF = 50

type Model struct {
	list  list.Model[row.Model]
	diff  diff.Model
	total int

	height int
	width  int
}

func (m Model) CokeCmd() tea.Cmd {
	parts := m.getCurrentSplit()
	path := parts[len(parts)-1]

	return cokeline.Cmd(
		path,
		fmt.Sprintf("%d/%d", m.list.ActiveRow+1, m.total),
		len(m.list.Children) > 0 && m.list.Children[m.list.ActiveRow].Staged,
	)
}

func InitialModel(width int, height int) Model {
	initialList := list.InitialModel(getHeight(height), []row.Model{}, 0, "No Files Found")

	m := Model{
		list: initialList,
		diff: diff.Model{},

		width:  getWidth(width),
		height: getHeight(height),
	}

	emptyRow := row.EmptyInitialModel("No Files Found", m.width)
	m.list.Children = []row.Model{emptyRow}
	m.diff = diff.InitialModel(emptyRow, m.width, m.height)
	m.list.SetCreateChild(func(name string) *row.Model {
		created := row.EmptyInitialModel("No Files Found", m.width)
		return &created
	})

	return m
}

func (m Model) Init() tea.Cmd {
	cmds := []tea.Cmd{tick.Cmd(0), func() tea.Msg {
		return tea.KeyMsg{Type: tea.KeyEsc, Runes: []rune{'\x1b'}}
	}}

	if len(m.list.Children) > 0 {
		cmds = append(cmds, m.list.Children[m.list.ActiveRow].Init(), m.CokeCmd())
	}

	return tea.Batch(cmds...)
}

func getWidth(width int) int {
	return width
}

func getHeight(height int) int {
	return height - 2
}

func (m Model) getFilesAxis() (int, int) {
	if m.width > CUTOFF {
		return m.width / 2, m.height
	}
	return m.width, m.height / 2
}

func (m Model) getDiffAxis() (int, int) {
	if m.width > CUTOFF {
		return m.width/2 - (m.width+1)%2, m.height
	}
	return m.width, m.height/2 - (m.height+1)%2
}

func GetFilesChanged(m snapshot) []row.Model {
	output, err := git.Exec("status", "--porcelain", "--untracked-files=all")
	if err != nil {
		return []row.Model{row.EmptyInitialModel("a files error has occured", m.width)}
	}

	fileLogs := strings.Split(output, "\n")
	fileLogs = fileLogs[:len(fileLogs)-1]

	if len(fileLogs) == 0 {
		return []row.Model{row.EmptyInitialModel("No Changes Made", m.width)}
	}

	rows := []row.Model{}
	modified := []row.Model{}
	index := 0

	for len(rows) < m.listNewSize && index < len(fileLogs) {
		path := getPath(fileLogs[index])

		if filterFn(path, m.listTextInputValue) {
			rows = append(rows, row.InitialModel(fileLogs[index], m.width))
		}

		if filterFn(path, m.listTextInputValue) && fileLogs[index][0] != ' ' && fileLogs[index][1] == 'M' {
			modified = append(modified, row.InitialModel(" "+fileLogs[index][1:], m.width))
		}
		index++
	}

	rows = append(rows, modified...)
	if len(rows) == 0 {
		rows = append(rows, row.EmptyInitialModel("No Changes Made", m.width))
	}

	return rows
}

func (m Model) getCurrentSplit() []string {
	return strings.Split(m.list.GetCurrent().Roller.Name, "/")
}

func getPath(fileStr string) string {
	return strings.Split(fileStr[2:], " ")[1]
}

func filterFn(path string, text string) bool {
	return strings.Contains(path, strings.ToLower(text))
}
