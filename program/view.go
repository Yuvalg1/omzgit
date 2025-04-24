package program

func (m Model) View() string {
	if m.Popup.Visible {
		return m.cokeline.View() + "\n" + m.Popup.View()
	}

	return m.Tabs[m.ActiveTab].View() + "\n" + m.cokeline.View()
}
