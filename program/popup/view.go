package popup

import "github.com/charmbracelet/lipgloss"

var (
	cancelColor  = lipgloss.Color("#CCCCCC")
	discardColor = lipgloss.Color("#FA7970")
)

func (m Model) View() string {
	if !m.Visible {
		return ""
	}

	containerStyle := lipgloss.NewStyle().
		Width(m.Width).
		Height(m.Height).
		Border(lipgloss.ThickBorder()).AlignHorizontal(lipgloss.Center)

	titleStyle := lipgloss.NewStyle().
		Width(m.Width - 2).AlignHorizontal(lipgloss.Center).Background(lipgloss.Color("#21262D"))

	contentStyle := lipgloss.NewStyle().Padding(1).Bold(true)

	optionStyle := lipgloss.NewStyle().
		Width((m.Width - 5) / 2).
		Foreground(lipgloss.Color("#21262D")).
		AlignHorizontal(lipgloss.Center)

	yesButtonStyle := lipgloss.NewStyle().Background(discardColor)
	noButtonStyle := lipgloss.NewStyle().Background(cancelColor)

	buttons := lipgloss.NewStyle().Foreground(cancelColor).Render("N ") +
		noButtonStyle.Inherit(optionStyle).Render("Cancel") + "  " +
		lipgloss.NewStyle().Foreground(discardColor).Render("Y ") +
		yesButtonStyle.Inherit(optionStyle).Render("Discard")

	return containerStyle.Render(
		titleStyle.Render("Attention!") + "\n" +
			contentStyle.Render("Are you sure you want to discard "+m.Name+"?") + "\n" + buttons,
	)
}
