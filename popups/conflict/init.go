package conflict

import (
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	path     string
	visible  bool
	conflict int

	ours   string
	theirs string

	width  int
	height int
}

func InitialModel(width int, height int) Model {
	return Model{
		visible: false,

		width:  width,
		height: height,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) GetVisible() bool {
	return m.visible
}

func (m *Model) getContent() { // inConflict is needed for malformed files
	file, err := os.Stat(m.path)
	if err != nil {
		m.ours = "An error has occured when reading file contents."
		m.theirs = "please try again later, or perhaps another file."
	}

	if file.Size() > 100*1000 { // bigger than 100kb
		m.ours = "File size is too big to render."
		m.theirs = "Yeah, what ours said"
	}

	data, err := os.ReadFile(m.path)
	rows := strings.Split(string(data), "\n")

	m.ours = ""
	m.theirs = ""

	inOurs := false
	inTheirs := false

	for _, element := range rows {
		if strings.HasPrefix(element, "<<<<<<< HEAD") {
			inOurs = true
			inTheirs = false
		}

		if strings.HasPrefix(element, "=======") {
			inOurs = false
			inTheirs = true
		}

		if strings.HasPrefix(element, ">>>>>>>") {
			inOurs = false
			inTheirs = false
		}

		if !inOurs && !inTheirs {
			m.ours += element + "\n"
			m.theirs += element + "\n"
		}

		if inOurs {
			m.ours += element + "\n"
		}

		if inTheirs {
			m.theirs += element + "\n"
		}
	}
}

func getWidth(width int) int {
	return width
}

func getHeight(height int) int {
	return height - 2
}
