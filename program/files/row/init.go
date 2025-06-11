package row

import (
	"program/messages"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Path   string
	status string

	Staged bool
	Active bool

	width  int
	height int
}

func InitialModel(fileStr string, width int, empty bool) Model {
	if empty {
		return Model{
			Active: true,
			Staged: false,
			Path:   fileStr,
			status: " ",

			width:  getWidth(width),
			height: getHeight(1),
		}
	}

	return Model{
		Active: false,
		Staged: getAdded(fileStr),
		Path:   getPath(fileStr),
		status: getStatus(fileStr),

		width:  getWidth(width),
		height: getHeight(1),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width/2 - 1
}

func getHeight(height int) int {
	return 1
}

func (m Model) PopupCmd(pType string, verb string, path string, fn func() bool) tea.Cmd {
	return func() tea.Msg {
		parts := strings.Split(path, "/")
		return messages.PopupMsg{
			Fn:   fn,
			Name: "'" + parts[len(parts)-1] + "'",
			Type: pType,
			Verb: verb,
		}
	}
}

func getAdded(fileStr string) bool {
	return fileStr[0] == 'A' || fileStr[0] == 'M' || fileStr[0] == 'D' || fileStr[0] == 'R'
}

func getPath(fileStr string) string {
	return strings.Split(fileStr[2:], " ")[1]
}

func getStatus(fileStr string) string {
	firstTwoChars := fileStr[:2]

	if strings.Contains(firstTwoChars, "A") {
		return "A"
	}

	if strings.Contains(firstTwoChars, "D") {
		return "D"
	}

	if strings.Contains(firstTwoChars, "M") {
		return "M"
	}

	if strings.Contains(firstTwoChars, "?") {
		return "U"
	}

	return "?"
}
