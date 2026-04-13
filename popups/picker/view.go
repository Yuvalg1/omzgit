package picker

import (
	"omzgit/consts"
	"omzgit/default/colors"
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	return lipgloss.NewStyle().
		Background(bg.C[0]).
		Render(consts.PadTitle("reset", m.width) +

			lipgloss.NewStyle().
				Background(bg.C[0]).
				BorderBackground(bg.C[0]).
				Border(lipgloss.NormalBorder(), false, true, true).
				Height(m.height).
				Width(m.width-2).
				Render(m.title+m.renderOptions()))
}

func (m Model) renderOptions() string {
	options := ""
	for key := range m.options {
		options += "\n" + m.renderOption(key)
	}
	return options
}

func (m Model) renderOption(letter string) string {
	return m.renderLetter(letter) + m.renderTitle(letter)
}

func (m Model) renderLetter(letter string) string {
	return lipgloss.NewStyle().
		Background(bg.C[0]).
		Bold(true).
		Foreground(colors.Red).
		Render(string(letter) + " ")
}

func (m Model) renderTitle(letter string) string {
	return lipgloss.NewStyle().Background(bg.C[0]).Foreground(colors.Orange).Render(m.options[letter].Desc)
}
