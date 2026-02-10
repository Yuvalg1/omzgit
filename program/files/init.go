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

type Model struct {
	list  list.Model[row.Model]
	diffs []diff.Model

	height int
	width  int
}

func (m Model) CokeCmd() tea.Cmd {
	parts := m.getCurrentSplit()
	path := parts[len(parts)-1]

	return cokeline.Cmd(
		path,
		fmt.Sprintf(
			"%d/%d", m.list.ActiveRow+1, len(m.list.Children)),
		len(m.list.Children) > 0 && m.list.Children[m.list.ActiveRow].Staged,
	)
}

func InitialModel(width int, height int) Model {
	tWidth := getWidth(width)
	tHeight := getHeight(height)

	initialList := list.InitialModel(tHeight, []row.Model{}, 0, "No Files Found")

	initialList.SetFilterFn(filterFn)

	m := Model{
		list:  initialList,
		diffs: []diff.Model{},

		width:  tWidth,
		height: tHeight,
	}

	emptyRow := row.EmptyInitialModel("No Files Found", m.width)
	m.diffs = []diff.Model{diff.InitialModel(emptyRow, m.width, m.height)}
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

func (m Model) GetFilesChanged() []row.Model {
	output, err := git.Exec("status", "--short", "--untracked-files=all")
	if err != nil {
		return []row.Model{row.EmptyInitialModel("a files error has occured", m.width)}
	}

	fileLogs := strings.Split(string(output), "\n")
	fileLogs = fileLogs[:len(fileLogs)-1]

	if len(fileLogs) == 0 {
		return []row.Model{row.EmptyInitialModel("No Changes Made", m.width)}
	}

	var rows []row.Model
	index := 0

	for len(rows) < m.list.NewSize() && index < len(fileLogs) {
		row := row.InitialModel(fileLogs[index], m.width)

		if filterFn(row, m.list.TextInput.Value()) {
			rows = append(rows, row)
		}

		index++
	}

	return rows
}

func getDiffs(files []row.Model, width int, height int) []diff.Model {
	var diffs []diff.Model

	for _, element := range files {
		diffs = append(diffs, diff.InitialModel(element, width, height))
	}

	return diffs
}

func (m Model) getCurrentSplit() []string {
	return strings.Split(m.list.GetCurrent().Roller.Name, "/")
}

func filterFn(row row.Model, text string) bool {
	return strings.Contains(strings.ToLower(row.Roller.Name), strings.ToLower(text))
}
