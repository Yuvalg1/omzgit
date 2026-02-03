package files

import (
	"fmt"
	"strings"
	"time"

	"omzgit/git"
	"omzgit/lib/list"
	"omzgit/messages"
	"omzgit/messages/tick"
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

func (m Model) TickCmd() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(10 * time.Second)
		return tick.Msg{RollOffset: m.list.Children[m.list.ActiveRow].Roller.Offset}
	}
}

func (m Model) CokeCmd() tea.Cmd {
	return func() tea.Msg {
		parts := m.getCurrentSplit()
		path := parts[len(parts)-1]

		return messages.CokeMsg{
			Center: path,
			Right: fmt.Sprintf(
				"%d/%d", m.list.ActiveRow+1, len(m.list.Children)),

			Primary: len(m.list.Children) > 0 && m.list.Children[m.list.ActiveRow].Staged,
		}
	}
}

func InitialModel(width int, height int) Model {
	tWidth := getWidth(width)
	tHeight := getHeight(height)

	emptyRow := row.EmptyInitialModel("No Files Found", getWidth(width))
	initialList := list.InitialModel(tHeight, []row.Model{emptyRow}, 0, "No Files Found")

	initialList.SetCreateChild(func(name string) *row.Model {
		created := row.EmptyInitialModel("No Files Found", getWidth(width))
		return &created
	})

	initialList.SetFilterFn(func(row row.Model, text string) bool {
		return strings.Contains(strings.ToLower(row.Roller.Name), strings.ToLower(text))
	})

	m := Model{
		list:  initialList,
		diffs: []diff.Model{diff.InitialModel(emptyRow, tWidth, tHeight)},

		width:  tWidth,
		height: tHeight,
	}

	m.list.SetGetContentFn(func() []row.Model {
		return GetFilesChanged(m.width)
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

func GetFilesChanged(width int) []row.Model {
	output, err := git.Exec("status", "--short", "--untracked-files=all")
	if err != nil {
		return []row.Model{row.EmptyInitialModel("a files error has occured", width)}
	}

	fileLogs := strings.Split(string(output), "\n")
	fileLogs = fileLogs[:len(fileLogs)-1]

	if len(fileLogs) == 0 {
		return []row.Model{row.EmptyInitialModel("No Changes Made", width)}
	}

	var rows []row.Model

	for _, element := range fileLogs {
		rows = append(rows, row.InitialModel(element, width))
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
