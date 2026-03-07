package content

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Conflict struct {
	Row    int
	Length int
}

type Model struct {
	Content   viewport.Model
	ours      bool
	conflicts []Conflict

	activeConflict int
}

func InitialModel(width int, height int, ours bool) Model {
	viewport := viewport.New(getWidth(width), getHeight(height))

	return Model{
		Content:   viewport,
		ours:      ours,
		conflicts: []Conflict{},

		activeConflict: 0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width - 2
}

func getHeight(height int) int {
	return height - 1
}

func (m *Model) SetContent(content string) {
	m.Content.SetContent(content)
}

func (m *Model) AppendConflict(conflict Conflict) {
	m.conflicts = append(m.conflicts, conflict)
}
