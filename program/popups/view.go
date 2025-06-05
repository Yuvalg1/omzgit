package popups

func (m Model[T]) View() string {
	return m.Popups[m.current].View()
}
