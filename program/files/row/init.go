package row

import (
	"slices"
	"strings"

	"omzgit/roller"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	staged    = []byte{'M', 'A', 'D', 'R', 'C', 'U'}
	conflicts = []string{"DD", "AU", "UD", "UA", "DU", "AA", "UU"}
)

type Model struct {
	Roller roller.Model
	status string

	Staged   bool
	Active   bool
	Conflict bool

	width int
}

func InitialModel(fileStr string, width int) Model {
	return Model{
		Staged:   getStaged(fileStr),
		Active:   false,
		Conflict: slices.Contains(conflicts, fileStr[:2]),
		Roller:   roller.InitialModel(getWidth(width), getPath(fileStr)),
		status:   getStatus(fileStr),

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

func getStaged(fileStr string) bool {
	return slices.Contains(staged, fileStr[0])
}

func getPath(fileStr string) string {
	return strings.Split(fileStr[2:], " ")[1]
}

func getStatus(fileStr string) string {
	if fileStr[:2] == "??" {
		return "U"
	}

	staged := getStaged(fileStr)

	if staged {
		return string(fileStr[0])
	} else {
		return string(fileStr[1])
	}
}
