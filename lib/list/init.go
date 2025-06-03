package list

import (
	"program/messages"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model[T tea.Model] struct {
	Children  []T
	ActiveRow int

	filterFn  func(row T, text string) bool
	TextInput textinput.Model

	width  int
	height int
}

func InitialModel[T tea.Model](width int, height int, children []T) Model[T] {
	var childrenTeaModels []T
	childrenTeaModels = append(childrenTeaModels, children...)

	ti := textinput.New()
	ti.CharLimit = 20
	ti.Width = width

	return Model[T]{
		Children:  childrenTeaModels,
		ActiveRow: 0,

		filterFn: func(row T, text string) bool {
			return true
		},
		TextInput: ti,

		width:  width,
		height: height,
	}
}

func (m Model[T]) Init() tea.Cmd {
	return nil
}

func (m Model[T]) ModeCmd(mode string) tea.Cmd {
	return func() tea.Msg {
		return messages.ModeMsg{Mode: mode}
	}
}

func getWidth(width int) int {
	return width
}

func getHeight(height int) int {
	return height
}

func (m *Model[T]) SetContent(children []T) {
	m.Children = children
	m.Children = m.getFilteredChildren()
}

func (m Model[T]) UpdateContent(msg tea.Msg) (Model[T], tea.Cmd) {
	var cmds []tea.Cmd
	for index, element := range m.Children {
		res, cmd := element.Update(msg)
		m.Children[index] = res.(T)

		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m Model[T]) GetCurrent() *T {
	return &m.Children[m.ActiveRow]
}

func (m Model[T]) UpdateCurrent(msg tea.Msg) (Model[T], tea.Cmd) {
	res, cmd := m.Children[m.ActiveRow].Update(msg)
	m.Children[m.ActiveRow] = res.(T)
	return m, cmd
}

func (m *Model[T]) SetFilterFn(fn func(row T, text string) bool) {
	m.filterFn = fn
}

func (m *Model[T]) ResetValue() {
	m.TextInput.SetValue("")
}
