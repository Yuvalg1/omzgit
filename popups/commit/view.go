package commit

import (
	"omzgit/consts"
	"omzgit/default/colors"
	"omzgit/default/colors/bg"
	"omzgit/default/colors/gray"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	titleStyle := lipgloss.NewStyle().Background(bg.C[0])
	return titleStyle.Render(consts.PadTitle("commit "+m.commitMessageType, m.width) +
		lipgloss.NewStyle().
			Background(bg.C[0]).
			Border(lipgloss.NormalBorder(), false, true, true).
			Height(m.height-1).
			Width(m.width-2).
			Render(
				m.textinput.View()+m.renderMoreOptions()))
}

func (m Model) renderOption(letter byte, desc string) string {
	return m.renderLetter(letter) + m.renderTitle(letter, " "+desc)
}

func (m Model) renderLetter(letter byte) string {
	if len(m.options[letter]) == 0 {
		return lipgloss.NewStyle().
			Background(bg.Dim).
			Foreground(colors.Yellow).
			Render(string(letter))
	}
	return lipgloss.NewStyle().
		Background(colors.Yellow).
		Foreground(bg.Dim).
		Render(string(letter))
}

func (m Model) renderTitle(letter byte, title string) string {
	if len(m.options[letter]) == 0 {
		return lipgloss.NewStyle().Background(bg.C[0]).Foreground(gray.C[2]).Render(title)
	}
	return lipgloss.NewStyle().Background(bg.C[0]).Render(title)
}

func (m Model) renderMoreOptions() string {
	if m.moreOptions {
		return "\n" +
			m.renderOption('a', "--amend") + "\n" +
			m.renderOption('e', "--edit") + "\n" +
			m.renderOption('E', "--no-edit") + "\n" +
			m.renderOption('n', "--no-verify") + "\n" +
			m.renderOption('y', "--allow-empty")
	}

	return lipgloss.NewStyle().Background(bg.Dim).Foreground(colors.Yellow).Render("o") +
		lipgloss.NewStyle().Background(bg.C[0]).Render(" options")
}
