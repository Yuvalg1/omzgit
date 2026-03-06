package conflict

import (
	"omzgit/consts"
	"omzgit/default/colors"
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

var (
	CUTOFF       = 50
	contentStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, true, true).
			BorderBackground(bg.C[0]).
			Background(bg.C[0])
)

func (m Model) View() string {
	if m.width > CUTOFF {
		our := getOurAxis(m.width)
		their := getTheirAxis(m.width)
		inverse := getInverseAxis(m.height - 1)

		return lipgloss.JoinHorizontal(lipgloss.Top,
			m.getOurContent(our, inverse),
			m.getTheirContent(their, inverse),
		)
	} else {
		our := getOurAxis(m.height - 2)
		their := getTheirAxis(m.height - 2)
		inverse := getInverseAxis(m.width)

		return lipgloss.JoinVertical(lipgloss.Top,
			m.getOurContent(inverse, our),
			m.getTheirContent(inverse, their))
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
			ourStyle.Render(m.ours),
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
			theirStyle.Render(m.theirs),
	)
}

func getOurAxis(size int) int {
	return size/2 + 1
}

func getTheirAxis(size int) int {
	return size/2 - (size+1)%2
}

func getInverseAxis(size int) int {
	return size
}
