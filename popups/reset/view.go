package reset

import (
	"omzgit/consts"
	"omzgit/default/colors"
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	title := "choose a reset type for " + m.name

	return lipgloss.NewStyle().
		Background(bg.C[0]).
		Render(consts.PadTitle("reset", m.width) +

			lipgloss.NewStyle().
				Background(bg.C[0]).
				BorderBackground(bg.C[0]).
				Border(lipgloss.NormalBorder(), false, true, true).
				Height(m.height).
				Width(m.width-2).
				Render(title+"\n"+
					m.renderOption('s')+"\n"+
					m.renderOption('h')+"\n"+
					m.renderOption('m')))
}

func (m Model) renderOption(letter byte) string {
	return m.renderLetter(letter) + m.renderTitle(letter)
}

func (m Model) renderLetter(letter byte) string {
	return lipgloss.NewStyle().
		Background(bg.C[0]).
		Foreground(colors.Yellow).
		Render(string(letter) + " ")
}

func (m Model) renderTitle(letter byte) string {
	return lipgloss.NewStyle().Background(bg.C[0]).Foreground(colors.Purple).Render(m.options[letter])
}
