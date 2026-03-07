package diff

import (
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
		m.content = m.getDiff()
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

func (m Model) getDiff() string {
	if m.Staged {
		output, _ := git.Exec("diff", "--staged", m.path)
		return output
	}

	output, _ := git.Exec("diff", m.path)
	return output
}
