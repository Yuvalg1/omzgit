package discard

import (
	"program/consts"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	cancelColor  = lipgloss.Color("#CCCCCC")
	discardColor = lipgloss.Color("#FA7970")
)

func (m Model) View() string {
	containerStyle := lipgloss.NewStyle().
		Width(m.Width-2).
		Height(m.Height).
		Border(lipgloss.NormalBorder(), false, true, true).
		AlignHorizontal(lipgloss.Center)

	contentStyle := lipgloss.NewStyle().Padding(1).Bold(true)

	optionStyle := lipgloss.NewStyle().
		Width((m.Width - 7) / 2).
		Foreground(lipgloss.Color("#21262D")).
		AlignHorizontal(lipgloss.Center)

	yesButtonStyle := lipgloss.NewStyle().Background(discardColor)
	noButtonStyle := lipgloss.NewStyle().Background(cancelColor)

	buttons := lipgloss.NewStyle().Foreground(cancelColor).Render("N ") +
		noButtonStyle.Inherit(optionStyle).Render("Cancel") + "  " +
		lipgloss.NewStyle().Foreground(discardColor).Render("Y ") +
		yesButtonStyle.Inherit(optionStyle).Render(cases.Title(language.English).String(m.verb))

	return consts.PadTitle("Attention", m.Width) + containerStyle.Render(
		contentStyle.Render("Are you sure you want to "+m.verb+" "+m.Name+"?")+"\n"+buttons,
	)
}
