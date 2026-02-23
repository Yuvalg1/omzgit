package conflict

import (
	"omzgit/consts"
	"omzgit/default/colors"
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

var CUTOFF = 40

func (m Model) View() string {
	titleStyle := lipgloss.NewStyle().Background(bg.C[0])

	dimensions := m.getContentDimensions()

	ourStyle := titleStyle.BorderForeground(colors.Aqua).Foreground(colors.Aqua)
	theirStyle := titleStyle.BorderForeground(colors.Orange).Foreground(colors.Orange)

	if m.width > CUTOFF {
		return lipgloss.JoinHorizontal(lipgloss.Top,
			ourStyle.Render(consts.PadTitle("ours", m.width/2)+dimensions.Render(m.ours)),
			theirStyle.Render(consts.PadTitle("theirs", m.width/2)+dimensions.Render(m.theirs)),
			consts.PadTitle(m.ours, m.width),
		)
	}

	return ourStyle.Render(consts.PadTitle("ours", m.width)+dimensions.Render(m.ours)) +
		theirStyle.Render(consts.PadTitle("theirs", m.width)+dimensions.Render(m.theirs))
}

func (m Model) getContentDimensions() lipgloss.Style {
	borderStyle := lipgloss.NewStyle().
		Background(bg.C[0]).
		Border(lipgloss.NormalBorder(), false, true, true).
		BorderBackground(bg.C[0])

	if m.width > CUTOFF {
		return lipgloss.NewStyle().Width(m.width/2 - 2).Height(m.height - 2).Inherit(borderStyle)
	}

	return lipgloss.NewStyle().Width(m.width - 2).Height(m.height/2 - 1).Inherit(borderStyle)
}
