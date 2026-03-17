package conflict

import (
	"omzgit/consts"
	"omzgit/default/colors"
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

var contentStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), false, true, true).
	BorderBackground(bg.C[0]).
	Background(bg.C[0])

func (m Model) View() string {
	ourWidth, ourHeight := m.getOurAxis()
	theirWidth, theirHeight := m.getTheirAxis()

	if m.width > CUTOFF {
		return lipgloss.JoinHorizontal(lipgloss.Top,
			m.getOurContent(ourWidth, ourHeight),
			m.getTheirContent(theirWidth, theirHeight),
		)
	} else {
		return lipgloss.JoinVertical(lipgloss.Top,
			m.getOurContent(ourWidth, ourHeight),
			m.getTheirContent(theirWidth, theirHeight),
		)
	}
}

func (m Model) getOurContent(width int, height int) string {
	startStyle := lipgloss.NewStyle().
		Background(bg.C[0]).
		Foreground(colors.Orange)

	ourStyle := lipgloss.NewStyle().
		BorderForeground(colors.Orange).
		Height(height - 1).
		Width(width - 2).
		Inherit(contentStyle)

	return startStyle.Render(
		consts.PadTitle("ours", width) +
			ourStyle.Render(m.ours.View()),
	)
}

func (m Model) getTheirContent(width int, height int) string {
	endStyle := lipgloss.NewStyle().
		Background(bg.C[0]).
		Foreground(colors.Aqua)

	theirStyle := lipgloss.NewStyle().
		BorderForeground(colors.Aqua).
		Height(height - 1).
		Width(width - 2).
		Inherit(contentStyle)

	return endStyle.Render(
		consts.PadTitle("theirs", width) +
			theirStyle.Render(m.theirs.View()),
	)
}
