package popups

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Func[F any] struct {
	Func F
}

type visible interface {
	GetVisible() bool
}

type InnerModel interface {
	visible
	tea.Model
}

type Model[T InnerModel] struct {
	current string
	Popups  map[string]T
}

func InitialModel[T InnerModel](current string) Model[InnerModel] {
	return Model[InnerModel]{
		current: current,
		Popups:  map[string]InnerModel{},
	}
}

func (m Model[T]) Init() tea.Cmd {
	return nil
}

func (m *Model[T]) AddPopup(name string, popup T) {
	m.Popups[name] = popup
}

func (m Model[T]) GetCurrent() T {
	current := m.Popups[m.current]
	return current
}
