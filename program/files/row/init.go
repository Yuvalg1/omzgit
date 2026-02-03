package row

import (
	"strings"

	"omzgit/roller"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Roller roller.Model
	status string

	Staged bool
	Active bool

	width int
}

func InitialModel(fileStr string, width int) Model {
	return Model{
		Active: false,
		Staged: getAdded(fileStr),
		Roller: roller.InitialModel(getWidth(width), getPath(fileStr)),
		status: getStatus(fileStr),

		width: getWidth(width),
	}
}

func EmptyInitialModel(fileStr string, width int) Model {
	return Model{
		Active: true,
		Staged: false,
		Roller: roller.InitialModel(getWidth(width), fileStr),
		status: " ",

		width: getWidth(width),
	}
}

func (m Model) Init() tea.Cmd {
	return m.Roller.Init()
}

func getWidth(width int) int {
	return width / 2
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
