package diff

import (
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ogios/cropviewport"
)

type Model struct {
	Content  string
	viewport cropviewport.CropViewportModel

	staged bool
	path   string

	width  int
	height int
}

func InitialModel(path string, staged bool, width int, height int) Model {
	tWidth := getWidth(width)
	tHeight := getHeight(height)
	cropviewport := cropviewport.NewCropViewportModel().(*cropviewport.CropViewportModel)
	cropviewport.SetBlock(0, 0, tWidth, tHeight)
	return Model{
		viewport: *cropviewport,
		staged:   staged,
		path:     path,

		width:  tWidth,
		height: tHeight,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width/2 + width%2
}

func getHeight(height int) int {
	return height - 2
}

func (m *Model) SetWidth(width int) {
	block := m.viewport.Block
	m.viewport.SetBlock(block[0], block[1], getWidth(width), block[3])
	m.width = getWidth(width)
}

func (m *Model) SetHeight(height int) {
	block := m.viewport.Block
	m.viewport.SetBlock(block[0], block[1], block[2], getHeight(height))
	m.height = getHeight(height)
}

func (m Model) GetContent() string {
	return m.getDiffStaged()
}

func (m Model) getDiffStaged() string {
	if m.staged {
		cmd := exec.Command("git", "diff", "--staged", m.path)

		stdout, err := cmd.Output()
		if err != nil {
			return "Staged File has been deleted."
		}

		return string(stdout)
	}

	file, err := os.Stat(m.path)
	if err != nil {
		return "Unstaged File has been deleted."
	}

	if file.Size() > 100*1000 { // bigger than 100kb
		return "File size is too big to render."
	}

	data, err := os.ReadFile(m.path)
	if err != nil {
		return "a file reading error has occured"
	}

	return string(data)
}
