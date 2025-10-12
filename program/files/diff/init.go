package diff

import (
	"os"

	"omzgit/git"
	"omzgit/program/files/row"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ogios/cropviewport"
)

type Model struct {
	content  string
	viewport cropviewport.CropViewportModel

	Staged bool
	path   string

	width  int
	height int
}

func InitialModel(row row.Model, width int, height int) Model {
	tWidth := getWidth(width)
	tHeight := getHeight(height)

	cropviewport := cropviewport.NewCropViewportModel().(*cropviewport.CropViewportModel)
	cropviewport.SetBlock(0, 0, tWidth, tHeight)

	m := Model{
		viewport: *cropviewport,
		Staged:   row.Staged,
		path:     row.Roller.Name,

		width:  tWidth,
		height: tHeight,
	}

	if row.Active {
		m.content = m.getDiffStaged()
	}

	return m
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

func (m Model) getDiffStaged() string {
	if m.Staged {
		output, _ := git.Exec("diff", "--staged", m.path)
		return string(output)
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
