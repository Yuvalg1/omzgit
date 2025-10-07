package list

import (
	"omzgit/messages"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model[T tea.Model] struct {
	Children  []T
	ActiveRow int

	createChildFn func(name string) *T
	filterFn      func(row T, text string) bool
	TextInput     textinput.Model
	emptyMsg      string

	height int
}

func InitialModel[T tea.Model](height int, children []T, initialActive int, emptyMsg string) Model[T] {
	ti := textinput.New()
	ti.CharLimit = 20

	return Model[T]{
		Children:  children,
		ActiveRow: initialActive,

		createChildFn: func(name string) *T { return nil },
		filterFn: func(row T, text string) bool {
			return true
		},
		TextInput: ti,
		emptyMsg:  emptyMsg,

		height: getHeight(height),
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

func getHeight(height int) int {
	return height
}

func (m *Model[T]) SetContent(children []T) {
	m.Children = children
	m.Children = m.getFilteredChildren()

	if len(m.Children) == 0 {
		m.Children = append(m.Children, *m.createChildFn(m.emptyMsg))
	}

	m.ActiveRow = min(m.ActiveRow, max(len(m.Children)-1, 0))
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

func (m *Model[T]) SetCreateChild(createChildFn func(name string) *T) {
	m.createChildFn = createChildFn
}

func (m *Model[T]) SetFilterFn(fn func(row T, text string) bool) {
	m.filterFn = fn
}

func (m *Model[T]) ResetValue() {
	m.TextInput.SetValue("")
}
