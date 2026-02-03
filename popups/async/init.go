package async

import (
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	callbackFn func() tea.Cmd
	spinner    spinner.Model
	title      string
	visible    bool

	width  int
	height int
}

func InitialModel(width int, height int, title string) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style.Background(bg.C[0])

	return Model{
		callbackFn: func() tea.Cmd { return nil },
		spinner:    s,
		title:      title,
		visible:    true,

		width:  getWidth(width),
		height: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) GetVisible() bool {
	return m.visible
}

func getWidth(width int) int {
	return 18
}

func getHeight(height int) int {
	return 3
}
