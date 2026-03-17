package conflict

import (
	"os"
	"strings"

	"omzgit/popups/conflict/chunk"
	"omzgit/popups/conflict/content"

	tea "github.com/charmbracelet/bubbletea"
)

var CUTOFF = 50

type Model struct {
	path    string
	visible bool

	ours   content.Model
	theirs content.Model

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

func getWidth(width int) int {
	return width
}

func getHeight(height int) int {
	return height - 2
}

func (m Model) getOurAxis() (int, int) {
	if m.width > CUTOFF {
		return m.width/2 + 1, m.height - 1
	}
	return m.width, m.height / 2
}

func (m Model) getTheirAxis() (int, int) {
	if m.width > CUTOFF {
		return m.width/2 - (m.width+1)%2, m.height - 1
	}
	return m.width, m.height/2 - (m.height+1)%2 - 1
}

func (m Model) GetVisible() bool {
	return m.visible
}

func (m *Model) getContent() {
	m.ours.Content.SetContent("")
	m.theirs.Content.SetContent("")

	ourWidth, ourHeight := m.getOurAxis()
	m.ours = content.InitialModel(ourWidth, ourHeight, true)

	theirWidth, theirHeight := m.getTheirAxis()
	m.theirs = content.InitialModel(theirWidth, theirHeight, false)

	file, err := os.Stat(m.path)
	if err != nil {
		m.ours.Content.SetContent("An error has occured when reading file contents.")
		m.theirs.Content.SetContent("please try again later, or perhaps another file.")
		return
	}

	if file.Size() > 100*1000 { // bigger than 100kb
		m.ours.Content.SetContent("File size is too big to render.")
		m.theirs.Content.SetContent("Yeah, what ours said")
		return
	}

	data, err := os.ReadFile(m.path)
	rows := strings.Split(string(data), "\n")
	rows = rows[:len(rows)-1]

	ourChunk := chunk.InitialModel(false, true, ourWidth)
	theirChunk := chunk.InitialModel(false, false, theirWidth)

	inOurs := false
	inTheirs := false

	for _, element := range rows {
		if strings.HasPrefix(element, "<<<<<<< HEAD") {
			inOurs = true
			inTheirs = false

			m.ours.Append(ourChunk)
			m.theirs.Append(theirChunk)

			ourChunk = chunk.InitialModel(true, true, ourWidth)
			theirChunk = chunk.InitialModel(false, false, theirWidth)

			continue
		}

		if strings.HasPrefix(element, "=======") {
			inOurs = false
			inTheirs = true

			m.ours.Append(ourChunk)
			theirChunk = chunk.InitialModel(true, false, theirWidth)

			continue
		}

		if strings.HasPrefix(element, ">>>>>>>") {
			inOurs = false
			inTheirs = false

			m.theirs.Append(theirChunk)

			ourChunk = chunk.InitialModel(false, true, ourWidth)
			theirChunk = chunk.InitialModel(false, false, theirWidth)

			continue
		}

		if !inOurs && !inTheirs {
			ourChunk.Append(element)
			theirChunk.Append(element)
		}

		if inOurs {
			ourChunk.Append(element)
		}

		if inTheirs {
			theirChunk.Append(element)
		}

	}
	if ourChunk.Content != "" {
		m.ours.Append(ourChunk)
	}
	if theirChunk.Content != "" {
		m.theirs.Append(theirChunk)
	}

	m.ours.Refresh()
	m.theirs.Refresh()
}
