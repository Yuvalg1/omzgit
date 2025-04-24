package diff

import (
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Content  string
	viewport textarea.Model

	staged bool
	path   string
}

func InitialModel(path string, staged bool, width int, height int) Model {
	textarea := textarea.New()
	textarea.SetWidth(getWidth(width))
	textarea.SetHeight(height)

	return Model{
		viewport: textarea,
		staged:   staged,
		path:     path,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width / 2
}

func getHeight(height int) int {
	return height
}

func (m *Model) SetWidth(width int) {
	m.viewport.SetWidth(getWidth(width))
}

func (m Model) SetHeight(height int) {
	m.viewport.SetHeight(getHeight(height))
}

func (m Model) GetContent() string {
	return m.getDiffStaged()
}

func (m Model) getDiffStaged() string {
	if m.staged {
		cmd := exec.Command("git", "diff", "--staged", m.path)

		stdout, err := cmd.Output()
		if err != nil {
			return "a git diff error has occured"
		}

		return string(stdout)
	}

	file, err := os.Stat(m.path)
	if err != nil {
		return "an os error has occured"
	}

	if file.Size() > 100*1000 { // bigger than 100kb
		return "file size is too big to render."
	}

	data, err := os.ReadFile(m.path)
	if err != nil {
		return "a file reading error has occured"
	}

	return string(data)
}
