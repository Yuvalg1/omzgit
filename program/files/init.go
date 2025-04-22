package files

import (
	"os/exec"
	"program/consts"
	"program/messages"
	"program/program/files/diff"
	"program/program/files/row"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	files     []row.Model
	Diffs     []diff.Model
	ActiveRow int
	title     string

	Height int
	Width  int
}

func (m Model) PopupCmd(path string, fn func()) tea.Cmd {
	return func() tea.Msg {
		return messages.PopupMsg{
			Fn:   fn,
			Name: path,
		}
	}
}

func (m Model) TickCmd() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(consts.REFRESH_INTERVAL)
		return messages.TickMsg{}
	}
}

func (m Model) TitleCmd(title string) tea.Cmd {
	return func() tea.Msg {
		return messages.TitleMsg{Title: title}
	}
}

func InitialModel(width int, height int) Model {
	tWidth := GetWidth(width)
	tHeight := GetHeight(height)

	files := GetFilesChanged(tWidth)

	files[0].Active = true

	return Model{
		files:     files,
		Diffs:     getDiffs(files, tWidth, tHeight),
		ActiveRow: 0,
		title:     "Files",

		Width:  tWidth,
		Height: tHeight,
	}
}

func (m Model) Init() tea.Cmd {
	cmds := []tea.Cmd{m.TickCmd()}

	if len(m.files) > 0 {
		cmds = append(cmds, m.files[0].Init())
	}

	return tea.Batch(cmds...)
}

func GetWidth(width int) int {
	return width - 2
}

func GetHeight(height int) int {
	return height - 2
}

func GetFilesChanged(width int) []row.Model {
	cmd := exec.Command("git", "status", "--short", "--untracked-files=all")

	stdout, err := cmd.Output()
	if err != nil {
		return []row.Model{row.InitialModel("a files error has occured", GetWidth(width), true)}
	}

	fileLogs := strings.Split(string(stdout), "\n")
	fileLogs = fileLogs[:len(fileLogs)-1]

	if len(fileLogs) == 0 {
		return []row.Model{row.InitialModel("No Changes Made", width, true)}
	}

	var rows []row.Model

	for _, element := range fileLogs {
		rows = append(rows, row.InitialModel(element, width, false))
	}

	return rows
}

func getDiffs(files []row.Model, width int, height int) []diff.Model {
	var diffs []diff.Model

	for _, element := range files {
		diffs = append(diffs, diff.InitialModel(element.Path, element.Staged, width, height))
	}

	if len(diffs) > 0 {
		diffs[0].Content = diffs[0].GetContent()
	}

	return diffs
}
