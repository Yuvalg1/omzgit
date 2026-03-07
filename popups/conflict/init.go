package conflict

import (
	"os"
	"strings"

	"omzgit/default/colors"
	"omzgit/default/colors/bg"
	"omzgit/popups/conflict/content"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var CUTOFF = 50

type Model struct {
	path     string
	visible  bool
	conflict int

	ours   content.Model
	theirs content.Model

	width  int
	height int
}

func InitialModel(width int, height int) Model {
	return Model{
		visible: false,
		ours:    content.InitialModel(width, height, true),
		theirs:  content.InitialModel(width, height, false),

		width:  width,
		height: height,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width
}

func getHeight(height int) int {
	return height - 2
}

func (m Model) getOurAxis() int {
	if m.width > CUTOFF {
		return m.width/2 + 1
	}
	return m.height / 2
}

func (m Model) getTheirAxis() int {
	if m.width > CUTOFF {
		return m.width/2 - (m.width+1)%2
	}
	return m.height/2 - (m.height+1)%2 - 1
}

func (m Model) GetVisible() bool {
	return m.visible
}

func (m *Model) getContent() { // inConflict is needed for malformed files
	if m.width > CUTOFF {
		m.ours = content.InitialModel(m.getOurAxis(), m.height-1, true)
		m.theirs = content.InitialModel(m.getTheirAxis(), m.height-1, false)
	} else {
		m.ours = content.InitialModel(m.width, m.getOurAxis(), true)
		m.theirs = content.InitialModel(m.width, m.getTheirAxis(), false)
	}

	ourContent := ""
	theirContent := ""

	file, err := os.Stat(m.path)
	if err != nil {
		ourContent = "An error has occured when reading file contents."
		theirContent = "please try again later, or perhaps another file."
	}

	if file.Size() > 100*1000 { // bigger than 100kb
		ourContent = "File size is too big to render."
		theirContent = "Yeah, what ours said"
	}

	data, err := os.ReadFile(m.path)
	rows := strings.Split(string(data), "\n")

	inOurs := false
	inTheirs := false

	ourRowStyle := lipgloss.NewStyle().Width(m.ours.Content.Width)
	theirRowStyle := lipgloss.NewStyle().Width(m.theirs.Content.Width)
	greenStyle := lipgloss.NewStyle().Background(bg.C[0]).Foreground(colors.Green).UnsetWidth()
	redStyle := lipgloss.NewStyle().Background(bg.C[0]).Foreground(colors.Red).UnsetWidth()

	for _, element := range rows {
		if strings.HasPrefix(element, "<<<<<<< HEAD") {
			inOurs = true
			inTheirs = false
			continue
		}

		if strings.HasPrefix(element, "=======") {
			inOurs = false
			inTheirs = true
			continue
		}

		if strings.HasPrefix(element, ">>>>>>>") {
			inOurs = false
			inTheirs = false
			continue
		}

		if !inOurs && !inTheirs {
			ourContent += ourRowStyle.Render(element) + "\n"
			theirContent += theirRowStyle.Render(element) + "\n"
		}

		if inOurs {
			ourContent += redStyle.Render(ourRowStyle.Render(element)) + "\n"
		}

		if inTheirs {
			theirContent += greenStyle.Render(theirRowStyle.Render(element)) + "\n"
		}

	}
	m.ours.SetContent(ourContent)
	m.theirs.SetContent(theirContent)
}
