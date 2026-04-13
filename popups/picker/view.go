package picker

import (
	"omzgit/consts"
	"omzgit/default/colors"
	"omzgit/default/colors/bg"
	"omzgit/default/colors/gray"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	return lipgloss.NewStyle().
		Background(bg.C[0]).
		Foreground(colors.Yellow).
		Render(consts.PadTitle("more options", m.width) +

			lipgloss.NewStyle().
				Background(bg.C[0]).
				BorderBackground(bg.C[0]).
				BorderForeground(colors.Yellow).
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
		Foreground(colors.Yellow).
		Render(string(letter) + " ")
}

func (m Model) renderTitle(letter string) string {
	return lipgloss.NewStyle().Background(bg.C[0]).Foreground(gray.C[2]).Render(m.options[letter].Desc)
}
