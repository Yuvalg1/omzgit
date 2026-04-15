package list

import (
	"strings"
	"time"

	"omzgit/git"

	"github.com/bep/debounce"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model[T tea.Model] struct {
	Children  []T
	ActiveRow int
	basename  string

	Page int

	createChildFn func(name string) *T
	debounceFn    func(f func())

	TextInput textinput.Model
	emptyMsg  string
	mode      string

	innerOffset int
	width       int
	height      int
}

func InitialModel[T tea.Model](height int, children []T, initialActive int, emptyMsg string) Model[T] {
	ti := textinput.New()
	ti.CharLimit = 20
	ti.Prompt = ""

	repo, _ := git.Exec("rev-parse", "--show-toplevel")
	parts := strings.Split(repo, "/")
	basename := strings.TrimSpace(parts[len(parts)-1])

	return Model[T]{
		Children:  children,
		ActiveRow: initialActive,

		Page: 0,

		createChildFn: func(name string) *T { return nil },
		debounceFn:    debounce.New(300 * time.Millisecond),
		TextInput:     ti,
		emptyMsg:      emptyMsg,
		basename:      basename,

		innerOffset: min(getHeight(height)-2, initialActive),
		height:      getHeight(height),
	}
}

func (m Model[T]) Init() tea.Cmd {
	return nil
}

func getHeight(height int) int {
	return height
}

func (m *Model[T]) SetContent(children []T) {
	m.Children = children

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

func (m Model[T]) NewSize() int {
	return (m.Page + 1) * (m.height - 1)
}

func (m Model[T]) debounceCmd(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		channel := make(chan tea.Msg)
		m.debounceFn(func() {
			channel <- msg
		})
		return <-channel
	}
}
