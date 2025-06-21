package discard

import (
	"omzgit/consts"
	"omzgit/default/colors"
	"omzgit/default/colors/bg"
	"omzgit/default/colors/gray"
	"omzgit/default/style"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	cancelColor  = gray.C[2]
	discardColor = colors.Red
)

func (m Model) View() string {
	containerStyle := lipgloss.NewStyle().
		AlignHorizontal(lipgloss.Center).
		Background(bg.C[0]).
		BorderBackground(bg.C[0]).
		BorderForeground(bg.C[4]).
		Border(lipgloss.NormalBorder(), false, true, true).
		Height(m.Height).
		Width(m.Width - 2)

	contentStyle := style.Bg.Padding(1).Bold(true)

	optionStyle := style.Bg.
		Width((m.Width - 7) / 2).
		Foreground(bg.C[0]).
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
