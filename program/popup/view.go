package popup

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	if !m.Visible {
		return ""
	}

	containerStyle := lipgloss.NewStyle().
		Width(30).
		Height(10).
		Border(lipgloss.ThickBorder())

	titleStyle := lipgloss.NewStyle().
		Width(28).
		Height(1).
		Underline(true)

	contentStyle := lipgloss.NewStyle().Padding(1).Bold(true)

	optionStyle := lipgloss.NewStyle().Margin(1, 2)

	buttonStyle := lipgloss.NewStyle().Background(lipgloss.Color("#FFFFFF")).Foreground(lipgloss.Color("#000000"))

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center,
		containerStyle.Render(
			titleStyle.Render("Attention!")+"\n"+
				contentStyle.Render("Are you sure you want to discard "+m.Name+"?")+"\n"+
				lipgloss.PlaceHorizontal(8, lipgloss.Bottom,
					optionStyle.Render(buttonStyle.Render("Y")+" Yes")+"  "+
						optionStyle.Render(buttonStyle.Render("N")+" No")),
		),
	)
}
