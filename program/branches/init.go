package branches

import (
	"os/exec"
	"program/program/branches/branch"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Title    string
	branches []branch.Model

	width  int
	height int
}

func InitialModel(width int, height int, title string) Model {
	return Model{
		branches: getBranches(getWidth(width), getHeight(height)),
		Title:    title,

		width:  getWidth(width),
		height: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width
}

func getHeight(height int) int {
	return height
}

func getBranches(width int, height int) []branch.Model {
	cmd := exec.Command("git", "branch")

	stdout, err := cmd.Output()
	if err != nil {
		return []branch.Model{}
	}

	branches := strings.Split(string(stdout), "\n")
	branches = branches[:len(branches)-1]

	var models []branch.Model
	for _, element := range branches {
		models = append(models, branch.InitialModel(width, height, element, "0 minutes ago", "0 | 0"))
	}

	return models
}
