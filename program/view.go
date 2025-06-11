package program

import (
	"program/default/colors/bg"
	"program/overlay"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	popup := m.Popup.View()

	current := m.Popup.GetCurrent()
	if current.GetVisible() {
		return lipgloss.NewStyle().
			Render(overlay.PlaceOverlay((m.Width-lipgloss.Width(popup))/2, (m.Height-1-lipgloss.Height(popup))/2, m.Popup.View(), m.Tabs[m.ActiveTab].View()+"\n"+m.cokeline.View()))
	}

	return lipgloss.NewStyle().
		Background(bg.C[0]).
		Render(m.Tabs[m.ActiveTab].View() + "\n" + m.cokeline.View())
}
