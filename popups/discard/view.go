package discard

import (
	"omzgit/consts"
	"omzgit/default/colors/bg"
	"omzgit/default/style"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	cancelColor  = lipgloss.Color("#CCCCCC")
	discardColor = lipgloss.Color("#FA7970")
)

func (m Model) View() string {
	containerStyle := style.Bg.
		Width(m.Width-2).
		Height(m.Height).
		Border(lipgloss.NormalBorder(), false, true, true).
		AlignHorizontal(lipgloss.Center)

	contentStyle := style.Bg.Padding(1).Bold(true)

	optionStyle := style.Bg.
		Width((m.Width - 7) / 2).
		Foreground(lipgloss.Color("#21262D")).
		AlignHorizontal(lipgloss.Center)

	yesButtonStyle := style.Bg.Background(discardColor)
	noButtonStyle := style.Bg.Background(cancelColor)

	buttons := style.Bg.Foreground(cancelColor).Render("N ") +
		noButtonStyle.Inherit(optionStyle).Render("Cancel") + style.Bg.Render("  ") +
		style.Bg.Foreground(discardColor).Render("Y ") +
		yesButtonStyle.Inherit(optionStyle).Render(cases.Title(language.English).String(m.verb))

	return style.Bg.Foreground(bg.C[4]).Render(consts.PadTitle("Attention", m.Width) + containerStyle.Render(
		contentStyle.Render("Are you sure you want to "+m.verb+" "+m.Name+"?")+"\n"+buttons))
}
