package program

import (
	"program/consts"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	popup := m.Popup.View()

	if m.Popup.Visible {
		return consts.PlaceOverlay((m.Width-lipgloss.Width(popup))/2, (m.Height-1-lipgloss.Height(popup))/2, m.Popup.View(), m.Tabs[m.ActiveTab].View()+"\n"+m.cokeline.View())
	}

	return m.Tabs[m.ActiveTab].View() + "\n" + m.cokeline.View()
}
