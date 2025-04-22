package program

func (m Model) View() string {
	if m.Popup.Visible {
		return m.Popup.View()
	}

	return m.Tabs[m.ActiveTab].View()
}
