package commit

import (
	"program/consts"
	"program/default/colors"
	"program/default/colors/bg"
	"program/default/colors/gray"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	return consts.PadTitle("commit "+m.commitMessageType, m.width) + lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, true).
		Height(m.height-1).
		Width(m.width-2).
		Render(
			m.textinput.View()+m.renderMoreOptions()+"\n"+m.getCommitStringString())
}

func (m Model) renderOption(letter byte, desc string) string {
	return m.renderLetter(letter) + m.renderTitle(letter, " "+desc)
}

func (m Model) renderLetter(letter byte) string {
	if len(m.options[letter]) == 0 {
		return lipgloss.NewStyle().
			Background(bg.Dim).
			Foreground(colors.Aqua).
			Render(string(letter))
	}
	return lipgloss.NewStyle().
		Background(colors.Aqua).
		Foreground(bg.Dim).
		Render(string(letter))
}

func (m Model) renderTitle(letter byte, title string) string {
	if len(m.options[letter]) == 0 {
		return lipgloss.NewStyle().Foreground(gray.C[2]).Render(title)
	}
	return lipgloss.NewStyle().Render(title)
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

	return lipgloss.NewStyle().Render(
		lipgloss.NewStyle().Background(bg.Dim).Foreground(colors.Aqua).Render("o") + " options")
}
