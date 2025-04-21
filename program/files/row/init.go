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
			Active: false,
			Staged: false,
			Path:   fileStr,
			status: " ",

			width:  GetWidth(width),
			height: GetHeight(1),
		}
	}

	return Model{
		Active: false,
		Staged: getAdded(fileStr),
		Path:   getPath(fileStr),
		status: getStatus(fileStr),

		width:  GetWidth(width),
		height: GetHeight(1),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func GetWidth(width int) int {
	return width / 2
}

func GetHeight(height int) int {
	return 1
}

func (m Model) PopupCmd(path string, fn func()) tea.Cmd {
	return func() tea.Msg {
		parts := strings.Split(path, "/")
		return messages.PopupMsg{
			Fn:   fn,
			Name: "'" + parts[len(parts)-1] + "'",
		}
	}
}

func getAdded(fileStr string) bool {
	return fileStr[0] == 'A' || fileStr[0] == 'M'
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
