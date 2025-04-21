package diff

import (
	"os"
	"os/exec"
	"program/program/files/diff/name"
	"program/program/files/diff/row"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Content []row.Model
	Name    name.Model
	staged  bool
	path    string

	Height int
	Width  int
}

func InitialModel(path string, staged bool, width int, height int) Model {
	return Model{
		Name:   name.InitialModel(path, GetWidth(width)),
		staged: staged,
		path:   path,

		Width:  GetWidth(width),
		Height: GetHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func GetWidth(width int) int {
	return width / 2
}

func GetHeight(height int) int {
	return height
}

func (m Model) GetContent() []row.Model {
	texts, isDesc := getTexts(m.path, m.staged)

	if len(texts) > m.Height {
		texts = texts[:m.Height]
	}

	var rows []row.Model

	for _, element := range texts {
		if strings.HasPrefix(element, "@@") {
			isDesc = false
		}

		rows = append(rows, row.InitialModel(element, isDesc, m.Width))
	}

	return rows
}

func getTexts(path string, staged bool) ([]string, bool) {
	if staged {
		cmd := exec.Command("git", "diff", "--staged", path)

		stdout, err := cmd.Output()
		if err != nil {
			return []string{"a git diff error has occured"}, true
		}

		texts := strings.Split(string(stdout), "\n")

		return texts[:len(texts)-1], true
	}

	file, err := os.Stat(path)
	if err != nil {
		return []string{"an os error has occured"}, false
	}

	if file.Size() > 100*1000 { // bigger than 100kb
		return []string{"file size is too big to render."}, false
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return []string{"a file reading error has occured"}, false
	}

	return strings.Split(string(data), "\n"), false
}
