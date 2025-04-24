package diff

func (m Model) View() string {
	m.viewport.SetValue(m.Content)

	return m.viewport.View()
}
